package attendance

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type deleteAttendance struct{}

func NewDeleteAttendance() *deleteAttendance {
	return &deleteAttendance{}
}

func (md *deleteAttendance) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))
	employeeId := helper.ParseToken(ctx)
	if id != employeeId {
		return ctx.JSON(map[string]interface{}{
			"error": "Unathorized",
		})
	}

	log.Println("Req", id)
	err := repo.NewAttendanceRepo().Delete(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
