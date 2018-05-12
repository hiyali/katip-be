package main

import (
  "strings"

  "github.com/asaskevich/govalidator"

  "github.com/hiyali/katip-be/config"
)

type (
  Validator struct {}

  ValidationErrors []error
)

func init() {
  govalidator.SetFieldsRequiredByDefault(true)
}

func errorLocalize(gv govalidator.Errors) ValidationErrors {
  var ve ValidationErrors
  for _, e := range gv {
    gve := e.(govalidator.Error)
    ve = append(ve, &config.JsonValidationError{
      Name: gve.Name,
      Err: gve.Err.Error(),
      Validator: gve.Validator,
    })
  }
  return ve
}

func (ve ValidationErrors) Error() string {
  var errs []string
  for _, e := range ve {
    errs = append(errs, e.Error())
  }
  return strings.Join(errs, ";")
}

func (cv *Validator) Validate (i interface{}) error {
  _, err := govalidator.ValidateStruct(i)
  if err != nil {
    errs := err.(govalidator.Errors)
    localErrs := errorLocalize(errs)
    return localErrs
  }
  return nil
}
