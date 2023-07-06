package domain

import (
	"context"

	"github.com/MuhAndriJP/crm/config"
	"github.com/jinzhu/gorm"
)

type MasterDepartmentInterface interface {
	Find(context.Context) ([]MasterDepartment, error)
	FindById(context.Context, int64) (MasterDepartment, error)
	Create(context.Context, MasterDepartment) error
	Delete(context.Context, int64) error
	Update(context.Context, int64, MasterDepartment) error
}

type masterDepartmentRepo struct {
	db *gorm.DB
}

func NewMasterDepartment() MasterDepartmentInterface {
	return &masterDepartmentRepo{
		db: config.DB,
	}
}

func (md *masterDepartmentRepo) Find(ctx context.Context) (res []MasterDepartment, err error) {
	result := md.db.Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (md *masterDepartmentRepo) FindById(ctx context.Context, id int64) (res MasterDepartment, err error) {
	result := md.db.Find(&res, id)
	if result.Error != nil {
		return
	}

	return
}

func (md *masterDepartmentRepo) Create(ctx context.Context, masterDepartment MasterDepartment) (err error) {
	result := md.db.Save(&masterDepartment)
	if result.Error != nil {
		return
	}
	return
}

func (md *masterDepartmentRepo) Delete(ctx context.Context, id int64) (err error) {
	data := MasterDepartment{}
	result := md.db.Delete(&data, id)
	if result.Error != nil {
		return
	}

	return
}

func (md *masterDepartmentRepo) Update(ctx context.Context, id int64, req MasterDepartment) (err error) {
	data := MasterDepartment{}
	result := md.db.First(&data, id)
	if result.Error != nil {
		return result.Error
	}

	result = md.db.Model(&data).Updates(req)
	if result.Error != nil {
		return result.Error
	}

	return
}
