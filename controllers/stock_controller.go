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

	stockUpdationID, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	if err := models.UpdateParticularsInAnStockUpdation(c.DB, stockUpdationID, stockUpdations.StockChanges); err != nil {
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
	page := ctx.QueryInt("page", 1)
	limit, err := ctx.ParamsInt("limit", 10)

	stockAdditions, err := models.GetAllStockUpdations(c.DB, true, (page-1)*limit, limit)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockAdditions)
}

func (c *StockController) GetStockAdditionsByMedicineID(ctx *fiber.Ctx) error {
	medicineID, err := ctx.ParamsInt("medicine_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "medicine_id", err)
	}
	stockAdditions, err := models.GetStockUpdationParticularsByMedicineID(c.DB, medicineID, true)
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

	if err, insufficientMedID := stockDeductions.DeductFromStock(c.DB); err != nil {
		if err == models.ErrInsufficientStock {
			return response.Response{
				HttpStatusCode: 400,
				Status:         false,
				ResponseCode:   respcode.INSUFFICIENT_STOCK,
				Error:          err,
				Data: map[string]int{
					"medicine_id": insufficientMedID,
				},
			}.WriteToJSON(ctx)
		}
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, nil)
}

func (c *StockController) GetAllStockDeductions(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit, err := ctx.ParamsInt("limit", 10)

	stockDeductions, err := models.GetAllStockUpdations(c.DB, false, (page-1)*limit, limit)
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
	stockDeductions, err := models.GetStockUpdationParticularsByMedicineID(c.DB, medicineID, false)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockDeductions)
}

func (c *StockController) GetMedicineStockByMedicineID(ctx *fiber.Ctx) error {
	medicineID, err := ctx.ParamsInt("medicine_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "medicine_id", err)
	}
	stock, err := models.GetMedicineStockByMedicineID(c.DB, medicineID)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stock)
}
