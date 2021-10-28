package requests

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goblong/pkg/model"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {

	// not_exists:users,email
	govalidator.AddCustomRule("not_exists",
		func(field string, rule string, message string, value interface{}) error {
			rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

			tableName := rng[0]
			dbFiled := rng[1]
			val := value.(string)

			var count int64
			model.DB.Table(tableName).Where(dbFiled+" = ?", val).Count(&count)
			if count != 0 {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("%v Aready taked", value)
			}
			return nil
		})

	// chinese character max_cn:8
	govalidator.AddCustomRule("max_cn",
		func(field string, rule string, message string, value interface{}) error {
			valLength := utf8.RuneCountInString(value.(string))
			ruleLength, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
			if valLength > ruleLength {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("Length cannot exceed %d character", ruleLength)
			}
			return nil
		})

	// chinese character min_cn:2
	govalidator.AddCustomRule("min_cn",
		func(field string, rule string, message string, value interface{}) error {
			valLength := utf8.RuneCountInString(value.(string))
			ruleLength, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
			if valLength < ruleLength {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("Length cannot less than %d character", ruleLength)
			}
			return nil
		})

}
