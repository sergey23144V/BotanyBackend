package auth

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/sergey23144V/BotanyBackend/test/graphqlTest"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorisation(t *testing.T) {
	client, _ := graphqlTest.GetClient()

	testTable := []struct {
		name     string
		user     *api.SignInUserInput
		expected bool
	}{
		{
			name: "Done",
			user: &api.SignInUserInput{
				Email:    "serg2",
				Password: "Sergey2222",
			},
			expected: true,
		},
		{
			name: "Error",
			user: &api.SignInUserInput{
				Email:    "sergeykalinin@gmail.con",
				Password: "Serg",
			},
			expected: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			req := graphql.NewRequest(`
		mutation auth($data : SignInUserInput ){
			auth{
    			signInUser(data:$data){
      			access_token
    			}
  			}
		}
			`)
			var respData interface{}
			req.Var("data", testCase.user)
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

func TestRegistration(t *testing.T) {
	client, _ := graphqlTest.GetClient()

	testTable := []struct {
		name     string
		user     *api.SignUpUserInput
		expected bool
	}{
		{
			name: "Done",
			user: &api.SignUpUserInput{
				Email:    "serg22",
				Password: "Sergey2222",
				Name:     "Sergey Kalinin",
			},
			expected: true,
		},
		{
			name: "Error",
			user: &api.SignUpUserInput{
				Email:    "sergeykalinin@gmail",
				Password: "Ser",
				Name:     "Sergey Kalinin",
			},
			expected: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			req := graphql.NewRequest(`
mutation reg($data : SignUpUserInput ){
  auth{
    signUpUser(data:$data){
      access_token
    }
  }
	}
			`)
			var respData interface{}
			req.Var("data", testCase.user)
			err := client.Run(ctx, req, &respData)
			g_rpc.Log(respData)
			log.Error(err)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
}
