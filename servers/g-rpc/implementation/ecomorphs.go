package implementation

import (
	"github.com/sergey23144V/BotanyBackend/pkg/service"
)

// EcomorphsServetImpl
type EcomorphsServetImpl struct {
	Sevice service.Service
}

func NewEcomorphsServetImplImpl(sevice service.Service) EcomorphsServetImpl {
	return EcomorphsServetImpl{sevice}
}
