package ecomorph_entity

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	log "github.com/sirupsen/logrus"
)

func CreateEcomorphsEntity(ctx context.Context, client g_rpc.ClientService) *resource.Identifier {
	Ecomorph := &api.InputEcomorphsEntity{
		Input: &api.InputFormEcomorphsEntity{
			Title:        "Не переносящий влагу",
			Ecomorphs:    &api.Ecomorph{Id: &resource.Identifier{ResourceId: "d7ba7908-2ae1-9abd-e1ff-6bed503d80c9"}},
			Description:  "Ну про вид",
			DisplayTable: "MsKs",
			Score:        0,
		},
	}

	result, err := client.Ecomorph_Emtity.InsertEcomorphEntity(ctx, Ecomorph)
	if err != nil {
		log.Error("Не создался Ecomorph")
		return nil
	}
	return result.Id
}

func DeleteEcomorphsEntity(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.Ecomorph_Emtity.DeleteEcomorphEntityByID(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
