package repo

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/helper"
)

type EmployeeRepo interface {
	Create(context.Context, domain.Employee) error
	Find(context.Context) ([]domain.Employee, error)
	FindById(context.Context, int64) (domain.Employee, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, domain.Employee) error
	Login(context.Context, domain.Employee) (string, error)
}

type employeeRepo struct {
}

func NewEmployeeRepo() EmployeeRepo {
	return &employeeRepo{}
}

func (e *employeeRepo) Create(ctx context.Context, employee domain.Employee) (err error) {
	employee.Password = helper.CryptPassword(employee.Password)
	res, err := domain.NewEmployee().Create(ctx, employee)
	if err != nil {
		return
	}

	getEmployee, err := domain.NewEmployee().FindById(ctx, res.Id)
	if err != nil {
		return
	}
	if getEmployee.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	employee.Code = helper.GenerateCode(getEmployee.Id)

	err = domain.NewEmployee().Update(ctx, getEmployee.Id, employee)
	return
}

func (e *employeeRepo) Find(ctx context.Context) (employee []domain.Employee, err error) {
	employee, err = domain.NewEmployee().Find(ctx)
	return
}

func (e *employeeRepo) FindById(ctx context.Context, id int64) (employee domain.Employee, err error) {
	employee, err = domain.NewEmployee().FindById(ctx, id)
	if employee.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func (e *employeeRepo) Delete(ctx context.Context, id int64) (err error) {
	err = domain.NewEmployee().Delete(ctx, id)
	return
}

func (e *employeeRepo) Update(ctx context.Context, id int64, req domain.Employee) (err error) {
	err = domain.NewEmployee().Update(ctx, id, req)
	return
}

func (e *employeeRepo) Login(ctx context.Context, req domain.Employee) (token string, err error) {
	getEmployee, err := domain.NewEmployee().FindByCode(ctx, req.Code)
	if err != nil {
		return
	}
	if getEmployee.Id == 0 {
		err = errors.New("Employee belum terdaftar")
		return
	}

	if err = helper.CheckPassword(getEmployee.Password, req.Password); err != nil {
		err = errors.New("Password kamu salah")
		return
	}

	token = helper.GenerateToken(getEmployee.Id)
	err = domain.NewEmployee().Update(ctx, getEmployee.Id, getEmployee)
	return
}
