package domain

import (
	"context"

	"github.com/MuhAndriJP/crm/config"
	"github.com/jinzhu/gorm"
)

type AttendanceInterface interface {
	Find(context.Context) ([]Attendance, error)
	FindById(context.Context, int64) (Attendance, error)
	FindByEmployeeId(context.Context, int64) (Attendance, error)
	Create(context.Context, Attendance) error
	Delete(context.Context, int64) error
	Update(context.Context, int64, Attendance) error
}

type attendanceRepo struct {
	db *gorm.DB
}

func NewAttendance() AttendanceInterface {
	return &attendanceRepo{
		db: config.DB,
	}
}

func (a *attendanceRepo) Find(ctx context.Context) (res []Attendance, err error) {
	result := a.db.Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (a *attendanceRepo) FindById(ctx context.Context, id int64) (res Attendance, err error) {
	result := a.db.Find(&res, id)
	if result.Error != nil {
		return
	}

	return
}

func (a *attendanceRepo) FindByEmployeeId(ctx context.Context, id int64) (res Attendance, err error) {
	result := a.db.Where("employee_id = ?", id).Order("id desc").Limit(1).Find(&res)
	if result.Error != nil {
		return
	}

	return
}

func (a *attendanceRepo) Create(ctx context.Context, attendance Attendance) (err error) {
	result := a.db.Save(&attendance)
	if result.Error != nil {
		return
	}
	return
}

func (a *attendanceRepo) Delete(ctx context.Context, id int64) (err error) {
	data := Attendance{}
	result := a.db.Delete(&data, id)
	if result.Error != nil {
		return
	}

	return
}

func (a *attendanceRepo) Update(ctx context.Context, id int64, req Attendance) (err error) {
	data := Attendance{}
	result := a.db.First(&data, id)
	if result.Error != nil {
		return result.Error
	}

	result = a.db.Model(&data).Updates(req)
	if result.Error != nil {
		return result.Error
	}

	return
}
