package analysis

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
		Transect *api.InputCreateAnalysis
		expected bool
	}{
		{
			name: "Done",
			Transect: &api.InputCreateAnalysis{
				Title:        "Хз",
				TypeAnalysis: api.TypeAnalysis_TypeAnalysisTransect,
				Transect:     &api.Transect{Id: &resource.Identifier{ResourceId: "fdb18250-bdcd-6671-b119-b5cc3f427c60"}},
				Ecomorph:     []*api.Ecomorph{{Id: &resource.Identifier{ResourceId: "d7ba7908-2ae1-9abd-e1ff-6bed503d80c9"}}},
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
