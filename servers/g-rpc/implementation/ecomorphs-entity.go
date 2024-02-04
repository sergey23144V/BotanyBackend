package implementation

import (
	"github.com/sergey23144V/BotanyBackend/pkg/service"
)

type EcomorphsEntityServetImpl struct {
	Sevice service.Service
}

func NewEcomorphsEntityServetImpl(sevice service.Service) EcomorphsEntityServetImpl {
	return EcomorphsEntityServetImpl{sevice}
}
