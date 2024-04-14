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
					Title:           "Терекон",
					Square:          20,
					SquareTrialSite: 1,
					TrialSite:       []*api.TrialSite{{Id: &resource.Identifier{ResourceId: "51d485bb-ec34-284b-d8a1-13e2f1608669"}}, {Id: &resource.Identifier{ResourceId: "380c728b-5b2c-1f5b-c68c-d164ddc7d2a5"}}},
					//Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
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
				//err := DeleteTransect(ctx, *client, result.Id)
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
					TrialSite: []*api.TrialSite{{Id: &resource.Identifier{ResourceId: "51d485bb-ec34-284b-d8a1-13e2f1608669"}}, {Id: &resource.Identifier{ResourceId: "380c728b-5b2c-1f5b-c68c-d164ddc7d2a5"}},
						{Id: &resource.Identifier{ResourceId: "ba1f1a30-9e78-c0f1-7cf9-1c9eba460ccf"}}, {Id: &resource.Identifier{ResourceId: "59a28d5f-10a4-44d4-fcf4-6c768e511f7c"}}},
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

func TestAddTrialSiteToTransect(t *testing.T) {
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
					TrialSite: []*api.TrialSite{{Id: trial_site.CreateTrialSite(ctx, *client)},
						{Id: trial_site.CreateTrialSite(ctx, *client)},
					},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r, err := client.Transect.AddTrialSiteToTransect(ctx, testCase.Transect)
			g_rpc.Log(r)
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
	SearchTitle := "Терекон"
	testTable := []struct {
		name     string
		request  *api.TransectListRequest
		expected bool
	}{
		{
			name: "GetListTransect",
			request: &api.TransectListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 2},
			},
			expected: true,
		},
		{
			name:     "GetListTransect all",
			request:  &api.TransectListRequest{},
			expected: true,
		},
		{
			name: "GetListTransect Id",
			request: &api.TransectListRequest{
				Filter: &api.FilterTransect{Id: []*resource.Identifier{CreateTransect(ctx, *client)}},
			},
			expected: true,
		},
		{
			name: "GetListTransect SearchTitle",
			request: &api.TransectListRequest{
				Filter: &api.FilterTransect{SearchTitle: SearchTitle},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Transect.GetAllTransect(ctx, testCase.request)
			g_rpc.Log(result)
			g_rpc.Log(len(result.List))
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
