package employee

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type createEmployee struct{}

func NewCreateEmployee() *createEmployee {
	return &createEmployee{}
}

func (md *createEmployee) Handle(ctx *fiber.Ctx) error {
	req := domain.Employee{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	log.Println("Req", req)
	err = repo.NewEmployeeRepo().Create(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
