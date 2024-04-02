package type_plant

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/test/g-rpc/ecomorph-entity"
	log "github.com/sirupsen/logrus"
)

func CreateTypePlant(ctx context.Context, client g_rpc.ClientService) *resource.Identifier {
	Ecomorph := &api.InputTypePlantRequest{
		Input: &api.InputFormTypePlantRequest{
			Title:           "Трава",
			EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, client)}},
			Subtitle:        "Ну про вид",
			Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
		},
	}

	result, err := client.TypePlant.CreateTypePlant(ctx, Ecomorph)
	if err != nil {
		log.Error("Не создался TypePlant")
		return nil
	}
	return result.Id
}

func DeleteTypePlant(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.TypePlant.DeleteTypePlant(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления TypePlant произашло успешно")
	return nil
}
