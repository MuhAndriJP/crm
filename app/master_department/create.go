package master_department

import (
	"log"

	"github.com/MuhAndriJP/crm/domain"
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type createMasterDepartment struct{}

func NewCreateMasterDepartment() *createMasterDepartment {
	return &createMasterDepartment{}
}

func (md *createMasterDepartment) Handle(ctx *fiber.Ctx) error {
	req := domain.MasterDepartment{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	log.Println("Req", req)
	err = repo.NewMasterDepartmentRepo().Create(ctx.Context(), req)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})
}
