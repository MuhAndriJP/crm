package master_location

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type createMasterLocation struct{}

func NewCreateMasterLocation() *createMasterLocation {
	return &createMasterLocation{}
}

func (md *createMasterLocation) Handle(ctx *fiber.Ctx) error {
	req := domain.MasterLocation{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	log.Println("Req", req)
	err = repo.NewMasterLocationRepo().Create(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
