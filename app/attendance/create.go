package attendance

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type Attendance struct{}

func NewCreateAttendance() *Attendance {
	return &Attendance{}
}

func (md *Attendance) Handle(ctx *fiber.Ctx) error {
	req := domain.Attendance{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	employeeId := helper.ParseToken(ctx)
	req.EmployeeId = employeeId

	log.Println("Req", req)
	err = repo.NewAttendanceRepo().Create(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
