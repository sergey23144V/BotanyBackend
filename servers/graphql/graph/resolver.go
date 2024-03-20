package graph

import (
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db             *gorm.DB
	AuthServerImpl *auth.AuthServerImpl
	service        *service.Service
}

func NewResolver(Db *gorm.DB, AuthServerImpl *auth.AuthServerImpl, Service *service.Service) *Resolver {
	return &Resolver{
		Db:             Db,
		AuthServerImpl: AuthServerImpl,
		service:        Service,
	}
}

func ToIdRequest(id string) *api.IdRequest {
	return &api.IdRequest{Id: &resource.Identifier{ResourceId: id}}
}
