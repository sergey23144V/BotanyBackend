package analysis

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

func TestCreateTransect(t *testing.T) {
	client, token := graphqlTest.GetClient()

	testTable := []struct {
		name     string
		analysis *api.InputCreateAnalysis
		expected bool
	}{
		//{
		//	name: "Done",
		//	Transect: &api.InputCreateAnalysis{
		//		Title:        "Хз",
		//		TypeAnalysis: api.TypeAnalysis_TypeAnalysisPlant,
		//		Transect:     &api.Transect{Id: &resource.Identifier{ResourceId: "fdb18250-bdcd-6671-b119-b5cc3f427c60"}},
		//		Ecomorph:     []*api.Ecomorph{{Id: &resource.Identifier{ResourceId: "d7ba7908-2ae1-9abd-e1ff-6bed503d80c9"}}, {Id: &resource.Identifier{ResourceId: "c15c4ebd-d583-6c3d-81fe-af286a2b72e1"}}},
		//	},
		//	expected: true,
		//},
		{
			name: "Done",
			analysis: &api.InputCreateAnalysis{
				Title:        "Хз",
				TypeAnalysis: api.TypeAnalysis_TypeAnalysisTransect,
				Transect:     &api.Transect{Id: &resource.Identifier{ResourceId: "fdb18250-bdcd-6671-b119-b5cc3f427c60"}},
				Ecomorph:     []*api.Ecomorph{{Id: &resource.Identifier{ResourceId: "d7ba7908-2ae1-9abd-e1ff-6bed503d80c9"}}, {Id: &resource.Identifier{ResourceId: "c15c4ebd-d583-6c3d-81fe-af286a2b72e1"}}},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req := graphql.NewRequest(`
mutation creatAnalysis($data:InputCreateAnalysis!) {
  analysis{
    creatAnalysis(input:$data){
      path
    }
  }
}
			`)
			var respData interface{}
			ctx := context.Background()
			data := pkg.StructToMap(testCase.analysis)
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
