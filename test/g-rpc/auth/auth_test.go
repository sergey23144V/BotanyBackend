package auth

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/test/g-rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorisation(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	authClient, conn := g_rpc.GetAuthClient()

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
			_, ctx, err := g_rpc.Authorisation(ctx, authClient, testCase.user)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}

	err := conn.Close()
	if err != nil {
		return
	}

	return

	<-done
}

func TestRegistration(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	authClient, conn := g_rpc.GetAuthClient()

	testTable := []struct {
		name     string
		user     *api.SignUpUserInput
		expected bool
	}{
		{
			name: "Done",
			user: &api.SignUpUserInput{
				Email:    "serg2",
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
			_, ctx, err := g_rpc.Registration(ctx, authClient, testCase.user)
			if testCase.expected {
				assert.NoError(t, err, "Done")
			} else {
				assert.Error(t, err, "Error")
			}
		})
	}
	err := conn.Close()
	if err != nil {
		return
	}

	return

	<-done
}
