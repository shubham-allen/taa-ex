package data

import (
	"taa-ex/internal/biz"
	"taa-ex/internal/data/entity"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type taaEngineRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewTaaEngineRepository(data *Data, logger log.Logger) biz.TaaEngineRepo {
	return &taaEngineRepository{data.db, log.NewHelper(logger)}
}

func (u *taaEngineRepository) Create(ctx context.Context, entity *entity.TaaEngineEntity) (err error) {
	// TODO : this needs to be filled with the right logic
	return nil
}
