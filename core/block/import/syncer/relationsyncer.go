package syncer

import (
	"strings"

	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

var logger = logging.Logger("import-file-relation-syncer")

type RelationSyncer interface {
	Sync(state *state.State, relationName string) []string
}

type FileRelationSyncer struct {
	service   *block.Service
	fileStore filestore.FileStore
}

func NewFileRelationSyncer(service *block.Service, fileStore filestore.FileStore) RelationSyncer {
	return &FileRelationSyncer{service: service, fileStore: fileStore}
}

func (fs *FileRelationSyncer) Sync(state *state.State, relationName string) []string {
	return fs.handleFileRelation(state, relationName)
}

func (fs *FileRelationSyncer) handleFileRelation(st *state.State, name string) []string {
	allFiles := fs.getFilesFromRelations(st, name)
	allFilesHashes := make([]string, 0)
	filesToDelete := make([]string, 0, len(allFiles))
	for _, f := range allFiles {
		if f == "" {
			continue
		}
		var hash string
		if hash = fs.uploadFile(f); hash != "" {
			allFilesHashes = append(allFilesHashes, hash)
			filesToDelete = append(filesToDelete, hash)
		}
		if hash == "" {
			if _, err := fs.fileStore.ListByTarget(f); err == nil {
				allFilesHashes = append(allFilesHashes, f)
				continue
			}
		}
	}
	fs.updateFileRelationsDetails(st, name, allFilesHashes)

	return filesToDelete
}

func (fs *FileRelationSyncer) getFilesFromRelations(st *state.State, name string) []string {
	var allFiles []string
	if files := pbtypes.GetStringList(st.Details(), name); files != nil {
		allFiles = append(allFiles, files...)
	}

	if files := pbtypes.GetString(st.Details(), name); files != "" {
		allFiles = append(allFiles, files)
	}
	return allFiles
}

func (fs *FileRelationSyncer) uploadFile(file string) string {
	var (
		hash string
		err  error
	)
	if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
		req := pb.RpcFileUploadRequest{LocalPath: file}
		req.Url = file
		req.LocalPath = ""
		hash, err = fs.service.UploadFile(req)
		if err != nil {
			logger.Errorf("file uploading %s", err)
		} else {
			file = hash
		}
	}
	return hash
}

func (fs *FileRelationSyncer) updateFileRelationsDetails(st *state.State, name string, allFilesHashes []string) {
	if st.Details() == nil || st.Details().GetFields() == nil {
		return
	}
	if st.Details().Fields[name].GetListValue() != nil && len(allFilesHashes) != 0 {
		st.SetDetail(name, pbtypes.StringList(allFilesHashes))
	}

	if st.Details().Fields[name].GetStringValue() != "" && len(allFilesHashes) != 0 {
		st.SetDetail(name, pbtypes.String(allFilesHashes[0]))
	}
}
