package repo

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/crm/domain"
)

type MasterDepartmentRepo interface {
	Create(context.Context, domain.MasterDepartment) error
	Find(context.Context) ([]domain.MasterDepartment, error)
	FindById(context.Context, int64) (domain.MasterDepartment, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, domain.MasterDepartment) error
}

type masterDepartmentRepo struct {
}

func NewMasterDepartmentRepo() MasterDepartmentRepo {
	return &masterDepartmentRepo{}
}

func (md *masterDepartmentRepo) Create(ctx context.Context, masterDepartment domain.MasterDepartment) (err error) {
	err = domain.NewMasterDepartment().Create(ctx, masterDepartment)
	return
}

func (md *masterDepartmentRepo) Find(ctx context.Context) (masterDepartment []domain.MasterDepartment, err error) {
	masterDepartment, err = domain.NewMasterDepartment().Find(ctx)
	return
}

func (md *masterDepartmentRepo) FindById(ctx context.Context, id int64) (masterDepartment domain.MasterDepartment, err error) {
	masterDepartment, err = domain.NewMasterDepartment().FindById(ctx, id)
	if masterDepartment.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func (md *masterDepartmentRepo) Delete(ctx context.Context, id int64) (err error) {
	err = domain.NewMasterDepartment().Delete(ctx, id)
	return
}

func (md *masterDepartmentRepo) Update(ctx context.Context, id int64, req domain.MasterDepartment) (err error) {
	err = domain.NewMasterDepartment().Update(ctx, id, req)
	return
}
