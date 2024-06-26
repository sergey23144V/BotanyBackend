package g_rpc

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func StartServerGRPC() {
	go func() {
		cfgDB := repository.Config{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			DBName:   "botany",
			SSLMode:  "disable",
			Password: "123321",
		}

		db, err := repository.ConnectDB(cfgDB)
		if err != nil {
			log.Fatal(err)
		}

		err = db.AutoMigrate(api.ImgORM{}, api.EcomorphORM{}, api.EcomorphsEntityORM{}, api.UserORM{}, api.TypePlantORM{}, api.TransectORM{}, api.TrialSiteORM{}, api.PlantORM{}, api.AnalysisORM{})
		if err != nil {
			return
		}
		log.Println("migrant")

		authServet := implementation.NewAuthServer(db)

		newRepository := repository.NewRepository(db)
		newService := service.NewService(newRepository)

		g_rpc.StartGrpc(&authServet, newService)
	}()

}

const titleToken = "Bearer "

func Authorisation(ctx context.Context, authClient api.AuthServiceClient, input *api.SignInUserInput) (*api.SignInUserResponse, context.Context, error) {
	user, err := authClient.SignInUser(ctx, input)
	if err != nil {
		log.Error("Error")
		return nil, nil, err
	}

	md := metadata.Pairs("authorization", titleToken+user.AccessToken)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return user, ctx, err
}

func GetToken(ctx context.Context, authClient api.AuthServiceClient) (context.Context, error) {

	input := &api.SignInUserInput{
		Email:    "Admin",
		Password: "Admin",
	}

	user, err := authClient.SignInUser(ctx, input)
	if err != nil {
		log.Error("Error")
		return nil, err
	}

	md := metadata.Pairs("authorization", titleToken+user.AccessToken)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, err
}

func Registration(ctx context.Context, authClient api.AuthServiceClient, input *api.SignUpUserInput) (*api.SignInUserResponse, context.Context, error) {
	user, err := authClient.SignUpUser(ctx, input)
	if err != nil {
		log.Error("Registration ERROR", err)
		return nil, ctx, err
	}

	md := metadata.Pairs("authorization", titleToken+user.AccessToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return user, ctx, err
}

func GetAuthClient() (api.AuthServiceClient, *grpc.ClientConn) {
	StartServerGRPC()

	time.Sleep(2 * time.Second)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error("Could not connect: ", err)
	}

	client := api.NewAuthServiceClient(conn)

	return client, conn

}

func GetClient() (*ClientService, context.Context) {
	//StartServerGRPC()

	time.Sleep(2 * time.Second)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error("Could not connect: ", err)
	}

	ecomorphClient := api.NewEcomorphServiceClient(conn)
	ecomorph_emtityClient := api.NewEcomorphEntityServiceClient(conn)
	typePlantClient := api.NewTypePlantServiceClient(conn)
	trialSiteClient := api.NewTrialSiteServiceClient(conn)
	transectClient := api.NewTransectServiceClient(conn)
	authClient := api.NewAuthServiceClient(conn)
	analysisClient := api.NewAnalysisServiceClient(conn)

	ctx := context.Background()
	ctxOut, err := GetToken(ctx, authClient)
	if err != nil {
		return nil, nil
	}
	return &ClientService{
		Ecomorph:        ecomorphClient,
		Ecomorph_Emtity: ecomorph_emtityClient,
		TypePlant:       typePlantClient,
		TrialSite:       trialSiteClient,
		Transect:        transectClient,
		Analysis:        analysisClient,
	}, ctxOut
}

type ClientService struct {
	Ecomorph        api.EcomorphServiceClient
	Ecomorph_Emtity api.EcomorphEntityServiceClient
	TypePlant       api.TypePlantServiceClient
	TrialSite       api.TrialSiteServiceClient
	Transect        api.TransectServiceClient
	Analysis        api.AnalysisServiceClient
}

func Log(args ...interface{}) {
	log.Println("___________________________________________________")
	log.Println(args...)
	log.Println("___________________________________________________")
}
