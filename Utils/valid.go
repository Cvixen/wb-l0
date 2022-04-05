package Utils

import (
	"github.com/go-playground/validator/v10"
	"wb-l0/Model"
)

var validate *validator.Validate

/**
Валидация данных джесон
*/
func ValidateStruct(order *Model.Order) error {
	validate = validator.New()

	err := validate.Struct(order)

	if err != nil {
		return err
	}
	return err
}
