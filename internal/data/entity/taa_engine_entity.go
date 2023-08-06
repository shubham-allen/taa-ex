package entity

import (
	v1 "taa-ex/api/taa_engine/v1"
)

type TaaEngineEntity struct {
	// please fill the appropriate fields
	Name string
}

func (entity *TaaEngineEntity) FromPB(info *v1.TaaEngineInfo) {
	entity.Name = info.Name
}

func (entity *TaaEngineEntity) ToPB(info *v1.TaaEngineInfo) {
	info.Name = entity.Name
}
