package trial_site

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	type_plant "github.com/sergey23144V/BotanyBackend/test/g-rpc/type-plant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePlant(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Plant    *api.InputPlantRequest
		expected bool
	}{
		{
			name: "Акация",
			Plant: &api.InputPlantRequest{
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: &resource.Identifier{ResourceId: "cafe5dcf-5d73-8c00-bc32-c99c31cd7a4b"}},
					Count:       5,
					Coverage:    40,
				},
			},
			expected: true,
		},
		{
			name: "Яблоня",
			Plant: &api.InputPlantRequest{
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: &resource.Identifier{ResourceId: "0f6f6bcb-33ff-2bd2-0986-cf3039053419"}},
					Count:       12,
					Coverage:    50,
				},
			},
			expected: true,
		},
		{
			name: "Подорожник",
			Plant: &api.InputPlantRequest{
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: &resource.Identifier{ResourceId: "cd9342da-bc9f-d35c-db56-2b664905f682"}},
					Count:       15,
					Coverage:    20,
				},
			},
			expected: true,
		},
		{
			name: "Водоросли",
			Plant: &api.InputPlantRequest{
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: &resource.Identifier{ResourceId: "67946220-a2a9-d1f6-a48d-63427a6e9e22"}},
					Count:       10,
					Coverage:    10,
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.CreatePlant(ctx, testCase.Plant)
			g_rpc.Log(result)
			if testCase.expected {
				//err := DeletePlant(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetPlantById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetPlant",
			id:       &api.IdRequest{Id: CreatePlant(ctx, *client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.GetPlant(ctx, testCase.id)
			g_rpc.Log(result)
			if testCase.expected {
				err := DeletePlant(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdatePlantById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Plant    *api.InputPlantRequest
		expected bool
	}{
		{
			name: "GetTransect",
			Plant: &api.InputPlantRequest{
				Id: CreatePlant(ctx, *client),
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: type_plant.CreateTypePlant(ctx, *client)},
					Count:       2,
					Coverage:    4,
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.TrialSite.UpdatePlant(ctx, testCase.Plant)
			if testCase.expected {
				err := DeletePlant(ctx, *client, testCase.Plant.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListPlant(t *testing.T) {
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
			result, err := client.TrialSite.GetAllPlant(ctx, &api.PagesRequest{Page: 1, Limit: 2})
			g_rpc.Log(result)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
