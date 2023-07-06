package master_position

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type createMasterPosition struct{}

func NewCreateMasterPosition() *createMasterPosition {
	return &createMasterPosition{}
}

func (md *createMasterPosition) Handle(ctx *fiber.Ctx) error {
	req := domain.MasterPosition{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	log.Println("Req", req)
	err = repo.NewMasterPositionRepo().Create(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
