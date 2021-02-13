package valid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Model(m interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("gvna", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "gvna" {
			return false
		}
		return true
	})
	var sb strings.Builder
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString("validation failed on field '" + err.Field() + "'")
			sb.WriteString(", condition: " + err.ActualTag())
			if err.Param() != "" {
				sb.WriteString(" { " + err.Param() + " }")
			}

			if err.Value() != nil && err.Value() != "" {
				sb.WriteString(fmt.Sprintf(", actual: %v", err.Value()))
			}
			sb.WriteString("\r\n")
			fmt.Print(sb.String())
		}
	}

	return errors.New(sb.String())
}
