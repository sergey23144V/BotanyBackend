package trial_site

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	type_plant "github.com/sergey23144V/BotanyBackend/test/graphqlTest/type-plant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePlant(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Plant    *api.InputFormPlant
		expected bool
	}{
		{
			name: "Done",
			Plant: &api.InputFormPlant{
				TypePlantId: &api.TypePlant{Id: type_plant.CreateTypePlant(ctx, token, client)},
				Count:       2,
				Coverage:    40,
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
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
}

			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.Plant)

			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
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
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetPlant",
			id:       &api.IdRequest{Id: CreatePlant(ctx, token, client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`

query getPlant($data: ID!){
  trialSite{
    getPlant(id:$data){
      id{
        resourceId
      }
      count
    }
  }
}

			`)
			var respData interface{}
			req.Var("data", testCase.id.Id.ResourceId)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeletePlant(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdatePlantById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Plant    *api.InputPlantRequest
		expected bool
	}{
		{
			name: "GetTransect",
			Plant: &api.InputPlantRequest{
				Id: CreatePlant(ctx, token, client),
				Input: &api.InputFormPlant{
					TypePlantId: &api.TypePlant{Id: type_plant.CreateTypePlant(ctx, token, client)},
					Count:       2,
					Coverage:    4,
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation updatePlant( $data: InputPlantRequest ){
  trialSite{
    updatePlant(input:$data){
      id{
        resourceId
      }
      count
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.Plant)
			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeletePlant(ctx, *client, testCase.Plant.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListPlant(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		page     *api.PagesRequest
		expected bool
	}{
		{
			name: "GetListEcomorphEntity",
			page: &api.PagesRequest{
				Limit: 2,
				Page:  1,
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getListPlant($data: PagesRequest){
  trialSite{
    getAllPlant(pages:$data){
    	page{
        total
        page
      }
      list{
        id{
          resourceId
        }
        count
      }
    }
  }
}
			`)
			var respData interface{}
			req.Var("data", testCase.page)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
