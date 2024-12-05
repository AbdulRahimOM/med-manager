package controllers

import (
	respcode "med-manager/domain/respcodes"
	"med-manager/domain/response"
	"med-manager/models"
	"med-manager/utils/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MedicineController struct {
	DB *gorm.DB
}

func NewMedicineController(db *gorm.DB) *MedicineController {
	return &MedicineController{DB: db}
}

func (c *MedicineController) CreateMedicine(ctx *fiber.Ctx) error {
	medicine := new(models.Medicine)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medicine); !ok {
		return errResponse
	}

	if err := medicine.Create(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, medicine)
}

func (c *MedicineController) GetAllMedicines(ctx *fiber.Ctx) error {
	medicines, err := models.GetAllMedicines(c.DB)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medicines)
}

func (c *MedicineController) GetMedicine(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	medicine, err := models.GetMedicineByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medicine)
}

func (c *MedicineController) UpdateMedicine(ctx *fiber.Ctx) error {
	medicine := new(models.Medicine)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medicine); !ok {
		return errResponse
	}

	if err := medicine.Update(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medicine)
}

func (c *MedicineController) DeleteMedicine(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	if err := models.DeleteMedicine(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

// func (c *MedicineController) GetAllMedTypes(ctx *fiber.Ctx) error {
// 	medTypes, err := models.GetAllMedTypes(c.DB)
// 	if err != nil {
// 		return response.DBErrorResponse(ctx, err)
// 	}
// 	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medTypes)
// }

// func (c *MedicineController) GetMedType(ctx *fiber.Ctx) error {
// 	id, err := ctx.ParamsInt("id")
// 	if err != nil {
// 		return response.InvalidURLParamResponse(ctx, "id", err)
// 	}
// 	medType, err := models.GetMedTypeByID(c.DB, id)
// 	if err != nil {
// 		return response.DBErrorResponse(ctx, err)
// 	}

// 	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medType)
// }

func (c *MedicineController) GetStockAdditionsByMedicineID(ctx *fiber.Ctx) error {
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

func (c *MedicineController) GetStockDeductionsByMedicineID(ctx *fiber.Ctx) error {
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

func (c *MedicineController) GetStockUpdation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	stockUpdation, err := models.GetStockUpdationByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, stockUpdation)
}

func (c *MedicineController) UpdateStockUpdation(ctx *fiber.Ctx) error {
	stockUpdations := new(models.UpdateStockUpdateRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockUpdations); !ok {
		return errResponse
	}

	if err := models.UpdateParticularsInAnStockUpdation(c.DB, stockUpdations.ID, stockUpdations.StockChanges); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *MedicineController) DeleteStockUpdation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	if err := models.DeleteStockUpdation(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *MedicineController) DeleteStockUpdationParticulars(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	medID, err := ctx.ParamsInt("medicine_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "medicine_id", err)
	}

	if err := models.DeleteStockUpdationParticulars(c.DB, id, medID); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

// func (c *MedicineController) DeductFromStock(ctx *fiber.Ctx) error {
// 	stockDeductions := new(models.StockUpdateRequest)
// 	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockDeductions); !ok {
// 		return errResponse
// 	}

// 	if err := models.DeductFromStock(c.DB, stockDeductions); err != nil {
// 		return response.DBErrorResponse(ctx, err)
// 	}

// 	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
// }

// func (c *MedicineController) AddToStock(ctx *fiber.Ctx) error {
// 	stockAdditions := new(models.StockUpdateRequest)
// 	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, stockAdditions); !ok {
// 		return errResponse
// 	}

// 	if err := models.AddToStock(c.DB, stockAdditions); err != nil {
// 		return response.DBErrorResponse(ctx, err)
// 	}

// 	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
// }
