package type_plant

import (
	"context"
	"encoding/json"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/test/graphqlTest/ecomorph-entity"
	log "github.com/sirupsen/logrus"
)

func CreateTypePlant(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	typePlant := api.InputFormTypePlantRequest{
		Title:           "Трава",
		EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, token, client)}},
		Subtitle:        "Ну про вид",
		Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
	}

	req := graphql.NewRequest(`
mutation insertTypePlant($data: InputFormTypePlantRequest){
  typePlant{
    createTypePlant(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}`)

	var respData interface{}
	data := graphqlTest.StructToMap(typePlant)

	req.Var("data", data)
	req.Header.Set("Authorization", token)
	err := client.Run(ctx, req, &respData)
	if err != nil {
		log.Error("Не создался typePlant")
		return nil
	}
	result, err := json.Marshal(respData)
	if err != nil {
		return nil
	}
	id, err := graphqlTest.JsonPathValue(string(result), "$.typePlant.createTypePlant.id.resourceId")
	if err != nil {
		log.Error("Не создался typePlant")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func DeleteTypePlant(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.TypePlant.DeleteTypePlant(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления TypePlant произашло успешно")
	return nil
}
