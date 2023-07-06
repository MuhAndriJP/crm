package master_location

import (
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type listMasterLocation struct{}

func NewListMasterLocation() *listMasterLocation {
	return &listMasterLocation{}
}

func (md *listMasterLocation) Handle(ctx *fiber.Ctx) error {
	res, err := repo.NewMasterLocationRepo().Find(ctx.Context())
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
