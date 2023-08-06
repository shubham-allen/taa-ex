package biz

import (
	"taa-ex/internal/data/entity"
	"context"
	v1 "taa-ex/api/taa_engine/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)


// TaaEngineRepo is a TaaEngine repo.
type TaaEngineRepo interface {
	BaseRepository[entity.TaaEngineEntity, string, uint]
}

// TaaEngineHandler is a TaaEngine usecase.
type TaaEngineHandler struct {
	repo TaaEngineRepo
	log  *log.Helper
}

// NewTaaEngineHandler is a new handler.
func NewTaaEngineHandler(repo TaaEngineRepo, logger log.Logger) *TaaEngineHandler {
	return &TaaEngineHandler{repo: repo, log: log.NewHelper(logger)}
}

// Sample function to generate to create taaEngine
func (handler *TaaEngineHandler) CreateTaaEngine(ctx context.Context, info *v1.TaaEngineInfo) (*v1.TaaEngineInfo, error) {
	handler.log.WithContext(ctx).Infof("CreateTaaEngine: %v", info)
	ent := &entity.TaaEngineEntity{}
	ent.FromPB(info)
	if err := handler.repo.Create(ctx, ent); err != nil {
		log.Errorf("entity not created: %v", info)
		return nil, err
	}
	ent.ToPB(info)
	return info, nil
}
