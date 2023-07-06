package domain

import (
	"context"

	"github.com/MuhAndriJP/crm/config"
	"github.com/jinzhu/gorm"
)

type MasterLocationInterface interface {
	Find(context.Context) ([]MasterLocation, error)
	FindById(context.Context, int64) (MasterLocation, error)
	Create(context.Context, MasterLocation) error
	Delete(context.Context, int64) error
	Update(context.Context, int64, MasterLocation) error
}

type masterLocationRepo struct {
	db *gorm.DB
}

func NewMasterLocation() MasterLocationInterface {
	return &masterLocationRepo{
		db: config.DB,
	}
}

func (ml *masterLocationRepo) Find(ctx context.Context) (res []MasterLocation, err error) {
	result := ml.db.Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (ml *masterLocationRepo) FindById(ctx context.Context, id int64) (res MasterLocation, err error) {
	result := ml.db.Find(&res, id)
	if result.Error != nil {
		return
	}

	return
}

func (ml *masterLocationRepo) Create(ctx context.Context, masterLocation MasterLocation) (err error) {
	result := ml.db.Save(&masterLocation)
	if result.Error != nil {
		return
	}
	return
}

func (ml *masterLocationRepo) Delete(ctx context.Context, id int64) (err error) {
	data := MasterLocation{}
	result := ml.db.Delete(&data, id)
	if result.Error != nil {
		return
	}

	return
}

func (ml *masterLocationRepo) Update(ctx context.Context, id int64, req MasterLocation) (err error) {
	data := MasterLocation{}
	result := ml.db.First(&data, id)
	if result.Error != nil {
		return result.Error
	}

	result = ml.db.Model(&data).Updates(req)
	if result.Error != nil {
		return result.Error
	}

	return
}
