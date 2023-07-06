package master_location

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type getMasterLocation struct {
}

func NewGetMasterLocation() *getMasterLocation {
	return &getMasterLocation{}
}

func (md *getMasterLocation) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("req", id)
	res, err := repo.NewMasterLocationRepo().FindById(ctx.Context(), id)
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
