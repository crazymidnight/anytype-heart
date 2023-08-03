package account

import (
	"path/filepath"
	"os"
	"github.com/anyproto/anytype-heart/pb"
	oserror "github.com/anyproto/anytype-heart/util/os"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/filestorage"
	"strings"
	cp "github.com/otiai10/copy"
	"errors"
)

func (s *Service) AccountMove(req *pb.RpcAccountMoveRequest) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	dirs := []string{filestorage.FlatfsDirName}
	conf := s.app.MustComponent(config.CName).(*config.Config)

	configPath := conf.GetConfigPath()
	srcPath := conf.RepoPath
	fileConf := config.ConfigRequired{}
	if err := config.GetFileConfig(configPath, &fileConf); err != nil {
		return domain.WrapErrorWithCode(err, pb.RpcAccountMoveResponseError_FAILED_TO_GET_CONFIG)
	}
	if fileConf.CustomFileStorePath != "" {
		srcPath = fileConf.CustomFileStorePath
	}

	parts := strings.Split(srcPath, string(filepath.Separator))
	accountDir := parts[len(parts)-1]
	if accountDir == "" {
		return domain.WrapErrorWithCode(errors.New("fail to identify account dir"), pb.RpcAccountMoveResponseError_FAILED_TO_IDENTIFY_ACCOUNT_DIR)
	}

	destination := filepath.Join(req.NewPath, accountDir)
	if srcPath == destination {
		return domain.WrapErrorWithCode(errors.New("source path should not be equal destination path"), pb.RpcAccountMoveResponseError_FAILED_TO_CREATE_LOCAL_REPO)
	}

	if _, err := os.Stat(destination); !os.IsNotExist(err) { // if already exist (in case of the previous fail moving)
		if err := removeDirsRelativeToPath(destination, dirs); err != nil {
			return domain.WrapErrorWithCode(oserror.TransformError(err), pb.RpcAccountMoveResponseError_FAILED_TO_REMOVE_ACCOUNT_DATA)
		}
	}

	err := os.MkdirAll(destination, 0700)
	if err != nil {
		return domain.WrapErrorWithCode(oserror.TransformError(err), pb.RpcAccountMoveResponseError_FAILED_TO_CREATE_LOCAL_REPO)
	}

	err = s.Stop()
	if err != nil {
		return domain.WrapErrorWithCode(err, pb.RpcAccountMoveResponseError_FAILED_TO_STOP_NODE)
	}

	for _, dir := range dirs {
		if _, err := os.Stat(filepath.Join(srcPath, dir)); !os.IsNotExist(err) { // copy only if exist such dir
			if err := cp.Copy(filepath.Join(srcPath, dir), filepath.Join(destination, dir), cp.Options{PreserveOwner: true}); err != nil {
				return domain.WrapErrorWithCode(err, pb.RpcAccountMoveResponseError_FAILED_TO_CREATE_LOCAL_REPO)
			}
		}
	}

	err = config.WriteJsonConfig(configPath, config.ConfigRequired{CustomFileStorePath: destination})
	if err != nil {
		return domain.WrapErrorWithCode(err, pb.RpcAccountMoveResponseError_FAILED_TO_WRITE_CONFIG)
	}

	if err := removeDirsRelativeToPath(srcPath, dirs); err != nil {
		return domain.WrapErrorWithCode(oserror.TransformError(err), pb.RpcAccountMoveResponseError_FAILED_TO_REMOVE_ACCOUNT_DATA)
	}

	if srcPath != conf.RepoPath { // remove root account dir, if move not from anytype source dir
		if err := os.RemoveAll(srcPath); err != nil {
			return domain.WrapErrorWithCode(oserror.TransformError(err), pb.RpcAccountMoveResponseError_FAILED_TO_REMOVE_ACCOUNT_DATA)
		}
	}
	return nil
}

func removeDirsRelativeToPath(rootPath string, dirs []string) error {
	for _, dir := range dirs {
		if err := os.RemoveAll(filepath.Join(rootPath, dir)); err != nil {
			return err
		}
	}
	return nil
}
