package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "taa-ex/api/taa_engine/v1"
	"taa-ex/internal/biz"
)

// TaaEngineService is a taaEngine service.
type TaaEngineService struct {
	v1.UnimplementedTaaEngineServer
	handler *biz.TaaEngineHandler
	log     *log.Helper
}

// NewTaaEngineService new a taaEngine service.
func NewTaaEngineService(handler *biz.TaaEngineHandler, logger log.Logger) *TaaEngineService {
	return &TaaEngineService{
		UnimplementedTaaEngineServer: v1.UnimplementedTaaEngineServer{},
		handler:                 handler,
		log:                     log.NewHelper(logger),
	}
}

// Test Create function to use request and response
func (s *TaaEngineService) CreateTaaEngine(ctx context.Context, req *v1.CreateTaaEngineRequest) (*v1.CreateTaaEngineResponse, error) {
	info := &v1.TaaEngineInfo{
		Name: req.Name,
	}

	result, err := s.handler.CreateTaaEngine(ctx, info)
	if err != nil {
		return nil, err
	}

	return &v1.CreateTaaEngineResponse{Info: result}, nil
}
