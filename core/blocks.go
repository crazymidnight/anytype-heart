package core

import (
	"fmt"
	"strings"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-library/pb/storage"
	"github.com/gogo/protobuf/types"
	mh "github.com/multiformats/go-multihash"
)

func (a *Anytype) GetBlock(id string) (Block, error) {
	_, err := mh.FromB58String(id)
	if err == nil {
		smartBlock, err := a.GetSmartBlock(id)
		if err != nil {
			return nil, err
		}

		switch strings.ToLower(smartBlock.thread.Schema.Name) {
		case "dashboard":
			return &Dashboard{smartBlock}, nil
		case "page":
			return &Page{smartBlock}, nil
		default:
			return nil, fmt.Errorf("for now only smartblocks are queriable")
		}
	}

	// todo: allow to query simple blocks via smart blocks
	return nil, fmt.Errorf("for now only smartblocks are queriable")
}

func isSmartBlock(block *model.Block) bool {
	switch block.Content.Content.(type) {
	case *model.BlockCoreContentOfDashboard, *model.BlockCoreContentOfPage:
		return true
	default:
		return false
	}
}

func (a *Anytype) blockToVersion(block *model.Block, parentSmartBlockVersion BlockVersion, versionId string, user string, date *types.Timestamp) BlockVersion {
	switch block.Content.Content.(type) {
	case *model.BlockCoreContentOfDashboard, *model.BlockCoreContentOfPage:
		return &SmartBlockVersion{
			model: &storage.BlockWithDependentBlocks{
				Block: block,
			},
			versionId: versionId,
			user:      user,
			date:      date,
			node:      a,
		}
	default:
		return &SimpleBlockVersion{
			model:                   block,
			parentSmartBlockVersion: parentSmartBlockVersion,
			node:                    a,
		}
	}
}

func (a *Anytype) createPredefinedBlocks() error {
	// archive
	thread, err := a.predefinedThreadAdd(threadDerivedIndexArchiveDashboard)
	if err != nil {
		return err
	}
	a.predefinedBlockIds.Archive = thread.Id
	block, err := a.GetBlock(thread.Id)
	if err != nil {
		return err
	}

	if version, _ := block.GetCurrentVersion(); version == nil {
		// version not yet created
		log.Debugf("create predefined archive block")
		_, err = block.AddVersion(&model.Block{
			Id: block.GetId(),
			Fields: &types.Struct{
				Fields: map[string]*types.Value{
					"name": {Kind: &types.Value_StringValue{StringValue: "Archive"}},
					"icon": {Kind: &types.Value_StringValue{StringValue: ":package:"}},
				},
			},
			Content: &model.BlockCore{
				Content: &model.BlockCoreContentOfDashboard{
					Dashboard: &model.BlockContentDashboard{
						Style: model.BlockContentDashboard_Archive,
					},
				},
			},
		})

		if err != nil {
			return err
		}
	}

	// home
	thread, err = a.predefinedThreadAdd(threadDerivedIndexHomeDashboard)
	if err != nil {
		return err
	}
	a.predefinedBlockIds.Home = thread.Id

	block, err = a.GetBlock(thread.Id)
	if err != nil {
		return err
	}

	if version, _ := block.GetCurrentVersion(); version == nil {
		// version not yet created
		log.Debugf("create predefined home block")

		_, err = block.AddVersion(&model.Block{
			Id:          block.GetId(),
			ChildrenIds: []string{a.predefinedBlockIds.Archive},
			Fields: &types.Struct{
				Fields: map[string]*types.Value{
					"name": {Kind: &types.Value_StringValue{StringValue: "Home"}},
					"icon": {Kind: &types.Value_StringValue{StringValue: ":house:"}},
				},
			},
			Content: &model.BlockCore{
				Content: &model.BlockCoreContentOfDashboard{
					Dashboard: &model.BlockContentDashboard{
						Style: model.BlockContentDashboard_MainScreen,
					},
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
