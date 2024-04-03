package ecomorph

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEcomorph(t *testing.T) {
	client, token := graphqlTest.GetClient()

	testTable := []struct {
		name     string
		Ecomorph *api.InputFormEcomorph
		expected bool
	}{
		{
			name: "Done",
			Ecomorph: &api.InputFormEcomorph{
				Title:       "Семейство",
				Description: "Ну про вид",
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation insertEcomorph($data: InputFormEcomorph){
  ecomorph{
    insertEcomorph(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			ctx := context.Background()
			req.Var("data", testCase.Ecomorph)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteEcomorphById(ctx, client.Ecomorph, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetEcomorphById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()
	testTable := []struct {
		name       string
		idEcomorph *api.IdRequest
		expected   bool
	}{
		{
			name:       "GetEcomorph",
			idEcomorph: &api.IdRequest{Id: CreateEcomorph(ctx, token, client)},
			expected:   true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
query getEcomorphById($data: ID!){
  ecomorph{
    getEcomorphById(id:$data){
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
				//err := DeleteEcomorphById(ctx, client.Ecomorph, testCase.idEcomorph.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateEcomorphById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		Ecomorph *api.InputEcomorph
		expected bool
	}{
		{
			name: "GetEcomorph",
			Ecomorph: &api.InputEcomorph{
				Id: CreateEcomorph(ctx, token, client),
				Payload: &api.InputFormEcomorph{
					Title: "Не семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation update($id: String! , $data:InputFormEcomorph! ){
  ecomorph{
    updateEcomorph(input:{id:{resourceId:$id} payload: $data} ){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			req.Var("data", testCase.Ecomorph.Payload)
			req.Var("id", testCase.Ecomorph.Id.ResourceId)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListEcomorph(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		page     *api.PagesRequest
		expected bool
	}{
		{
			name: "GetEcomorph",
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
query getListEcomorph($data: PagesRequest){
  ecomorph{
    getListEcomorph(pages:$data){
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
