package employee

import (
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type listEmployee struct{}

func NewListEmployee() *listEmployee {
	return &listEmployee{}
}

func (md *listEmployee) Handle(ctx *fiber.Ctx) error {
	res, err := repo.NewEmployeeRepo().Find(ctx.Context())
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
