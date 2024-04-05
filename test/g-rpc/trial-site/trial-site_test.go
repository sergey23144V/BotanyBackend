package trial_site

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	g_rpc "github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTrialSite(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name      string
		TrialSite *api.InputTrialSiteRequest
		expected  bool
	}{
		{
			name: "Done",
			TrialSite: &api.InputTrialSiteRequest{
				Input: &api.InputFormTrialSiteRequest{
					Title:      "N1",
					CountTypes: 20,
					Rating:     2,
					Covered:    40,
					Plant:      []*api.Plant{{Id: &resource.Identifier{ResourceId: "9e958acb-5d06-8cbb-f204-71dc57529edd"}}, {Id: &resource.Identifier{ResourceId: "a0696b75-3f02-10ec-d307-bc3cb818ef46"}}},
					//Img:        &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
			},
			expected: true,
		},
		{
			name: "Done",
			TrialSite: &api.InputTrialSiteRequest{
				Input: &api.InputFormTrialSiteRequest{
					Title:      "N2",
					CountTypes: 20,
					Rating:     2,
					Covered:    40,
					Plant:      []*api.Plant{{Id: &resource.Identifier{ResourceId: "ace7aaf5-8083-23d2-79e3-0fdd644c00f4"}}, {Id: &resource.Identifier{ResourceId: "71fd04cb-ddc2-2672-9ccb-2a3ca49b4378"}}},
					//Img:        &api.Img{Id: &resource.Identifier{ResourceId: "5622f6d5-9dd1-1567-d198-0ca6a1600c2d"}},
				},
			},
			expected: true,
		},
		//{
		//	name: "Error ",
		//	TrialSite: &api.InputTrialSiteRequest{
		//		Input: &api.InputFormTrialSiteRequest{
		//			Title:      "N2",
		//			CountTypes: 20,
		//			Rating:     2,
		//			Covered:    40,
		//			Plant:      []*api.Plant{{Id: CreatePlant(ctx, *client)}, {Id: CreatePlant(ctx, *client)}, {Id: CreatePlant(ctx, *client)}},
		//		},
		//	},
		//	expected: false,
		//},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.CreateTrialSite(ctx, testCase.TrialSite)
			g_rpc.Log(result)
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
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name     string
		id       *api.IdRequest
		expected bool
	}{
		{
			name:     "GetTrialSite",
			id:       &api.IdRequest{Id: CreateTrialSite(ctx, *client)},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.GetTrialSite(ctx, testCase.id)
			g_rpc.Log(result)
			if testCase.expected {
				err := DeleteTrialSite(ctx, *client, testCase.id.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestUpdateTrialSiteById(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name      string
		TrialSite *api.InputTrialSiteRequest
		expected  bool
	}{
		{
			name: "GetTransect",
			TrialSite: &api.InputTrialSiteRequest{
				Id: CreateTrialSite(ctx, *client),
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
			_, err := client.TrialSite.UpTrialSite(ctx, testCase.TrialSite)
			if testCase.expected {
				err := DeleteTrialSite(ctx, *client, testCase.TrialSite.Id)
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}

func TestGetListTrialSite(t *testing.T) {
	client, ctx := g_rpc.GetClient()

	testTable := []struct {
		name string

		expected bool
	}{
		{
			name:     "GetListTransect",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.GetAllTrialSite(ctx, &api.PagesRequest{Page: 1, Limit: 2})
			g_rpc.Log(result)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
