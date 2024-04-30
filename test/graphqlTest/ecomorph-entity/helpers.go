package ecomorph_entity

import (
	"context"
	"encoding/json"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest/ecomorph"
	log "github.com/sirupsen/logrus"
)

func CreateEcomorphsEntity(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	EcomorphsEntity := api.InputFormEcomorphsEntity{
		Title:       "Семейство",
		Ecomorphs:   &api.Ecomorph{Id: ecomorph.CreateEcomorph(ctx, token, client)},
		Description: "Ну про вид",
	}

	req := graphql.NewRequest(`
mutation insertecomorphsEntity($data: InputFormEcomorphsEntity){
  ecomorphsEntity{
    insertEcomorphEntity(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}`)

	var respData interface{}
	data := pkg.StructToMap(EcomorphsEntity)

	req.Var("data", data)
	req.Header.Set("Authorization", token)
	err := client.Run(ctx, req, &respData)
	if err != nil {
		log.Error("Не создался EcomorphsEntity")
		return nil
	}
	result, err := json.Marshal(respData)
	if err != nil {
		return nil
	}
	id, err := graphqlTest.JsonPathValue(string(result), "$.ecomorphsEntity.insertEcomorphEntity.id.resourceId")
	if err != nil {
		log.Error("Не создался EcomorphsEntity")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func DeleteEcomorphsEntity(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.Ecomorph_Emtity.DeleteEcomorphEntityByID(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
