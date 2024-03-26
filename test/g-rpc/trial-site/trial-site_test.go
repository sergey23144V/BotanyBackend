package trial_site

import (
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
					Title:      "Трава",
					CountTypes: 20,
					Rating:     2,
					Covered:    4,
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := client.TrialSite.CreateTrialSite(ctx, testCase.TrialSite)
			g_rpc.Log(result)
			if testCase.expected {
				err := DeleteTrialSite(ctx, *client, result.Id)
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
