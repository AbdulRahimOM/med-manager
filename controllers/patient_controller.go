package controllers

import (
	"med-manager/domain/request"
	respcode "med-manager/domain/respcodes"
	"med-manager/domain/response"
	models "med-manager/models"
	"med-manager/utils/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PatientController struct {
	DB *gorm.DB
}

func NewPatientController(db *gorm.DB) *PatientController {
	return &PatientController{DB: db}
}

func (c *PatientController) CreatePatient(ctx *fiber.Ctx) error {
	PatientReq := new(request.PatientReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, PatientReq); !ok {
		return errResponse
	}

	patient := PatientReq.ToPatient()
	if err := patient.Create(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, patient)
}

func (c *PatientController) GetAllPatients(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	patients, err := models.GetAllPatients(c.DB, (page-1)*limit, limit)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, patients)
}

func (c *PatientController) GetPatient(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	patient, err := models.GetPatientByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, patient)
}

func (c *PatientController) UpdatePatient(ctx *fiber.Ctx) error {
	PatientReq := new(request.PatientReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, PatientReq); !ok {
		return errResponse
	}

	patient := PatientReq.ToPatient()
	var err error
	patient.ID, err = ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}

	if err := patient.Update(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, patient)
}

func (c *PatientController) DeletePatient(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	if err := models.DeletePatient(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *PatientController) CreateVisit(ctx *fiber.Ctx) error {
	VisitReq := new(request.VisitReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, VisitReq); !ok {
		return errResponse
	}

	visit := VisitReq.ToVisit()
	if err := visit.Create(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 201, respcode.SUCCESS, visit)
}

func (c *PatientController) GetAllVisits(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	visits, err := models.GetAllVisits(c.DB, (page-1)*limit, limit)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, visits)
}

func (c *PatientController) GetVisit(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	visit, err := models.GetVisitByID(c.DB, id)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, visit)
}

func (c *PatientController) UpdateVisit(ctx *fiber.Ctx) error {
	VisitReq := new(request.VisitReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(ctx, VisitReq); !ok {
		return errResponse
	}

	visit := VisitReq.ToVisit()
	var err error
	visit.ID, err = ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	if err := visit.Update(c.DB); err != nil {
		return response.DBErrorResponse(ctx, err)
	}

	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, visit)
}

func (c *PatientController) DeleteVisit(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "id", err)
	}
	if err := models.DeleteVisit(c.DB, id); err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, nil)
}

func (c *PatientController) GetAllVisitsByPatientID(ctx *fiber.Ctx) error {
	patientID, err := ctx.ParamsInt("patient_id")
	if err != nil {
		return response.InvalidURLParamResponse(ctx, "patient_id", err)
	}
	visits, err := models.GetAllVisitsByPatientID(c.DB, patientID)
	if err != nil {
		return response.DBErrorResponse(ctx, err)
	}
	return response.CreateSuccess(ctx, 200, respcode.SUCCESS, visits)
}
