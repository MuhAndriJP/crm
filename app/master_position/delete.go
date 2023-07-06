package master_position

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type deleteMasterPosition struct{}

func NewDeleteMasterPosition() *deleteMasterPosition {
	return &deleteMasterPosition{}
}

func (md *deleteMasterPosition) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("Req", id)
	err := repo.NewMasterPositionRepo().Delete(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
