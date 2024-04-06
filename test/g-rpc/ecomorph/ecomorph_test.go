package ecomorph

import (
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEcomorph(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Ecomorph *api.InputEcomorph
		expected bool
	}{
		{
			name: "Done",
			Ecomorph: &api.InputEcomorph{
				Payload: &api.InputFormEcomorph{
					Title:       "Семейства",
					Description: "Ну про вид",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			result, err := client.Ecomorph.InsertEcomorph(ctx, testCase.Ecomorph)
			g_rpc.Log(result)
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
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name       string
		idEcomorph *api.IdRequest
		expected   bool
	}{
		{
			name:       "GetEcomorph",
			idEcomorph: &api.IdRequest{Id: CreateEcomorph(ctx, client.Ecomorph)},
			expected:   true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.Ecomorph.GetEcomorphById(ctx, testCase.idEcomorph)
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
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name       string
		idEcomorph *api.InputEcomorph
		expected   bool
	}{
		{
			name: "GetEcomorph",
			idEcomorph: &api.InputEcomorph{
				Id: CreateEcomorph(ctx, client.Ecomorph),
				Payload: &api.InputFormEcomorph{
					Title: "Не семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.Ecomorph.UpdateEcomorph(ctx, testCase.idEcomorph)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListEcomorph(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		request  *api.EcomorphListRequest
		expected bool
	}{
		{
			name: "GetEcomorph",
			request: &api.EcomorphListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 2},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r, err := client.Ecomorph.GetListEcomorph(ctx, testCase.request)
			g_rpc.Log(r.List)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
