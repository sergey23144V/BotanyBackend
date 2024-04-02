package transect

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	trial_site "github.com/sergey23144V/BotanyBackend/test/g-rpc/trial-site"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransect(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Transect *api.InputTransectRequest
		expected bool
	}{
		{
			name: "Done",
			Transect: &api.InputTransectRequest{
				Input: &api.InputFormTransectRequest{
					Title:           "Семейство",
					Square:          20,
					SquareTrialSite: 1,
					TrialSite:       []*api.TrialSite{{Id: trial_site.CreateTrialSite(ctx, *client)}},
					Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Transect.CreateTransect(ctx, testCase.Transect)
			g_rpc.Log(result)
			if testCase.expected {
				err := DeleteTransect(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetTransectById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTransect",
			id:       &api.IdRequest{Id: CreateTransect(ctx, *client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Transect.GetTransect(ctx, testCase.id)
			g_rpc.Log(result)
			if testCase.expected {
				err := DeleteTransect(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTransectById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Transect *api.InputTransectRequest
		expected bool
	}{
		{
			name: "GetTransect",
			Transect: &api.InputTransectRequest{
				Id: CreateTransect(ctx, *client),
				Input: &api.InputFormTransectRequest{
					Title: "не Семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.Transect.UpTransect(ctx, testCase.Transect)
			if testCase.expected {
				err := DeleteTransect(ctx, *client, testCase.Transect.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListTransect(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name string

		expected bool
	}{
		{
			name:     "GetListTransect",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Transect.GetAllTransect(ctx, &api.PagesRequest{Page: 1, Limit: 2})
			g_rpc.Log(result)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
