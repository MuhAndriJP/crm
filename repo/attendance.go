package repo

import (
	"context"
	"errors"
	"time"

	"github.com/MuhAndriJP/crm/domain"
)

type AttendanceRepo interface {
	Create(context.Context, domain.Attendance) error
	Find(context.Context) ([]domain.Attendance, error)
	FindById(context.Context, int64) (domain.Attendance, error)
	Delete(context.Context, int64) error
	Update(context.Context, int64, domain.Attendance) error
}

type attendanceRepo struct {
}

func NewAttendanceRepo() AttendanceRepo {
	return &attendanceRepo{}
}

func (a *attendanceRepo) Create(ctx context.Context, attendance domain.Attendance) (err error) {
	getAttendanceByEmployeeID, err := domain.NewAttendance().FindByEmployeeId(ctx, attendance.EmployeeId)
	if err != nil {
		return
	}

	now := time.Now()
	if getAttendanceByEmployeeID.AbsentIn.IsZero() {
		attendance.AbsentIn = now
		err = domain.NewAttendance().Create(ctx, attendance)
		return
	} else if !getAttendanceByEmployeeID.AbsentIn.IsZero() && getAttendanceByEmployeeID.AbsentOut != nil {
		attendance.AbsentIn = now
		err = domain.NewAttendance().Create(ctx, attendance)
		return
	}

	attendance.AbsentOut = &now
	err = domain.NewAttendance().Update(ctx, getAttendanceByEmployeeID.Id, attendance)
	return
}

func (a *attendanceRepo) Find(ctx context.Context) (attendance []domain.Attendance, err error) {
	attendance, err = domain.NewAttendance().Find(ctx)
	return
}

func (a *attendanceRepo) FindById(ctx context.Context, id int64) (attendance domain.Attendance, err error) {
	attendance, err = domain.NewAttendance().FindById(ctx, id)
	if attendance.Id == 0 {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func (a *attendanceRepo) Delete(ctx context.Context, id int64) (err error) {
	err = domain.NewAttendance().Delete(ctx, id)
	return
}

func (a *attendanceRepo) Update(ctx context.Context, id int64, req domain.Attendance) (err error) {
	err = domain.NewAttendance().Update(ctx, id, req)
	return
}
