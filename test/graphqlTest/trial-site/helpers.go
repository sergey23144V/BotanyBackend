package trial_site

import (
	"context"
	"encoding/json"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	type_plant "github.com/sergey23144V/BotanyBackend/test/graphqlTest/type-plant"
	log "github.com/sirupsen/logrus"
)

func CreateTrialSite(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	plants := &api.Plant{Id: CreatePlant(ctx, token, client)}
	Transect := api.InputFormTrialSiteRequest{
		Title:   "Трава",
		Rating:  2,
		Covered: 4,
		Plant:   []*api.Plant{plants, plants, {Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}},
		Img:     &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
	}

	req := graphql.NewRequest(`
mutation createTrialSite($data: InputFormTrialSiteRequest){
  trialSite{
    createTrialSite(input:$data){
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
	id, err := graphqlTest.JsonPathValue(string(result), "$.trialSite.createTrialSite.id.resourceId")
	if err != nil {
		log.Error("Не создался typePlant")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func CreateTrialSiteForTransect(ctx context.Context, client g_rpc.ClientService, plants []*api.Plant) *resource.Identifier {
	Transect := &api.InputTrialSiteRequest{
		Input: &api.InputFormTrialSiteRequest{
			Title:      "Трава",
			CountTypes: 20,
			Rating:     2,
			Covered:    4,
			Plant:      plants,
			Img:        &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
		},
	}

	result, err := client.TrialSite.CreateTrialSite(ctx, Transect)
	if err != nil {
		log.Error("Не создался TrialSite")
		return nil
	}
	return result.Id
}

func DeleteTrialSite(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.TrialSite.DeleteTrialSite(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления TrialSite произашло успешно")
	return nil
}

func CreatePlant(ctx context.Context, token string, client *graphql.Client) *resource.Identifier {
	Plant := &api.InputFormPlant{
		TypePlantId: &api.TypePlant{Id: type_plant.CreateTypePlant(ctx, token, client)},
		Count:       2,
		Coverage:    20,
	}

	req := graphql.NewRequest(`
mutation insertPlant($data: InputFormPlant){
  trialSite{
    createPlant(input:$data){
      id{
        resourceId
      }
      coverage
    }
  }
}`)
	var respData interface{}
	data := pkg.StructToMap(Plant)
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
	id, err := graphqlTest.JsonPathValue(string(result), "$.trialSite.createPlant.id.resourceId")
	if err != nil {
		log.Error("Не создался typePlant")
		return nil
	}
	return &resource.Identifier{ResourceId: id.(string)}
}

func DeletePlant(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.TrialSite.DeletePlant(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления TrialSite произашло успешно")
	return nil
}
