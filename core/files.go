package core

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/anytypeio/go-anytype-library/schema"
	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"
	ipld "github.com/ipfs/go-ipld-format"
	ipfspath "github.com/ipfs/go-path"
	"github.com/textileio/go-textile/core"
	"github.com/textileio/go-textile/ipfs"
	ipfsutil "github.com/textileio/go-textile/ipfs"
	"github.com/textileio/go-textile/mill"
	tpb "github.com/textileio/go-textile/pb"
	tschema "github.com/textileio/go-textile/schema"
	tutil "github.com/textileio/go-textile/util"
)

func (a *Anytype) FileByHash(hash string) (File, error) {
	indexes, err := a.getFileIndexByTarget(hash)
	if err != nil {
		return nil, err
	}
	if len(indexes) == 0 {
		return nil, fmt.Errorf("file not found")
	}

	fileIndex := indexes[0]
	return &file{
		hash:  hash,
		index: &fileIndex,
		node:  a,
	}, nil
}

func (a *Anytype) FileAddWithBytes(content []byte, media string, name string) (File, error) {
	fileIndex, err := a.Textile.Node().AddFileIndex(&mill.Blob{}, core.AddFileConfig{
		Input: content,
		Media: media,
		Name:  name,
	})

	node, keys, err := a.Textile.Node().AddNodeFromFiles([]*tpb.FileIndex{fileIndex})
	if err != nil {
		return nil, err
	}

	nodeHash := node.Cid().Hash().B58String()

	err = a.indexFileData(node, nodeHash)
	if err != nil {
		return nil, err
	}

	filesKeysCacheMutex.Lock()
	defer filesKeysCacheMutex.Unlock()
	filesKeysCache[nodeHash] = keys.Files

	return &file{
		hash:  nodeHash,
		index: fileIndex,
		node:  a,
	}, nil
}

func (a *Anytype) FileAddWithReader(content io.Reader, media string, name string) (File, error) {
	// todo: PR textile to be able to use reader instead of bytes
	contentBytes, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}

	return a.FileAddWithBytes(contentBytes, media, name)
}

func (a *Anytype) getFileIndexByTarget(target string) ([]tpb.FileIndex, error) {
	var list []tpb.FileIndex
	rows, err := a.Textile.Node().Datastore().Files().PrepareAndExecuteQuery("SELECT * FROM files WHERE targets=?", target)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var mill, checksum, source, opts, hash, key, media, name string
		var size int64
		var addedInt int64
		var metab []byte
		var targets *string

		if err := rows.Scan(&mill, &checksum, &source, &opts, &hash, &key, &media, &name, &size, &addedInt, &metab, &targets); err != nil {
			log.Errorf("error in db scan: %s", err)
			continue
		}

		meta := &structpb.Struct{}
		if metab != nil {
			if err := jsonpb.Unmarshal(bytes.NewReader(metab), meta); err != nil {
				log.Errorf("failed to unmarshal file meta: %s", err)
				continue
			}
		}

		tlist := make([]string, 0)
		if targets != nil {
			tlist = tutil.SplitString(*targets, ",")
		}

		list = append(list, tpb.FileIndex{
			Mill:     mill,
			Checksum: checksum,
			Source:   source,
			Opts:     opts,
			Hash:     hash,
			Key:      key,
			Media:    media,
			Name:     name,
			Size:     size,
			Added:    tutil.ProtoTs(addedInt),
			Meta:     meta,
			Targets:  tlist,
		})
	}

	return list, nil
}

func (a *Anytype) getFileConfig(reader io.ReadSeeker, filename string, mill mill.Mill, use string, plaintext bool) (*core.AddFileConfig, error) {
	conf := &core.AddFileConfig{}

	if use == "" {
		conf.Name = filename
	} else {
		ref, err := ipfspath.ParsePath(use)
		if err != nil {
			return nil, err
		}
		parts := strings.Split(ref.String(), "/")
		hash := parts[len(parts)-1]
		var file *tpb.FileIndex
		reader, file, err = a.textile().FileContent(hash)
		if err != nil {
			if err == core.ErrFileNotFound {
				// just cat the data from ipfs
				b, err := ipfsutil.DataAtPath(a.ipfs(), ref.String())
				if err != nil {
					return nil, err
				}
				reader = bytes.NewReader(b)
				conf.Use = ref.String()
			} else {
				return nil, err
			}
		} else {
			conf.Use = file.Checksum
		}
	}

	media, err := a.textile().GetMillMedia(reader, mill)
	if err != nil {
		return nil, err
	}
	conf.Media = media
	_, _ = reader.Seek(0, 0)

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	conf.Input = data
	conf.Plaintext = plaintext

	return conf, nil
}

func (a *Anytype) buildDirectory(reader io.ReadSeeker, filename string, sch *tpb.Node) (*tpb.Directory, error) {
	dir := &tpb.Directory{
		Files: make(map[string]*tpb.FileIndex),
	}

	mil, err := schema.GetMill(sch.Mill, sch.Opts)
	if err != nil {
		return nil, err
	}
	if mil != nil {
		conf, err := a.getFileConfig(reader, filename, mil, "", sch.Plaintext)
		if err != nil {
			return nil, err
		}

		added, err := a.textile().AddFileIndex(mil, *conf)
		if err != nil {
			return nil, err
		}
		dir.Files[tschema.SingleFileTag] = added

	} else if len(sch.Links) > 0 {
		// determine order
		steps, err := tschema.Steps(sch.Links)
		if err != nil {
			return nil, err
		}

		// send each link
		for _, step := range steps {
			stepMill, err := schema.GetMill(step.Link.Mill, step.Link.Opts)
			if err != nil {
				return nil, err
			}
			var conf *core.AddFileConfig
			_, _ = reader.Seek(0, 0)
			if step.Link.Use == tschema.FileTag {
				conf, err = a.getFileConfig(
					reader,
					filename,
					stepMill,
					"",
					step.Link.Plaintext,
				)
				if err != nil {
					return nil, err
				}

			} else {
				if dir.Files[step.Link.Use] == nil {
					return nil, fmt.Errorf(step.Link.Use + " not found")
				}

				conf, err = a.getFileConfig(nil,
					filename,
					stepMill,
					dir.Files[step.Link.Use].Hash,
					step.Link.Plaintext,
				)
				if err != nil {
					return nil, err
				}
			}

			added, err := a.textile().AddFileIndex(stepMill, *conf)
			if err != nil {
				return nil, err
			}
			dir.Files[step.Name] = added
		}
	} else {
		return nil, tschema.ErrEmptySchema
	}

	return dir, nil
}

func (a *Anytype) getFileIndexForPath(pth string) (*tpb.FileIndex, error) {
	plaintext, err := ipfs.DataAtPath(a.Textile.Node().Ipfs(), pth+core.MetaLinkName)
	if err != nil {
		return nil, err
	}

	var file tpb.FileIndex
	err = jsonpb.Unmarshal(bytes.NewReader(plaintext), &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

// IndexFileData walks a file data node, indexing file links
func (a *Anytype) indexFileData(inode ipld.Node, data string) error {
	for _, link := range inode.Links() {
		nd, err := ipfs.NodeAtLink(a.ipfs(), link)
		if err != nil {
			return err
		}
		err = a.indexFileNode(nd, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// indexFileNode walks a file node, indexing file links
func (a *Anytype) indexFileNode(inode ipld.Node, data string) error {
	links := inode.Links()

	if looksLikeFileNode(inode) {
		return a.indexFileLink(inode, data)
	}

	for _, link := range links {
		n, err := ipfs.NodeAtLink(a.ipfs(), link)
		if err != nil {
			return err
		}

		err = a.indexFileLink(n, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// indexFileLink indexes a file link
func (a *Anytype) indexFileLink(inode ipld.Node, data string) error {
	dlink := tschema.LinkByName(inode.Links(), core.ValidContentLinkNames)
	if dlink == nil {
		return core.ErrMissingContentLink
	}

	return a.Textile.Node().Datastore().Files().AddTarget(dlink.Cid.Hash().B58String(), data)
}

// looksLikeFileNode returns whether or not a node appears to
// be a textile node. It doesn't inspect the actual data.
func looksLikeFileNode(node ipld.Node) bool {
	links := node.Links()
	if len(links) != 2 {
		return false
	}
	if tschema.LinkByName(links, core.ValidMetaLinkNames) == nil ||
		tschema.LinkByName(links, core.ValidContentLinkNames) == nil {
		return false
	}
	return true
}
