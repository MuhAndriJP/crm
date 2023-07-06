package master_department

import (
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type listMasterDepartment struct{}

func NewListMasterDepartment() *listMasterDepartment {
	return &listMasterDepartment{}
}

func (md *listMasterDepartment) Handle(ctx *fiber.Ctx) error {
	res, err := repo.NewMasterDepartmentRepo().Find(ctx.Context())
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
