package response

import (
	"fmt"
	respcode "med-manager/domain/respcodes"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	HttpStatusCode int         `json:"-"`
	Status         bool        `json:"status"`
	ResponseCode   string      `json:"resp_code"`
	Error          error       `json:"-"` //will be marshalled to string when WriteToJSON is called
	Data           interface{} `json:"data,omitempty"`
}

func CreateError(ctx *fiber.Ctx, statusCode int, respcode string, err error) error {
	return Response{
		HttpStatusCode: statusCode,
		Status:         false,
		ResponseCode:   respcode,
		Error:          err,
	}.WriteToJSON(ctx)
}

func CreateSuccess(ctx *fiber.Ctx, statusCode int, respcode string, data interface{}) error {
	return Response{
		HttpStatusCode: statusCode,
		Status:         true,
		ResponseCode:   respcode,
		Data:           data,
	}.WriteToJSON(ctx)
}

func DBErrorResponse(ctx *fiber.Ctx, err error) error {
	return Response{
		HttpStatusCode: 500,
		Status:         false,
		ResponseCode:   respcode.DB_ERROR,
		Error:          err,
	}.WriteToJSON(ctx)
}

func InvalidURLParamResponse(ctx *fiber.Ctx, param string, err error) error {
	return CreateError(ctx, http.StatusBadRequest, respcode.INVALID_URL_PARAM, fmt.Errorf("error parsing %v from url: %w", param, err))
}

func BugResponse(ctx *fiber.Ctx, err error) error {
	return CreateError(ctx, http.StatusInternalServerError, respcode.BUG, fmt.Errorf("bug found, notify BE: %w", err))
}

func UnauthorizedResponse(ctx *fiber.Ctx, err error) error {
	return CreateError(ctx, http.StatusUnauthorized, respcode.UNAUTHORIZED, fmt.Errorf("unauthorized: %w", err))
}

type custError struct {
	Response
	Error string `json:"error"`
}

func (resp Response) WriteToJSON(c *fiber.Ctx) error {
	if resp.Error == nil {
		fmt.Println("resp.HttpStatusCode:", resp.HttpStatusCode)
		return c.Status(resp.HttpStatusCode).JSON(resp)
	}
	newCustError := custError{
		Response: resp,
	}
	if resp.Error != nil {
		newCustError.Error = resp.Error.Error()
	}

	return c.Status(resp.HttpStatusCode).JSON(newCustError)
}
