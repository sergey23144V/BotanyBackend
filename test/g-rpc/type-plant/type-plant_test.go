package type_plant

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/test/g-rpc/ecomorph-entity"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesTypePlant(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		Ecomorph *api.InputTypePlantRequest
		expected bool
	}{
		{
			name: "Done",
			Ecomorph: &api.InputTypePlantRequest{
				Input: &api.InputFormTypePlantRequest{
					Title:           "Водоросли",
					Subtitle:        "Ну про вид",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)}},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			result, err := client.TypePlant.CreateTypePlant(ctx, testCase.Ecomorph)
			if testCase.expected {
				err := DeleteTypePlant(ctx, *client, result.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetTypePlantById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTypePlant",
			id:       &api.IdRequest{Id: CreateTypePlant(ctx, *client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.TypePlant.GetTypePlant(ctx, testCase.id)
			if testCase.expected {
				err := DeleteTypePlant(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTypePlantById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name           string
		InputTypePlant *api.InputTypePlantRequest
		expected       bool
	}{
		{
			name: "Done",
			InputTypePlant: &api.InputTypePlantRequest{
				Id: &resource.Identifier{ResourceId: "cc518ea6-c38d-481c-ec0e-2e2c2d91dbb0"},
				Input: &api.InputFormTypePlantRequest{
					Title:           "не Семейство",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)}},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TypePlant.UpdateTypePlant(ctx, testCase.InputTypePlant)
			log.Println(result.EcomorphsEntity)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestAddEcomorphsEntityById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name           string
		InputTypePlant *api.InputTypePlant_EcomorphsEntityRequest
		expected       bool
	}{
		{
			name: "Done",
			InputTypePlant: &api.InputTypePlant_EcomorphsEntityRequest{
				Id: CreateTypePlant(ctx, *client),
				EcomorphsEntity: []*api.EcomorphsEntity{
					{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)},
					{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)},
					{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TypePlant.AddEcomorphsEntity(ctx, testCase.InputTypePlant)
			g_rpc.Log(result.EcomorphsEntity)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListTypePlant(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name string

		expected bool
	}{
		{
			name:     "GetListTypePlant",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := client.TypePlant.GetAllTypePlant(ctx, &api.PagesRequest{Page: 1, Limit: 2})
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
