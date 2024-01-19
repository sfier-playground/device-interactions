package validators

import (
	"fmt"

	"github.com/sifer169966/device-interactions/pkg/apperror"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validators "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validator *validators.Validate
	trans     ut.Translator
}

func New() (*Validator, error) {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	vld := &Validator{
		validator: validators.New(),
		trans:     trans,
	}
	err := enTranslations.RegisterDefaultTranslations(vld.validator, trans)
	if err != nil {
		return nil, err
	}
	return vld, nil
}

func (vld *Validator) ValidateStruct(v interface{}) error {
	err := vld.validator.Struct(v)
	if err == nil {
		return nil
	}
	return apperror.NewBadRequestWithFieldError(vld.translateValidateError(err))
}

func (vld *Validator) translateValidateError(err error) (errs []apperror.ErrorModel) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validators.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(vld.trans))
		errs = append(errs, apperror.ErrorModel{
			Field:   e.Field(),
			Message: apperror.ErrorMessage(translatedErr.Error()),
		})
	}
	return errs
}
