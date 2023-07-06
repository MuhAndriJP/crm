package attendance

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type getAttendance struct {
}

func NewGetAttendance() *getAttendance {
	return &getAttendance{}
}

func (md *getAttendance) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("req", id)
	res, err := repo.NewAttendanceRepo().FindById(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
		"data":    res,
	})
}
