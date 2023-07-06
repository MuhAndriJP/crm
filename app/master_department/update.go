package master_department

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type updateMasterDepartment struct{}

func NewUpdateMasterDepartment() *updateMasterDepartment {
	return &updateMasterDepartment{}
}
func (md *updateMasterDepartment) Handle(ctx *fiber.Ctx) error {
	id := helper.StrToInt64(ctx.Params("id"))

	req := domain.MasterDepartment{}
	errParser := ctx.BodyParser(&req)
	if errParser != nil {
		return errParser
	}

	log.Println("Req", id, req)
	err := repo.NewMasterDepartmentRepo().Update(ctx.Context(), id, req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
