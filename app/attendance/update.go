package attendance

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type updateAttendance struct{}

func NewUpdateAttendance() *updateAttendance {
	return &updateAttendance{}
}
func (md *updateAttendance) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	req := domain.Attendance{}
	errParser := ctx.BodyParser(&req)
	if errParser != nil {
		return errParser
	}

	log.Println("Req", id, req)
	err := repo.NewAttendanceRepo().Update(ctx.Context(), id, req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
