package ecomorph_entity

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest/ecomorph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEcomorphsEntity(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()
	testTable := []struct {
		name     string
		Ecomorph api.InputFormEcomorphsEntity
		expected bool
	}{
		{
			name: "Done",
			Ecomorph: api.InputFormEcomorphsEntity{
				Title:       "Семейство",
				Description: "Ну про вид",
				Ecomorphs:   &api.Ecomorph{Id: ecomorph.CreateEcomorph(ctx, token, client)},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
				mutation insertecomorphsEntity($data: InputFormEcomorphsEntity){
				  ecomorphsEntity{
					insertEcomorphEntity(input:$data){
					  id{
						resourceId
					  }
					  title
					}
				  }
				}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.Ecomorph)

			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteEcomorphsEntity(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetEcomorphEntityById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()
	testTable := []struct {
		name       string
		idEcomorph *api.IdRequest
		expected   bool
	}{
		{
			name:       "GetEcomorphEntity",
			idEcomorph: &api.IdRequest{Id: CreateEcomorphsEntity(ctx, token, client)},
			expected:   true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`

query getEcomorphsEntity($data: ID!){
  ecomorphsEntity{
    getEcomorphEntityByID(id:$data){
      id{
        resourceId
      }
      title
    }
  }
}

			`)
			var respData interface{}
			req.Var("data", testCase.idEcomorph.Id.ResourceId)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteEcomorphsEntity(ctx, *client, testCase.idEcomorph.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateEcomorphEntityById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Ecomorph api.InputEcomorphsEntity
		expected bool
	}{
		{
			name: "GetEcomorphEntity",
			Ecomorph: api.InputEcomorphsEntity{
				Id: CreateEcomorphsEntity(ctx, token, client),
				Input: &api.InputFormEcomorphsEntity{
					Title: "Не семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation update( $data: InputEcomorphsEntity ){
  ecomorphsEntity{
    updateEcomorphEntity(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.Ecomorph)
			req.Var("data", data)
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

func TestGetListEcomorphEntity(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()
	SearchTitle := "Не "
	testTable := []struct {
		name     string
		request  *api.EcomorphsEntityListRequest
		expected bool
	}{
		{
			name: "GetEcomorphsEntity",
			request: &api.EcomorphsEntityListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 10},
			},
			expected: true,
		},
		{
			name: "GetEcomorphsEntity FilterEcomorphsEntity ID",
			request: &api.EcomorphsEntityListRequest{
				Filter: &api.FilterEcomorphsEntity{Id: []*resource.Identifier{CreateEcomorphsEntity(ctx, token, client)}},
			},
			expected: true,
		},
		{
			name: "GetEcomorphsEntity FilterEcomorphsEntity Title",
			request: &api.EcomorphsEntityListRequest{
				Filter: &api.FilterEcomorphsEntity{SearchTitle: &SearchTitle},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getListEcomorphsEntity($data: EcomorphsEntityListRequest){
  ecomorphsEntity{
    getAllEcomorphEntity(pages:$data){
    	page{
        total
        page
      }
      list{
        id{
          resourceId
        }
        title
      }
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.request)
			req.Var("data", data)
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
