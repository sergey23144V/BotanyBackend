package transect

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	trial_site "github.com/sergey23144V/BotanyBackend/test/graphqlTest/trial-site"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransect(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Transect api.InputFormTransectRequest
		expected bool
	}{
		{
			name: "Done",
			Transect: api.InputFormTransectRequest{
				Title:           "Семейство",
				Square:          20,
				SquareTrialSite: 1,
				TrialSite:       []*api.TrialSite{{Id: trial_site.CreateTrialSite(ctx, token, client)}, {Id: trial_site.CreateTrialSite(ctx, token, client)}, {Id: trial_site.CreateTrialSite(ctx, token, client)}},
				Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
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
			data := graphqlTest.StructToMap(testCase.Transect)

			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
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
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTransect",
			id:       &api.IdRequest{Id: CreateTransect(ctx, token, client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getTrialSite($data: ID!){
  transect{
    getTransect(id:$data){
      id{
        resourceId
      }
      title
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
				//err := DeleteTransect(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTransectById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Transect *api.InputTransectRequest
		expected bool
	}{
		{
			name: "GetTransect",
			Transect: &api.InputTransectRequest{
				Id: CreateTransect(ctx, token, client),
				Input: &api.InputFormTransectRequest{
					Title: "не Семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation upTrialSite( $data: InputTransectRequest ){
  transect{
    upTransect(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.Transect)
			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteTransect(ctx, *client, testCase.Transect.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListTransect(t *testing.T) {
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
			query getAllTrialSite($data: PagesRequest){
			  transect{
				getAllTransect(pages:$data){
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
