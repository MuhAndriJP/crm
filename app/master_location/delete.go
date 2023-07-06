package master_location

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type deleteMasterLocation struct{}

func NewDeleteMasterLocation() *deleteMasterLocation {
	return &deleteMasterLocation{}
}

func (md *deleteMasterLocation) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("Req", id)
	err := repo.NewMasterLocationRepo().Delete(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
