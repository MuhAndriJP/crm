package domain

import (
	"context"

	"github.com/MuhAndriJP/crm/config"
	"github.com/jinzhu/gorm"
)

type EmployeeInterface interface {
	Find(context.Context) ([]Employee, error)
	FindById(context.Context, int64) (Employee, error)
	FindByCode(context.Context, string) (Employee, error)
	Create(context.Context, Employee) (Employee, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, Employee) error
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployee() EmployeeInterface {
	return &employeeRepo{
		db: config.DB,
	}
}

func (e *employeeRepo) Find(ctx context.Context) (res []Employee, err error) {
	result := e.db.Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (e *employeeRepo) FindById(ctx context.Context, id int64) (res Employee, err error) {
	result := e.db.Find(&res, id)
	if result.Error != nil {
		return
	}

	return
}

func (e *employeeRepo) FindByCode(ctx context.Context, code string) (res Employee, err error) {
	result := e.db.Where("code = ?", code).First(&res)
	if result.Error != nil {
		return
	}

	return
}

func (e *employeeRepo) Create(ctx context.Context, Employee Employee) (res Employee, err error) {
	result := e.db.Save(&Employee)
	if result.Error != nil {
		return
	}
	res = Employee
	return
}

func (e *employeeRepo) Delete(ctx context.Context, id int64) (err error) {
	data := Employee{}
	result := e.db.Delete(&data, id)
	if result.Error != nil {
		return
	}

	return
}

func (e *employeeRepo) Update(ctx context.Context, id int64, req Employee) (err error) {
	data := Employee{}
	result := e.db.First(&data, id)
	if result.Error != nil {
		return result.Error
	}

	result = e.db.Model(&data).Updates(req)
	if result.Error != nil {
		return result.Error
	}

	return
}
