package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Validator provides methods for struct validation.
type validator struct{}

// New creates a new instance of Validator.
func New() *validator {
	return &validator{}
}

// Validate validates the fields of a struct based on the `validate` tags and returns errors, if any.
func (v *validator) Validate(input interface{}) map[string]interface{} {
	t := reflect.TypeOf(input)
	val := reflect.ValueOf(input)
	if t.Kind() != reflect.Struct {
		return map[string]interface{}{
			"error": "Input must be a struct",
		}
	}

	errors := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)

		// Check if the field is a string.
		if fieldValue.Kind() != reflect.String {
			continue
		}

		validateTag := field.Tag.Get("validate")
		jsonTag := field.Tag.Get("json")
		errField := getErrorFieldName(field, jsonTag)

		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, ",")
		err := validateField(rules, fieldValue.String(), errField)
		if err != "" {
			errors[errField] = err
		}
	}

	if len(errors) > 0 {
		return map[string]interface{}{
			"error": errors,
		}
	}

	return nil
}

// getErrorFieldName determines the error field name based on the `json` tag or struct field name.
func getErrorFieldName(field reflect.StructField, jsonTag string) string {
	if jsonTag != "" {
		return jsonTag
	}
	return strings.ToLower(field.Name)
}

// validateField performs validation based on rules and returns an error message if validation fails.
func validateField(rules []string, value string, field string) string {
	for _, rule := range rules {
		switch {
		case rule == "required":
			if strings.TrimSpace(value) == "" {
				return "This field is required."
			}
		case rule == "email":
			if !isValidEmail(value) {
				return "Enter a valid email."
			}
		case strings.HasPrefix(rule, "min="):
			if err := validateMinLength(rule, value, field); err != "" {
				return err
			}
		}
	}
	return ""
}

// isValidEmail checks if the value is a valid email address.
func isValidEmail(value string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(strings.TrimSpace(value))
}

// validateMinLength validates if the string meets the minimum length requirement.
func validateMinLength(rule, value, field string) string {
	minValue := strings.TrimPrefix(rule, "min=")
	minLength, err := strconv.Atoi(minValue)
	if err != nil {
		return fmt.Sprintf("Invalid min value for field '%s'", field)
	}
	if len(strings.TrimSpace(value)) < minLength {
		return fmt.Sprintf("'%s' field must be longer than %d characters", field, minLength)
	}
	return ""
}
