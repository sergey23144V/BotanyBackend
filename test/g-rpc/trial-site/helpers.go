package trial_site

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	type_plant "github.com/sergey23144V/BotanyBackend/test/g-rpc/type-plant"
	log "github.com/sirupsen/logrus"
)

func CreateTrialSite(ctx context.Context, client g_rpc.ClientService) *resource.Identifier {
	Transect := &api.InputTrialSiteRequest{
		Input: &api.InputFormTrialSiteRequest{
			Title:      "Трава",
			CountTypes: 20,
			Rating:     2,
			Covered:    4,
			Plant:      []*api.Plant{{Id: CreatePlant(ctx, client)}, {Id: CreatePlant(ctx, client)}},
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

func CreatePlant(ctx context.Context, client g_rpc.ClientService) *resource.Identifier {
	Plant := &api.InputPlantRequest{
		Input: &api.InputFormPlant{
			TypePlantId: &api.TypePlant{Id: type_plant.CreateTypePlant(ctx, client)},
			Count:       2,
			Coverage:    40,
		},
	}

	result, err := client.TrialSite.CreatePlant(ctx, Plant)
	if err != nil {
		log.Error("Не создался Plant")
		return nil
	}
	return result.Id
}

func DeletePlant(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.TrialSite.DeletePlant(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления TrialSite произашло успешно")
	return nil
}
