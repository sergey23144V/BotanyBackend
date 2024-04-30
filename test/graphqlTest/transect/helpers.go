package transect

import (
	"context"
	"encoding/json"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	trial_site "github.com/sergey23144V/BotanyBackend/test/graphqlTest/trial-site"
	log "github.com/sirupsen/logrus"
)

func CreateTransect(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	Transect := api.InputFormTransectRequest{
		Title:           "Трава",
		Square:          20,
		SquareTrialSite: 2,
		TrialSite:       []*api.TrialSite{{Id: trial_site.CreateTrialSite(ctx, token, client)}},
		Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
	}

	req := graphql.NewRequest(`
mutation createTrialSite($data:InputFormTransectRequest ){
  transect{
    createTransect(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
`)
	var respData interface{}
	data := pkg.StructToMap(Transect)
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
	id, err := graphqlTest.JsonPathValue(string(result), "$.transect.createTransect.id.resourceId")
	if err != nil {
		log.Error("Не создался typePlant")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func DeleteTransect(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.Transect.DeleteTransect(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
