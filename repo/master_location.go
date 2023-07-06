package repo

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/crm/domain"
)

type MasterLocationRepo interface {
	Create(context.Context, domain.MasterLocation) error
	Find(context.Context) ([]domain.MasterLocation, error)
	FindById(context.Context, int64) (domain.MasterLocation, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, domain.MasterLocation) error
}

type masterLocationRepo struct {
}

func NewMasterLocationRepo() MasterLocationRepo {
	return &masterLocationRepo{}
}

func (ml *masterLocationRepo) Create(ctx context.Context, masterLocation domain.MasterLocation) (err error) {
	err = domain.NewMasterLocation().Create(ctx, masterLocation)
	return
}

func (ml *masterLocationRepo) Find(ctx context.Context) (masterLocation []domain.MasterLocation, err error) {
	masterLocation, err = domain.NewMasterLocation().Find(ctx)
	return
}

func (ml *masterLocationRepo) FindById(ctx context.Context, id int64) (masterLocation domain.MasterLocation, err error) {
	masterLocation, err = domain.NewMasterLocation().FindById(ctx, id)
	if masterLocation.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func (ml *masterLocationRepo) Delete(ctx context.Context, id int64) (err error) {
	err = domain.NewMasterLocation().Delete(ctx, id)
	return
}

func (ml *masterLocationRepo) Update(ctx context.Context, id int64, req domain.MasterLocation) (err error) {
	err = domain.NewMasterLocation().Update(ctx, id, req)
	return
}
