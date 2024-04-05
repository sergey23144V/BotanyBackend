package ecomorph

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func CreateEcomorph(ctx context.Context, client api.EcomorphServiceClient) *resource.Identifier {
	Ecomorph := &api.InputEcomorph{
		Payload: &api.InputFormEcomorph{
			Title:       "Гидроморф",
			Description: "Определяет вид растения",
		},
	}

	result, err := client.InsertEcomorph(ctx, Ecomorph)
	if err != nil {
		log.Error("Не создался Ecomorph")
		return nil
	}
	return result.Id
}

func DeleteEcomorphById(ctx context.Context, client api.EcomorphServiceClient, id *resource.Identifier) error {
	_, err := client.DeleteEcomorphById(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
