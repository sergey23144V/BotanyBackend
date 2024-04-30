package trial_site

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTrialSite(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name      string
		TrialSite *api.InputFormTrialSiteRequest
		expected  bool
	}{
		{
			name: "Done",
			TrialSite: &api.InputFormTrialSiteRequest{
				Title:      "N2",
				CountTypes: 20,
				Rating:     2,
				Covered:    40,
				Plant:      []*api.Plant{{Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}},
				Img:        &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
			},
			expected: true,
		},
		{
			name: "Error ",
			TrialSite: &api.InputFormTrialSiteRequest{
				Title:      "N2",
				CountTypes: 20,
				Rating:     2,
				Covered:    40,
				Plant:      []*api.Plant{{Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}, {Id: CreatePlant(ctx, token, client)}},
			},
			expected: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
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
			data := pkg.StructToMap(testCase.TrialSite)

			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteTrialSite(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetTrialSiteById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTrialSite",
			id:       &api.IdRequest{Id: CreateTrialSite(ctx, token, client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getTrialSite($data: ID!){
  trialSite{
    getTrialSite(id:$data){
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
				//err := DeleteTrialSite(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTrialSiteById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name      string
		TrialSite *api.InputTrialSiteRequest
		expected  bool
	}{
		{
			name: "GetTransect",
			TrialSite: &api.InputTrialSiteRequest{
				Id: CreateTrialSite(ctx, token, client),
				Input: &api.InputFormTrialSiteRequest{
					Title:      "не Трава",
					CountTypes: 27,
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation upTrialSite( $data: InputTrialSiteRequest ){
  trialSite{
    upTrialSite(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			data := pkg.StructToMap(testCase.TrialSite)
			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteTrialSite(ctx, *client, testCase.TrialSite.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListTrialSite(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()
	SearchTitle := "не "
	testTable := []struct {
		name     string
		request  *api.TrialSiteListRequest
		expected bool
	}{
		{
			name: "GetEcomorphsEntity",
			request: &api.TrialSiteListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 10},
			},
			expected: true,
		},
		{
			name: "GetEcomorphsEntity FilterEcomorphsEntity ID",
			request: &api.TrialSiteListRequest{
				Filter: &api.FilterTrialSite{Id: []*resource.Identifier{CreateTrialSite(ctx, token, client)}},
			},
			expected: true,
		},
		{
			name: "GetEcomorphsEntity FilterEcomorphsEntity Title",
			request: &api.TrialSiteListRequest{
				Filter: &api.FilterTrialSite{SearchTitle: SearchTitle},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getAllTrialSite($data: TrialSiteListRequest){
  trialSite{
    getAllTrialSite(pages:$data){
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
			data := pkg.StructToMap(testCase.request)
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
