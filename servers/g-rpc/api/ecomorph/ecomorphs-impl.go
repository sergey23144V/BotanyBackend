package ecomorph

import (
	context "context"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

// EcomorphsServetImpl
type EcomorphsServetImpl struct {
	db *gorm.DB
}

func (e EcomorphsServetImpl) InsertEcomorph(ctx context.Context, ecomorph *Ecomorph) (*Ecomorph, error) {
	return DefaultCreateEcomorph(ctx, ecomorph, e.db)
}

func (e EcomorphsServetImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*Ecomorph, error) {
	return DefaultReadEcomorph(ctx, &Ecomorph{Id: request.Id}, e.db)
}

func (e EcomorphsServetImpl) UpEcomorphById(ctx context.Context, ecomorph *Ecomorph) (*Ecomorph, error) {
	return DefaultStrictUpdateEcomorph(ctx, ecomorph, e.db)
}

func (e EcomorphsServetImpl) DeleteEcomorphById(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteEcomorph(ctx, &Ecomorph{Id: request.Id}, e.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphsServetImpl) GetListEcomorphById(ctx context.Context, request *api.EmptyRequest) (*EcomorphsList, error) {
	list, err := DefaultListEcomorph(ctx, e.db)
	if err != nil {
		return nil, err
	}
	return &EcomorphsList{Ecomorph: list}, nil
}

func (e EcomorphsServetImpl) mustEmbedUnimplementedEcomorphServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewEcomorphsServetImplImpl(db *gorm.DB) EcomorphsServetImpl {
	return EcomorphsServetImpl{db}
}
