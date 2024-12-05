package validation

import (
	"log"
	"med-manager/domain/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	bindingErrCode      = "BINDING_ERROR"
	validationErrCode   = "VALIDATION_ERROR"
	queryBindingErrCode = "URL QUERY BINDING ERROR"
)

type ValidationErrorResponse struct {
	Status       bool           `json:"status"`
	ResponseCode string         `json:"resp_code"`
	Errors       []InvalidField `json:"errors"`
}

type InvalidField struct {
	FailedField string      `json:"field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

// BindAndValidateRequest binds and validates the request.
// Req should be a pointer to the request struct.
func BindAndValidateJSONRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   bindingErrCode,
			Error:          err,
		}.WriteToJSON(c)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}
	return true, nil
}

func BindAndValidateURLQueryRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.QueryParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   queryBindingErrCode,
			Error:          err,
		}.WriteToJSON(c)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}
	return true, nil
}
