package ecomorph_entity

import (
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/g-rpc/ecomorph"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateEcomorphsEntity(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Ecomorph *api.InputEcomorphsEntity
		expected bool
	}{
		{
			name: "Done",
			Ecomorph: &api.InputEcomorphsEntity{
				Input: &api.InputFormEcomorphsEntity{
					Title:       "Семейство",
					Description: "Ну про вид",
					Ecomorphs:   &api.Ecomorph{Id: ecomorph.CreateEcomorph(ctx, client.Ecomorph)},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Ecomorph_Emtity.InsertEcomorphEntity(ctx, testCase.Ecomorph)
			log.Println(result.Ecomorphs.Title)
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
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name       string
		idEcomorph *api.IdRequest
		expected   bool
	}{
		{
			name:       "GetEcomorphEntity",
			idEcomorph: &api.IdRequest{Id: CreateEcomorphsEntity(ctx, *client)},
			expected:   true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.Ecomorph_Emtity.GetEcomorphEntityByID(ctx, testCase.idEcomorph)
			g_rpc.Log(result.Ecomorphs.Title)
			if testCase.expected {
				err := DeleteEcomorphsEntity(ctx, *client, testCase.idEcomorph.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateEcomorphEntityById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name       string
		idEcomorph *api.InputEcomorphsEntity
		expected   bool
	}{
		{
			name: "GetEcomorphEntity",
			idEcomorph: &api.InputEcomorphsEntity{
				Id: CreateEcomorphsEntity(ctx, *client),
				Input: &api.InputFormEcomorphsEntity{
					Title: "Не семейство",
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.Ecomorph_Emtity.UpdateEcomorphEntity(ctx, testCase.idEcomorph)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListEcomorphEntity(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		request  *api.EcomorphsEntityListRequest
		expected bool
	}{
		{
			name: "GetListEcomorphEntity",
			request: &api.EcomorphsEntityListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 2},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.Ecomorph_Emtity.GetAllEcomorphEntity(ctx, testCase.request)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
