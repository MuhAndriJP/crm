package employee

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type getEmployee struct {
}

func NewGetEmployee() *getEmployee {
	return &getEmployee{}
}

func (md *getEmployee) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("req", id)
	res, err := repo.NewEmployeeRepo().FindById(ctx.Context(), id)
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
