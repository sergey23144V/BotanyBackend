package ecomorph

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
)

// EcomorphsServetImpl
type EcomorphsServetImpl struct {
	db *gorm.DB
}

func NewEcomorphsServetImplImpl(db *gorm.DB) EcomorphsServetImpl {
	return EcomorphsServetImpl{db}
}

func (e EcomorphsServetImpl) UpdateEcomorph(ctx context.Context, ecomorph *InputEcomorph) (*Ecomorph, error) {
	createEcomorph, err := DefaultStrictUpdateEcomorph(ctx, e.ToPB(ctx, ecomorph), e.db)
	if err != nil {
		log.Error("Update Ecomorph:", err)
		return nil, err
	}
	log.Debug("Update Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphsServetImpl) GetListEcomorph(ctx context.Context, request *api.EmptyRequest) (*EcomorphsList, error) {
	list, err := DefaultListEcomorph(ctx, e.db)
	if err != nil {
		return nil, err
	}
	return &EcomorphsList{Ecomorph: list}, nil
}

func (e EcomorphsServetImpl) InsertEcomorph(ctx context.Context, ecomorph *InputEcomorph) (*Ecomorph, error) {
	createEcomorph, err := DefaultCreateEcomorph(ctx, e.ToPB(ctx, ecomorph), e.db)
	if err != nil {
		log.Error("Insert Ecomorph:", err)
		return nil, err
	}
	log.Debug("Insert Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphsServetImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*Ecomorph, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	ecomorph, err := DefaultReadEcomorph(ctx, &Ecomorph{Id: request.Id, UserId: userId}, e.db)
	if err != nil {
		log.Error("Get Ecomorph:", err)
		return nil, err
	}
	log.Debug("Get Ecomorph: good")

	return ecomorph, nil
}

func (e EcomorphsServetImpl) DeleteEcomorphById(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteEcomorph(ctx, &Ecomorph{Id: request.Id, UserId: userId}, e.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphsServetImpl) mustEmbedUnimplementedEcomorphServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (e EcomorphsServetImpl) ToPB(ctx context.Context, ecomorph *InputEcomorph) *Ecomorph {
	var id *resource.Identifier

	if ecomorph.Id != nil {
		id = ecomorph.Id
	} else {
		id.ResourceId = pkg.GenerateUUID()
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &Ecomorph{
		Id:          id,
		Title:       ecomorph.Input.Title,
		Description: ecomorph.Input.Description,
		UserId:      userId,
	}
}
