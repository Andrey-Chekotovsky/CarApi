package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validator.New()
	validate.RegisterValidation("SerialNum", validateSerialNum)
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func validateSerialNum(f1 validator.FieldLevel) bool {
	//format aaa-11
	re := regexp.MustCompile(`[a-z]+-[0-9]+`)
	serialNum := re.FindAllString(f1.Field().String(), -1)
	if len(serialNum) == 1 {
		return true
	}
	return false
}
