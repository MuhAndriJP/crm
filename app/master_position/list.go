package master_position

import (
	"github.com/MuhAndriJP/crm/repo"
	"github.com/gofiber/fiber/v2"
)

type listMasterPosition struct{}

func NewListMasterPosition() *listMasterPosition {
	return &listMasterPosition{}
}

func (md *listMasterPosition) Handle(ctx *fiber.Ctx) error {
	res, err := repo.NewMasterPositionRepo().Find(ctx.Context())
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
