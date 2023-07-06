package domain

import (
	"context"

	"github.com/MuhAndriJP/crm/config"
	"github.com/jinzhu/gorm"
)

type MasterPositionInterface interface {
	Find(context.Context) ([]MasterPosition, error)
	FindById(context.Context, int64) (MasterPosition, error)
	Create(context.Context, MasterPosition) error
	Delete(context.Context, int64) error
	Update(context.Context, int64, MasterPosition) error
}

type masterPositionRepo struct {
	db *gorm.DB
}

func NewMasterPosition() MasterPositionInterface {
	return &masterPositionRepo{
		db: config.DB,
	}
}

func (mp *masterPositionRepo) Find(ctx context.Context) (res []MasterPosition, err error) {
	result := mp.db.Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (mp *masterPositionRepo) FindById(ctx context.Context, id int64) (res MasterPosition, err error) {
	result := mp.db.Find(&res, id)
	if result.Error != nil {
		return
	}

	return
}

func (mp *masterPositionRepo) Create(ctx context.Context, masterPosition MasterPosition) (err error) {
	result := mp.db.Save(&masterPosition)
	if result.Error != nil {
		return
	}
	return
}

func (mp *masterPositionRepo) Delete(ctx context.Context, id int64) (err error) {
	data := MasterPosition{}
	result := mp.db.Delete(&data, id)
	if result.Error != nil {
		return
	}

	return
}

func (mp *masterPositionRepo) Update(ctx context.Context, id int64, req MasterPosition) (err error) {
	data := MasterPosition{}
	result := mp.db.First(&data, id)
	if result.Error != nil {
		return result.Error
	}

	result = mp.db.Model(&data).Updates(req)
	if result.Error != nil {
		return result.Error
	}

	return
}
