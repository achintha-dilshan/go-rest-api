package validator

import (
	"testing"
)

func TestValidateRequiredField(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name" validate:"required"`
	}

	v := New()
	result := v.Validate(TestStruct{})

	expectedError := "This field is required."
	if err, ok := result["error"].(map[string]string); ok {
		if err["name"] != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err["name"])
		}
	} else {
		t.Errorf("Expected validation errors, got none")
	}
}

func TestValidateEmailField(t *testing.T) {
	type TestStruct struct {
		Email string `json:"email" validate:"email"`
	}

	v := New()
	result := v.Validate(TestStruct{Email: "invalid-email"})

	expectedError := "Enter a valid email."
	if err, ok := result["error"].(map[string]string); ok {
		if err["email"] != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err["email"])
		}
	} else {
		t.Errorf("Expected validation errors, got none")
	}
}

func TestValidateMinLengthField(t *testing.T) {
	type TestStruct struct {
		Username string `json:"username" validate:"min=5"`
	}

	v := New()
	result := v.Validate(TestStruct{Username: "abc"})

	expectedError := "'username' field must be longer than 5 characters"
	if err, ok := result["error"].(map[string]string); ok {
		if err["username"] != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err["username"])
		}
	} else {
		t.Errorf("Expected validation errors, got none")
	}
}

func TestValidateValidInput(t *testing.T) {
	type TestStruct struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Username string `json:"username" validate:"min=5"`
	}

	v := New()
	input := TestStruct{
		Name:     "John",
		Email:    "john.doe@example.com",
		Username: "johndoe",
	}
	result := v.Validate(input)

	if result != nil {
		t.Errorf("Expected no validation errors, got '%v'", result)
	}
}
