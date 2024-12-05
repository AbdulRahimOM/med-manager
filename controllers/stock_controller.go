package controllers

import (
	respcode "med-manager/domain/respcodes"
	"med-manager/domain/response"
	models "med-manager/models"
	"med-manager/utils/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type StockController struct {
	DB *gorm.DB
}

func NewStockController(db *gorm.DB) *StockController {
	return &StockController{DB: db}
}

func (c *StockController) AddToStock(ctx *fiber.Ctx) error {
	stockUpdationReq := new(models.StockUpdateRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockUpdationReq); !ok {
		return errResponse
	}

	if err := stockUpdationReq.AddToStock(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, nil)
}

func (c *StockController) GetAllStockAdditions(ctx *fiber.Ctx) error {
	page, err := ctx.ParamsInt("page")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "page", err)
	}

	limit, err := ctx.ParamsInt("limit")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "limit", err)
	}

	stockAdditions, err := models.GetAllStockUpdations(c.DB, true, page, limit)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockAdditions)
}

func (c *StockController) GetStockUpdation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	stockAddition, err := models.GetStockUpdationByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockAddition)
}

func (c *StockController) UpdateStockUpdation(ctx *fiber.Ctx) error {
	stockUpdations := new(models.UpdateStockUpdateRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockUpdations); !ok {
		return errResponse
	}

	if err := models.UpdateParticularsInAnStockUpdation(c.DB, stockUpdations.ID, stockUpdations.StockChanges); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *StockController) DeleteStockUpdation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	if err := models.DeleteStockUpdation(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *StockController) GetStockAdditionsByMedicineID(ctx *fiber.Ctx) error {
	medicineID, err := ctx.ParamsInt("medicine_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "medicine_id", err)
	}
	stockAdditions, err := models.GetStockUpdationParticularsByMedicineID(c.DB, medicineID)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockAdditions)
}

func (c *StockController) DeductFromStock(ctx *fiber.Ctx) error {
	stockDeductions := new(models.StockUpdateRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockDeductions); !ok {
		return errResponse
	}

	if err := stockDeductions.DeductFromStock(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, nil)
}

func (c *StockController) GetAllStockDeductions(ctx *fiber.Ctx) error {
	page, err := ctx.ParamsInt("page")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "page", err)
	}

	limit, err := ctx.ParamsInt("limit")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "limit", err)
	}

	stockDeductions, err := models.GetAllStockUpdations(c.DB, false, page, limit)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockDeductions)
}

func (c *StockController) GetStockDeductionsByMedicineID(ctx *fiber.Ctx) error {
	medicineID, err := ctx.ParamsInt("medicine_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "medicine_id", err)
	}
	stockDeductions, err := models.GetStockUpdationParticularsByMedicineID(c.DB, medicineID)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockDeductions)
}
