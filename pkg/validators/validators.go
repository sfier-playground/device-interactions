package validators

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/shopspring/decimal"
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
	vld := validators.New()
	trans := createTranslator()
	err := enTranslations.RegisterDefaultTranslations(vld, trans)
	if err != nil {
		return nil, err
	}
	registerCustomType(vld)
	err = registerDecimalGreaterThanOrEqual(vld)
	if err != nil {
		return nil, err
	}
	err = registerDecimalLessThanOrEqual(vld)
	if err != nil {
		return nil, err
	}
	err = registerDeviceNameFormat(vld)
	if err != nil {
		return nil, err
	}
	return &Validator{
		validator: vld,
		trans:     trans,
	}, nil
}

func createTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	return trans
}

func registerCustomType(vld *validators.Validate) {
	vld.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if valuer, ok := field.Interface().(decimal.Decimal); ok {
			return valuer.String()
		}
		return nil
	}, decimal.Decimal{})
}

func registerDecimalGreaterThanOrEqual(vld *validators.Validate) error {
	err := vld.RegisterValidation("dgte", func(fl validators.FieldLevel) bool {
		data, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		actualValue, err := decimal.NewFromString(data)
		if err != nil {
			return false
		}
		expectValue, err := decimal.NewFromString(fl.Param())
		if err != nil {
			return false
		}
		return actualValue.GreaterThanOrEqual(expectValue)
	})
	if err != nil {
		return err
	}
	return nil
}

func registerDecimalLessThanOrEqual(vld *validators.Validate) error {
	err := vld.RegisterValidation("dlte", func(fl validators.FieldLevel) bool {
		data, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		actualValue, err := decimal.NewFromString(data)
		if err != nil {
			return false
		}
		expectValue, err := decimal.NewFromString(fl.Param())
		if err != nil {
			return false
		}
		return actualValue.LessThanOrEqual(expectValue)
	})
	if err != nil {
		return err
	}
	return nil

}

func registerDeviceNameFormat(vld *validators.Validate) error {
	err := vld.RegisterValidation("device_name", func(fl validators.FieldLevel) bool {
		deviceNameConvention := "^[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+$"
		deviceNameRegexp := regexp.MustCompile(deviceNameConvention)
		return deviceNameRegexp.MatchString(fl.Field().String())
	})
	if err != nil {
		return err
	}
	return nil

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
