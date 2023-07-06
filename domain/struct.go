package domain

import "time"

type MasterDepartment struct {
	Id        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type MasterPosition struct {
	Id           int64      `json:"id" db:"id"`
	DepartmentId int64      `json:"department_id" db:"department_id"`
	Name         string     `json:"name" db:"name"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type MasterLocation struct {
	Id        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Employee struct {
	Id           int64      `json:"id" db:"id"`
	Code         string     `json:"code" db:"code"`
	Name         string     `json:"name" db:"name"`
	Password     string     `json:"password" db:"password"`
	DepartmentId int64      `json:"department_id" db:"department_id"`
	PositionId   int64      `json:"position_id" db:"position_id"`
	Superior     int64      `json:"superior" db:"superior"`
	Created_at   time.Time  `json:"created_at" db:"created_at"`
	Created_by   string     `json:"created_by" db:"created_by"`
	Updated_at   time.Time  `json:"updated_at" db:"updated_at"`
	Updated_by   string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type Attendance struct {
	Id         int64      `json:"id" db:"id"`
	EmployeeId int64      `json:"employee_id" db:"employee_id"`
	LocationId int64      `json:"location_id" db:"location_id"`
	AbsentIn   time.Time  `json:"absent_in" db:"absent_in"`
	AbsentOut  *time.Time `json:"absent_out"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
	Created_by string     `json:"created_by" db:"created_by"`
	Updated_at time.Time  `json:"updated_at" db:"updated_at"`
	Updated_by string     `json:"updated_by" db:"updated_by"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
