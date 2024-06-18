package graphqlTest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	client "github.com/machinebox/graphql"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"github.com/sergey23144V/BotanyBackend/servers/graphql"
	"log"
	"time"
)

func StartServerGraphQl() {
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

		err = db.AutoMigrate(api.ImgORM{}, api.EcomorphORM{}, api.EcomorphsEntityORM{}, api.UserORM{}, api.TypePlantORM{}, api.TrialSiteORM{}, api.TransectORM{})
		log.Println("migrant")

		authServet := implementation.NewAuthServer(db)

		newRepository := repository.NewRepository(db)
		newService := service.NewService(newRepository)

		graphql.StartGraphQl(db, &authServet, newService)
	}()
}

func GetClient() (*client.Client, string) {

	//StartServerGraphQl()

	time.Sleep(2 * time.Second)

	user := &api.SignInUserInput{
		Email:    "Admin",
		Password: "Admin",
	}

	clientServer := client.NewClient("http://localhost:80/api")
	ctx := context.Background()
	req := client.NewRequest(`
		mutation auth($data : SignInUserInput ){
			auth{
    			signInUser(data:$data){
      			access_token
    			}
  			}
		}
			`)
	var respData interface{}
	req.Var("data", user)

	err := clientServer.Run(ctx, req, &respData)
	if err != nil {
		return nil, ""
	}
	data, err := json.Marshal(respData)
	if err != nil {
		return nil, ""
	}
	token, err := JsonPathValue(string(data), "$.auth.signInUser.access_token")
	if err != nil {
		return nil, ""
	}
	return clientServer, "Bearer " + token.(string)
}

func JsonPathValue(jsonData string, path string) (interface{}, error) {
	var jsonObj interface{}
	err := json.Unmarshal([]byte(jsonData), &jsonObj)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	res, err := jsonpath.Get(path, jsonObj)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON path: %v", err)
	}

	return res, nil
}
