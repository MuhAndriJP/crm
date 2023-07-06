package attendance

import (
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type listAttendance struct{}

func NewListAttendance() *listAttendance {
	return &listAttendance{}
}

func (md *listAttendance) Handle(ctx *fiber.Ctx) error {
	res, err := repo.NewAttendanceRepo().Find(ctx.Context())
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
		"data":    res,
	})
}
