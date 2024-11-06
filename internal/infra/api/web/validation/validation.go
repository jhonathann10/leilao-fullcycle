package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorEn "github.com/go-playground/validator/v10/translations/en"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ := enTransl.GetTranslator("en")
		validatorEn.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError(jsonErr.Error())
	} else if errors.As(validationErr, &jsonValidation) {
		errorCauses := []rest_err.Causes{}
		for _, e := range validationErr.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}
		return rest_err.NewBadRequestError("invalid filed values", errorCauses...)
	} else {
		return rest_err.NewInternalServerError("internal server error")
	}
}
