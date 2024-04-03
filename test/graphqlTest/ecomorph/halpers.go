package ecomorph

import (
	"encoding/json"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func CreateEcomorph(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	Ecomorph := &api.InputFormEcomorph{
		Title:       "Семейство",
		Description: "Ну про вид",
	}

	req := graphql.NewRequest(`
	mutation insertEcomorph($data: InputFormEcomorph){
  ecomorph{
    insertEcomorph(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}`)

	var respData interface{}
	req.Var("data", Ecomorph)
	req.Header.Set("Authorization", token)
	err := client.Run(ctx, req, &respData)
	if err != nil {
		log.Error("Не создался Ecomorph")
		return nil
	}
	data, err := json.Marshal(respData)
	if err != nil {
		return nil
	}
	id, err := graphqlTest.JsonPathValue(string(data), "$.ecomorph.insertEcomorph.id.resourceId")
	if err != nil {
		log.Error("Не создался Ecomorph")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func DeleteEcomorphById(ctx context.Context, client api.EcomorphServiceClient, id *resource.Identifier) error {
	_, err := client.DeleteEcomorphById(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
