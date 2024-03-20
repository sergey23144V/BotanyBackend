package auth

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
	"github.com/sergey23144V/BotanyBackend/test/g-rpc"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

const titleToken = "Bearer "

func GetAuthClient() (auth.AuthServiceClient, *grpc.ClientConn) {
	g_rpc.StartServerGRPC()

	time.Sleep(2 * time.Second)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error("Could not connect: ", err)
	}

	client := auth.NewAuthServiceClient(conn)

	return client, conn

}

func Authorisation(ctx context.Context, authClient auth.AuthServiceClient, input *auth.SignInUserInput) (*auth.SignInUserResponse, context.Context, error) {
	user, err := authClient.SignInUser(ctx, input)
	if err != nil {
		log.Error("Error")
		return nil, nil, err
	}

	md := metadata.Pairs("authorization", titleToken+user.AccessToken)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return user, ctx, err
}

func Registration(ctx context.Context, authClient auth.AuthServiceClient, input *auth.SignUpUserInput) (*auth.SignInUserResponse, context.Context, error) {
	user, err := authClient.SignUpUser(ctx, input)
	if err != nil {
		log.Error("Registration ERROR", err)
		return nil, ctx, err
	}

	md := metadata.Pairs("authorization", titleToken+user.AccessToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return user, ctx, err
}

func TestAuthorisation(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	authClient, conn := GetAuthClient()

	testTable := []struct {
		name     string
		user     *auth.SignInUserInput
		expected bool
	}{
		{
			name: "Done",
			user: &auth.SignInUserInput{
				Email:    "serg2",
				Password: "Sergey2222",
			},
			expected: true,
		},
		{
			name: "Error",
			user: &auth.SignInUserInput{
				Email:    "sergeykalinin@gmail.con",
				Password: "Serg",
			},
			expected: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			_, ctx, err := Authorisation(ctx, authClient, testCase.user)
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
	authClient, conn := GetAuthClient()

	testTable := []struct {
		name     string
		user     *auth.SignUpUserInput
		expected bool
	}{
		{
			name: "Done",
			user: &auth.SignUpUserInput{
				Email:    "sergeyklinin@gmail.com",
				Password: "Sergey24222",
				Name:     "Sergey Kalinin",
			},
			expected: true,
		},
		{
			name: "Error",
			user: &auth.SignUpUserInput{
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
			_, ctx, err := Registration(ctx, authClient, testCase.user)
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
