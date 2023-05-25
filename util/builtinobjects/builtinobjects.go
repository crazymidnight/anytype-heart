package builtinobjects

import (
	"archive/zip"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/anyproto/anytype-heart/core/session"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/anyproto/any-sync/app"

	"github.com/anyproto/anytype-heart/core/block"
	importer "github.com/anyproto/anytype-heart/core/block/import"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/constant"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

const (
	CName            = "builtinobjects"
	injectionTimeout = 30 * time.Second

	// TODO: GO-1387 Need to use profile.pb to handle dashboard injection during migration
	migrationDashboardName = "bafyreiha2hjbrzmwo7rpiiechv45vv37d6g5aezyr5wihj3agwawu6zi3u"
)

//go:embed data/skip.zip
var skipZip []byte

//go:embed data/personal_projects.zip
var personalProjectsZip []byte

//go:embed data/knowledge_base.zip
var knowledgeBaseZip []byte

//go:embed data/notes_diary.zip
var notesDiaryZip []byte

//go:embed data/migration_dashboard.zip
var migrationDashboardZip []byte

var (
	log = logging.Logger("anytype-mw-builtinobjects")

	archives = map[pb.RpcObjectImportUseCaseRequestUseCase][]byte{
		pb.RpcObjectImportUseCaseRequest_SKIP:              skipZip,
		pb.RpcObjectImportUseCaseRequest_PERSONAL_PROJECTS: personalProjectsZip,
		pb.RpcObjectImportUseCaseRequest_KNOWLEDGE_BASE:    knowledgeBaseZip,
		pb.RpcObjectImportUseCaseRequest_NOTES_DIARY:       notesDiaryZip,
	}
)

type BuiltinObjects interface {
	app.Component

	CreateObjectsForUseCase(*session.Context, pb.RpcObjectImportUseCaseRequestUseCase) (code pb.RpcObjectImportUseCaseResponseErrorCode, err error)
	InjectMigrationDashboard() error
}

type builtinObjects struct {
	service     *block.Service
	coreService core.Service
	importer    importer.Importer
	store       objectstore.ObjectStore
}

func New() BuiltinObjects {
	return &builtinObjects{}
}

func (b *builtinObjects) Init(a *app.App) (err error) {
	b.service = a.MustComponent(block.CName).(*block.Service)
	b.coreService = a.MustComponent(core.CName).(core.Service)
	b.importer = a.MustComponent(importer.CName).(importer.Importer)
	b.store = app.MustComponent[objectstore.ObjectStore](a)
	return
}

func (b *builtinObjects) Name() (name string) {
	return CName
}

func (b *builtinObjects) CreateObjectsForUseCase(
	ctx *session.Context, useCase pb.RpcObjectImportUseCaseRequestUseCase,
) (code pb.RpcObjectImportUseCaseResponseErrorCode, err error) {
	start := time.Now()

	archive, found := archives[useCase]
	if !found {
		return pb.RpcObjectImportUseCaseResponseError_BAD_INPUT,
			fmt.Errorf("failed to import builtinObjects: invalid Use Case value: %v", useCase)
	}

	if err = b.inject(ctx, archive, false); err != nil {
		return pb.RpcObjectImportUseCaseResponseError_UNKNOWN_ERROR,
			fmt.Errorf("failed to import builtinObjects for Use Case %s: %s",
				pb.RpcObjectImportUseCaseRequestUseCase_name[int32(useCase)], err.Error())
	}

	spent := time.Now().Sub(start)
	if spent > injectionTimeout {
		log.Debugf("built-in objects injection time exceeded timeout of %s and is %s", injectionTimeout.String(), spent.String())
	}

	return pb.RpcObjectImportUseCaseResponseError_NULL, nil
}

func (b *builtinObjects) InjectMigrationDashboard() error {
	return b.inject(nil, migrationDashboardZip, true)
}

func (b *builtinObjects) inject(ctx *session.Context, archive []byte, isMigration bool) (err error) {
	path := filepath.Join("./", time.Now().Format("tmp.20060102.150405.99")+".zip")
	if err = os.WriteFile(path, archive, 0644); err != nil {
		return fmt.Errorf("failed to save use case archive to temporary file: %s", err.Error())
	}

	if err = b.importArchive(ctx, path); err != nil {
		return err
	}

	b.setCreationDateToday()

	var oldId string
	if isMigration {
		oldId = migrationDashboardName
	} else {
		oldId, err = b.getOldSpaceDashboardId(archive)
		if err != nil {
			log.Errorf("Failed to get old id of space dashboard object: %s", err.Error())
			return nil
		}
	}

	newId, err := b.getNewSpaceDashboardId(oldId)
	if err != nil {
		log.Errorf("Failed to get new id of space dashboard object: %s", err.Error())
		return nil
	}

	b.handleSpaceDashboard(newId)
	return
}

func (b *builtinObjects) importArchive(ctx *session.Context, path string) (err error) {
	if err = b.importer.Import(ctx, &pb.RpcObjectImportRequest{
		UpdateExistingObjects: false,
		Type:                  pb.RpcObjectImportRequest_Pb,
		Mode:                  pb.RpcObjectImportRequest_ALL_OR_NOTHING,
		NoProgress:            true,
		IsMigration:           false,
		Params: &pb.RpcObjectImportRequestParamsOfPbParams{
			PbParams: &pb.RpcObjectImportRequestPbParams{
				Path:         []string{path},
				NoCollection: true,
			}},
	}); err != nil {
		return err
	}

	if err = os.Remove(path); err != nil {
		log.Errorf("failed to remove temporary file %s: %s", path, err.Error())
	}

	return nil
}

func (b *builtinObjects) setCreationDateToday() {
	ids, _, err := b.store.QueryObjectIds(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				Condition:   model.BlockContentDataviewFilter_Exists,
				RelationKey: bundle.RelationKeyOldAnytypeID.String(),
			},
		},
	}, nil)
	if err != nil {
		log.Errorf("Failed to query created built-in objects for setting CreationDate")
		return
	}
	for _, id := range ids {
		if err := b.service.SetDetails(nil, pb.RpcObjectSetDetailsRequest{
			ContextId: id,
			Details: []*pb.RpcObjectSetDetailsDetail{
				{
					Key:   bundle.RelationKeyCreatedDate.String(),
					Value: pbtypes.Float64(float64(time.Now().Unix())),
				},
			},
		}); err != nil {
			log.Errorf("Failed to set CreatedDate relation to built-in object '%s': %s", id, err.Error())
		}
	}
}

func (b *builtinObjects) getOldSpaceDashboardId(archive []byte) (id string, err error) {
	var (
		rd      io.ReadCloser
		openErr error
	)
	zr, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		return "", err
	}
	profileFound := false
	for _, zf := range zr.File {
		if zf.Name == constant.ProfileFile {
			profileFound = true
			rd, openErr = zf.Open()
			if openErr != nil {
				return "", openErr
			}
			break
		}
	}

	if !profileFound {
		return "", fmt.Errorf("no profile file included in archive")
	}

	defer rd.Close()
	data, err := io.ReadAll(rd)

	profile := &pb.Profile{}
	if err = profile.Unmarshal(data); err != nil {
		return "", err
	}
	return profile.SpaceDashboardId, nil
}

func (b *builtinObjects) getNewSpaceDashboardId(oldId string) (id string, err error) {
	ids, _, err := b.store.QueryObjectIds(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				Condition:   model.BlockContentDataviewFilter_Equal,
				RelationKey: bundle.RelationKeyOldAnytypeID.String(),
				Value:       pbtypes.String(oldId),
			},
		},
	}, nil)
	if err == nil && len(ids) > 0 {
		return ids[0], nil
	}
	return "", err
}

func (b *builtinObjects) handleSpaceDashboard(id string) {
	if err := b.service.SetDetails(nil, pb.RpcObjectSetDetailsRequest{
		ContextId: b.coreService.PredefinedBlocks().Account,
		Details: []*pb.RpcObjectSetDetailsDetail{
			{
				Key:   bundle.RelationKeySpaceDashboardId.String(),
				Value: pbtypes.String(id),
			},
		},
	}); err != nil {
		log.Errorf("Failed to set SpaceDashboardId relation to Account object: %s", err.Error())
	}
	b.createSpaceDashboardWidget(id)
}

func (b *builtinObjects) createSpaceDashboardWidget(id string) {
	targetID, err := b.getFirstWidgetBlockId()
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	if _, err := b.service.CreateWidgetBlock(nil, &pb.RpcBlockCreateWidgetRequest{
		ContextId:    b.coreService.PredefinedBlocks().Widgets,
		TargetId:     targetID,
		Position:     model.Block_Top,
		WidgetLayout: model.BlockContentWidget_Link,
		Block: &model.Block{
			Id:          "",
			ChildrenIds: nil,
			Content: &model.BlockContentOfLink{
				Link: &model.BlockContentLink{
					TargetBlockId: id,
					Style:         model.BlockContentLink_Page,
					IconSize:      model.BlockContentLink_SizeNone,
					CardStyle:     model.BlockContentLink_Inline,
					Description:   model.BlockContentLink_None,
				},
			},
		},
	}); err != nil {
		log.Errorf("Failed to link SpaceDashboard to Widget object: %s", err.Error())
	}
}

func (b *builtinObjects) getFirstWidgetBlockId() (string, error) {
	w, err := b.service.GetObject(context.Background(), b.coreService.PredefinedBlocks().Account, b.coreService.PredefinedBlocks().Widgets)
	if err != nil {
		return "", fmt.Errorf("failed to get Widget object: %s", err.Error())
	}
	if len(w.Blocks()) < 2 {
		return "", fmt.Errorf("failed to get first block of Widget object: %s", err.Error())
	}
	if w.Blocks()[1] == nil {
		return "", fmt.Errorf("failed to get id of first block of Widget object: %s", err.Error())
	}
	return w.Blocks()[1].Id, nil
}
