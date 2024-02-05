package service

//type EcomorphsEntityService interface {
//	Insert(input g_rpc.CreateEcomorphsEntityRequest) (*g_rpc.EcomorphsEntity, error)
//	GetById(inputId string) (*g_rpc.EcomorphsEntity, error)
//	GetList() ([]*g_rpc.EcomorphsEntity, error)
//}
//
//type EcomorphService interface {
//	Insert(input g_rpc.CreateEcomorphsRequest) (*g_rpc.Ecomorph, error)
//	GetById(inputId string) (*g_rpc.Ecomorph, error)
//	GetList() ([]*g_rpc.Ecomorph, error)
//}
//
//type Service struct {
//	EcomorphsEntityService
//	EcomorphService
//}
//
//func NewService(repos *repository.Repository) *Service {
//	return &Service{
//		EcomorphsEntityService: NewEcomorphsEntityServiceImpl(repos),
//		EcomorphService:        NewEcomorphServiceImpl(repos),
//	}
//}
