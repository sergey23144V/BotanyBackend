package transect

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
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
				Id: &resource.Identifier{ResourceId: "fdb18250-bdcd-6671-b119-b5cc3f427c60"},
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
