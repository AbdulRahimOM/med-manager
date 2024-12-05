package controllers

import (
	"fmt"
	"med-manager/domain/request"
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
	medicineReq := new(request.MedicineRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medicineReq); !ok {
		return errResponse
	}

	medicine := medicineReq.ToMedicine()

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
	medicineReq := new(request.MedicineRequest)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medicineReq); !ok {
		return errResponse
	}

	medicine := medicineReq.ToMedicine()
	var err error
	medicine.ID, err = ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
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

func (c *MedicineController) GetAllMedTypes(ctx *fiber.Ctx) error {
	medTypes, err := models.GetAllMedTypes(c.DB)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medTypes)
}

func (c *MedicineController) GetMedType(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	medType, err := models.GetMedTypeByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medType)
}

func (c *MedicineController) CreateMedType(ctx *fiber.Ctx) error {
	medType := new(models.MedType)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medType); !ok {
		return errResponse
	}

	if err := medType.Create(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, medType)
}

func (c *MedicineController) UpdateMedType(ctx *fiber.Ctx) error {
	medType := new(models.MedType)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, medType); !ok {
		return errResponse
	}

	var err error
	medType.ID, err = ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	fmt.Println("ID: ", medType.ID)

	if err := medType.Update(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, medType)
}

func (c *MedicineController) DeleteMedType(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	if err := models.DeleteMedType(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}
