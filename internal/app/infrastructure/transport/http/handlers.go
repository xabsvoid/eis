package http

import "github.com/xabsvoid/eis/internal/app/domain/service"

type Handlers struct {
	service *service.Service
}

func NewHandlers(service *service.Service) Handlers {
	return Handlers{
		service: service,
	}
}
