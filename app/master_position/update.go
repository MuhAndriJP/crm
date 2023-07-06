package master_position

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type updateMasterPosition struct{}

func NewUpdateMasterPosition() *updateMasterPosition {
	return &updateMasterPosition{}
}
func (md *updateMasterPosition) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	req := domain.MasterPosition{}
	errParser := ctx.BodyParser(&req)
	if errParser != nil {
		return errParser
	}

	log.Println("Req", id, req)
	err := repo.NewMasterPositionRepo().Update(ctx.Context(), id, req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
