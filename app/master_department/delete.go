package master_department

import (
	"log"

	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type deleteMasterDepartment struct{}

func NewDeleteMasterDepartment() *deleteMasterDepartment {
	return &deleteMasterDepartment{}
}

func (md *deleteMasterDepartment) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	log.Println("Req", id)
	err := repo.NewMasterDepartmentRepo().Delete(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
