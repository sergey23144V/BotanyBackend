package type_plant

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/test/graphqlTest/ecomorph-entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesTypePlant(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name      string
		TypePlant api.InputFormTypePlantRequest
		expected  bool
	}{
		{
			name: "Done",
			TypePlant: api.InputFormTypePlantRequest{
				Title:           "Водоросли",
				Subtitle:        "Ну про вид",
				EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, token, client)}},
				Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation insertTypePlant($data: InputFormTypePlantRequest){
  typePlant{
    createTypePlant(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.TypePlant)

			req.Var("data", data)
			req.Header.Set("Authorization", token)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			if testCase.expected {
				//err := DeleteTypePlant(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetTypePlantById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTypePlant",
			id:       &api.IdRequest{Id: CreateTypePlant(ctx, token, client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`

query getEcomorphsEntity($data: ID!){
  typePlant{
    getTypePlant(id:$data){
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
				//err := DeleteTypePlant(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTypePlantById(t *testing.T) {
	client, token := graphqlTest.GetClient()
	ctx := context.Background()

	testTable := []struct {
		name           string
		InputTypePlant api.InputTypePlantRequest
		expected       bool
	}{
		{
			name: "Done",
			InputTypePlant: api.InputTypePlantRequest{
				Id: CreateTypePlant(ctx, token, client),
				Input: &api.InputFormTypePlantRequest{
					Title:           "не Семейство",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, token, client)}},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation update( $data: InputTypePlantRequest ){
  typePlant{
    updateTypePlant(input:$data){
      id{
        resourceId
      }
      title
    }
  }
}
			`)
			var respData interface{}
			data := graphqlTest.StructToMap(testCase.InputTypePlant)
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

func TestGetListTypePlant(t *testing.T) {
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
query getListEcomorphsEntity($data: PagesRequest){
  typePlant{
    getAllTypePlant(pages:$data){
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
