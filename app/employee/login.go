package employee

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type loginEmployee struct{}

func NewLoginEmployee() *loginEmployee {
	return &loginEmployee{}
}

func (le *loginEmployee) Handle(ctx *fiber.Ctx) error {
	req := domain.Employee{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	log.Println("Req", req)
	token, err := repo.NewEmployeeRepo().Login(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}
