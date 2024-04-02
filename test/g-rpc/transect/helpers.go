package transect

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	trial_site "github.com/sergey23144V/BotanyBackend/test/g-rpc/trial-site"
	log "github.com/sirupsen/logrus"
)

func CreateTransect(ctx context.Context, client g_rpc.ClientService) *resource.Identifier {
	Transect := &api.InputTransectRequest{
		Input: &api.InputFormTransectRequest{
			Title:           "Трава",
			Square:          20,
			SquareTrialSite: 2,
			TrialSite:       []*api.TrialSite{{Id: trial_site.CreateTrialSite(ctx, client)}},
			Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
		},
	}

	result, err := client.Transect.CreateTransect(ctx, Transect)
	if err != nil {
		log.Error("Не создался Ecomorph")
		return nil
	}
	return result.Id
}

func DeleteTransect(ctx context.Context, client g_rpc.ClientService, id *resource.Identifier) error {
	_, err := client.Transect.DeleteTransect(ctx, &api.IdRequest{Id: id})
	if err != nil {
		return err
	}
	log.Debug("Удаления Ecomorph произашло успешно")
	return nil
}
