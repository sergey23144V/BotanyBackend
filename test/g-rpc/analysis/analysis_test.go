package analysis

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAnalysis(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Transect *api.InputCreateAnalysis
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
			Transect: &api.InputCreateAnalysis{
				Title:        "Хз",
				TypeAnalysis: api.TypeAnalysis_TypeAnalysisPlant,
				Transect:     &api.Transect{Id: &resource.Identifier{ResourceId: "3d9f3076-4b23-33e0-40f5-0e7bde4e5a56"}},
				Ecomorph:     []*api.Ecomorph{{Id: &resource.Identifier{ResourceId: "d7ba7908-2ae1-9abd-e1ff-6bed503d80c9"}}, {Id: &resource.Identifier{ResourceId: "c15c4ebd-d583-6c3d-81fe-af286a2b72e1"}}},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Analysis.CreatAnalysis(ctx, testCase.Transect)
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
