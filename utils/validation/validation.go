package validation

import (
	"fmt"
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

func validateJSONRequestDetailed(req interface{}) []InvalidField {

	Response := []InvalidField{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		// Get the JSON tag name using reflection
		jsonFieldName := getJSONTagName(req, err.Field())

		e := InvalidField{
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
