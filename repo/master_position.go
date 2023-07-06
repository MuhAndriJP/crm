package repo

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/crm/domain"
)

type MasterPositionRepo interface {
	Create(context.Context, domain.MasterPosition) error
	Find(context.Context) ([]domain.MasterPosition, error)
	FindById(context.Context, int64) (domain.MasterPosition, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, domain.MasterPosition) error
}

type masterPositionRepo struct {
}

func NewMasterPositionRepo() MasterPositionRepo {
	return &masterPositionRepo{}
}

func (mp *masterPositionRepo) Create(ctx context.Context, masterPosition domain.MasterPosition) (err error) {
	err = domain.NewMasterPosition().Create(ctx, masterPosition)
	return
}

func (mp *masterPositionRepo) Find(ctx context.Context) (masterPosition []domain.MasterPosition, err error) {
	masterPosition, err = domain.NewMasterPosition().Find(ctx)
	return
}

func (mp *masterPositionRepo) FindById(ctx context.Context, id int64) (masterPosition domain.MasterPosition, err error) {
	masterPosition, err = domain.NewMasterPosition().FindById(ctx, id)
	if masterPosition.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func (mp *masterPositionRepo) Delete(ctx context.Context, id int64) (err error) {
	err = domain.NewMasterPosition().Delete(ctx, id)
	return
}

func (mp *masterPositionRepo) Update(ctx context.Context, id int64, req domain.MasterPosition) (err error) {
	err = domain.NewMasterPosition().Update(ctx, id, req)
	return
}
