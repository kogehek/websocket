package store

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// if exist return false
func checkUniqueDB(table string, field string, value interface{}, db *sql.DB) bool {
	var exists bool
	userSql := "SELECT EXISTS(SELECT 1 FROM " + table + " WHERE " + field + " = $1)"
	row := db.QueryRow(userSql, value)
	row.Scan(&exists)
	return !exists
}

func checkUnique(validate *validator.Validate, m interface{}, db *sql.DB) {
	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		table := getTabaleName(m)
		field := strings.ToLower(fl.FieldName())
		value := fl.Field().String()
		return checkUniqueDB(table, field, value, db)
	})
}

func validModel(m interface{}, db *sql.DB) error {
	validate := validator.New()
	checkUnique(validate, m, db)

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
		}

		return errors.New(sb.String())
	}

	return nil
}

func getTabaleName(m interface{}) (res string) {
	t := reflect.TypeOf(m)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strings.ToLower(res + t.Name() + "s")
}
