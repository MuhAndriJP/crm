package route

import (
	a "github.com/MuhAndriJP/crm/app/attendance"
	e "github.com/MuhAndriJP/crm/app/employee"
	md "github.com/MuhAndriJP/crm/app/master_department"
	ml "github.com/MuhAndriJP/crm/app/master_location"
	mp "github.com/MuhAndriJP/crm/app/master_position"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	department := app.Group("v1/department")
	department.Get("/", helper.JWTMiddleware(), md.NewListMasterDepartment().Handle)
	department.Get("/:id", helper.JWTMiddleware(), md.NewGetMasterDepartment().Handle)
	department.Delete("/:id", helper.JWTMiddleware(), md.NewDeleteMasterDepartment().Handle)
	department.Put("/:id", helper.JWTMiddleware(), md.NewUpdateMasterDepartment().Handle)
	department.Post("/", helper.JWTMiddleware(), md.NewCreateMasterDepartment().Handle)

	location := app.Group("v1/location")
	location.Get("/", helper.JWTMiddleware(), ml.NewListMasterLocation().Handle)
	location.Get("/:id", helper.JWTMiddleware(), ml.NewGetMasterLocation().Handle)
	location.Delete("/:id", helper.JWTMiddleware(), ml.NewDeleteMasterLocation().Handle)
	location.Put("/:id", helper.JWTMiddleware(), ml.NewUpdateMasterLocation().Handle)
	location.Post("/", helper.JWTMiddleware(), ml.NewCreateMasterLocation().Handle)

	position := app.Group("v1/position")
	position.Get("/", helper.JWTMiddleware(), mp.NewListMasterPosition().Handle)
	position.Get("/:id", helper.JWTMiddleware(), mp.NewGetMasterPosition().Handle)
	position.Delete("/:id", helper.JWTMiddleware(), mp.NewDeleteMasterPosition().Handle)
	position.Put("/:id", helper.JWTMiddleware(), mp.NewUpdateMasterPosition().Handle)
	position.Post("/", helper.JWTMiddleware(), mp.NewCreateMasterPosition().Handle)

	employee := app.Group("v1/employee")
	employee.Get("/", helper.JWTMiddleware(), e.NewGetEmployee().Handle)
	employee.Get("/:id", helper.JWTMiddleware(), e.NewListEmployee().Handle)
	employee.Delete("/:id", helper.JWTMiddleware(), e.NewDeleteEmployee().Handle)
	employee.Put("/:id", helper.JWTMiddleware(), e.NewUpdateEmployee().Handle)
	employee.Post("/", e.NewCreateEmployee().Handle)
	employee.Post("/login", e.NewLoginEmployee().Handle)

	attendance := app.Group("v1/attendance")
	attendance.Get("/", helper.JWTMiddleware(), a.NewListAttendance().Handle)
	attendance.Get("/:id", helper.JWTMiddleware(), a.NewGetAttendance().Handle)
	attendance.Delete("/:id", helper.JWTMiddleware(), a.NewDeleteAttendance().Handle)
	attendance.Put("/:id", helper.JWTMiddleware(), a.NewUpdateAttendance().Handle)
	attendance.Post("/", helper.JWTMiddleware(), a.NewCreateAttendance().Handle)
}
