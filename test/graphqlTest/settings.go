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
	"reflect"
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

	StartServerGraphQl()

	time.Sleep(2 * time.Second)

	user := &api.SignInUserInput{
		Email:    "serg2",
		Password: "Sergey2222",
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

func toLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]+32) + s[1:]
}

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	if objType.Kind() != reflect.Struct {
		if objType.Kind() == reflect.Ptr {
			if objValue.IsNil() {
				return nil
			}
			objValue = objValue.Elem()
			objType = objValue.Type()
		}
	}

	data := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)

		fieldValue := objValue.Field(i)
		if field.Name == "state" || field.Name == "sizeCache" || field.Name == "unknownFields" {
			continue
		}

		if field.Tag.Get("json") == "-" {
			continue
		}

		// Если тип поля является указателем, следует разыменовать его
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue
			}
			fieldValue = fieldValue.Elem()
		}

		fieldName := toLowerFirst(field.Name)
		// Если тип поля является структурой, рекурсивно вызываем функцию structToMap
		if fieldValue.Kind() == reflect.Struct {
			nestedData := StructToMap(fieldValue.Interface())
			data[fieldName] = nestedData
		} else if fieldValue.Kind() == reflect.Slice {
			var arrData []interface{}
			for j := 0; j < fieldValue.Len(); j++ {
				d := fieldValue.Index(j).Interface()
				nestedData := StructToMap(d)
				arrData = append(arrData, nestedData)
			}
			data[fieldName] = arrData
		} else {
			if fieldValue.Kind() == reflect.String && fieldValue.Interface().(string) == "" {
				continue
			}
			data[fieldName] = fieldValue.Interface()
		}
	}

	return data
}
