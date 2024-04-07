package graph

import (
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"gorm.io/gorm"

	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db             *gorm.DB
	AuthServerImpl *implementation.AuthServerImpl
	service        *service.Service
}

func NewResolver(Db *gorm.DB, AuthServerImpl *implementation.AuthServerImpl, Service *service.Service) *Resolver {
	return &Resolver{
		Db:             Db,
		AuthServerImpl: AuthServerImpl,
		service:        Service,
	}
}

func ToIdRequest(id string) *api.IdRequest {
	return &api.IdRequest{Id: &resource.Identifier{ResourceId: id}}
}
