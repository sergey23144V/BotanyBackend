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
			name: "Подорожник",
			Ecomorph: &api.InputTypePlantRequest{
				Input: &api.InputFormTypePlantRequest{
					Title:           "Подорожник",
					Subtitle:        "Ну про вид",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: &resource.Identifier{ResourceId: "3b087a90-573b-d64d-16b9-d0ad730dc891"}}, {Id: &resource.Identifier{ResourceId: "5aea8c3e-6bc0-9f9b-716e-d71b7f70565d"}}},
					//Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
			},
			expected: true,
		},
		{
			name: "Яблоня",
			Ecomorph: &api.InputTypePlantRequest{
				Input: &api.InputFormTypePlantRequest{
					Title:           "Яблоня",
					Subtitle:        "Ну про вид",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: &resource.Identifier{ResourceId: "09f59b48-64bf-79ed-f980-8d675ef182c6"}}, {Id: &resource.Identifier{ResourceId: "380fcbb8-6faf-4bf6-b125-36617daef527"}}},
					//Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
			},
			expected: true,
		},
		{
			name: "Акация",
			Ecomorph: &api.InputTypePlantRequest{
				Input: &api.InputFormTypePlantRequest{
					Title:           "Акация",
					Subtitle:        "Ну про вид",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: &resource.Identifier{ResourceId: "09f59b48-64bf-79ed-f980-8d675ef182c6"}}, {Id: &resource.Identifier{ResourceId: "5aea8c3e-6bc0-9f9b-716e-d71b7f70565d"}}},
					//Img:             &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
				Publicly: true,
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			result, err := client.TypePlant.CreateTypePlant(ctx, testCase.Ecomorph)
			g_rpc.Log(result)
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
			result, err := client.TypePlant.GetTypePlant(ctx, testCase.id)
			g_rpc.Log(result)
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
				Id: CreateTypePlant(ctx, *client),
				Input: &api.InputFormTypePlantRequest{
					Title:           "не Семейство",
					EcomorphsEntity: []*api.EcomorphsEntity{{Id: ecomorph_entity.CreateEcomorphsEntity(ctx, *client)}},
				},
			},
			expected: true,
		},
		{
			name: "Done",
			InputTypePlant: &api.InputTypePlantRequest{
				Id: &resource.Identifier{ResourceId: "92bd51e1-33f2-0f8b-f27b-4f52bd06961a"},
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
			result, err := client.TypePlant.AddEcomorphsEntityToTypePlant(ctx, testCase.InputTypePlant)
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
		name     string
		request  *api.TypePlantListRequest
		expected bool
	}{
		{
			name: "GetListTypePlant",
			request: &api.TypePlantListRequest{
				Page: &api.PagesRequest{Page: 1, Limit: 2},
			},
			expected: true,
		},
		{
			name:     "GetListTypePlant",
			request:  &api.TypePlantListRequest{},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TypePlant.GetAllTypePlant(ctx, testCase.request)
			g_rpc.Log(result)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
