package validation

import (
	"fmt"
	"med-manager/domain/response"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func getJSONTagName(req interface{}, fieldName string) string {
	val := reflect.TypeOf(req)

	// Check if the value passed is a pointer and get the element type
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Find the struct field by name and return the JSON tag value
	field, found := val.FieldByName(fieldName)
	if !found {
		return fieldName // Return the field name if no JSON tag is found
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldName // Return the field name if the JSON tag is not defined
	}
	return jsonTag
}

// Function to get the 'form' tag name of a struct field
func getFormTagName(req interface{}, fieldName string) string {
	val := reflect.TypeOf(req)

	// Check if the value passed is a pointer and get the element type
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Find the struct field by name and return the 'form' tag value
	field, found := val.FieldByName(fieldName)
	if !found {
		return fieldName // Return the field name if no 'form' tag is found
	}

	formTag := field.Tag.Get("form")
	if formTag == "" {
		return fieldName // Return the field name if the 'form' tag is not defined
	}
	return formTag
}
func validateJSONRequestDetailed(req interface{}) []response.InvalidField {

	Response := []response.InvalidField{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		// Get the JSON tag name using reflection
		jsonFieldName := getJSONTagName(req, err.Field())

		e := response.InvalidField{
			FailedField: jsonFieldName, // Use JSON field name instead of Go field name
			Tag:         err.Tag(),
			Value:       err.Value(),
		}

		message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)
		fmt.Println("validation fail message: ", message)

		Response = append(Response, e)
	}
	return Response
}
func validateFormDataRequestDetailed(req interface{}) []response.InvalidField {

	Response := []response.InvalidField{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		// Get the 'form' tag name using reflection
		formFieldName := getFormTagName(req, err.Field())

		e := response.InvalidField{
			FailedField: formFieldName, // Use 'form' tag field name instead of Go field name
			Tag:         err.Tag(),
			Value:       err.Value(),
		}

		message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)
		fmt.Println("validation fail message: ", message)

		Response = append(Response, e)
	}
	return Response
}
