package form

import (
	"fmt"
	"net/http"
	"net/mail"
)

type Field struct {
	Name       string
	Required   bool
	Validators []Validator
	Value      string
}

func (f *Field) GetValue() string {
	return f.Value
}

type Validators []Validator

type Validator interface {
	Validate(*Field, string) (bool, string)
}

type Email struct{}

func (e Email) Validate(f *Field, v string) (bool, string) {
	if _, err := mail.ParseAddress(v); err != nil {
		return false, "Require valid email address."
	}
	return true, ""
}

func (f *Field) Validate(r *http.Request) (string, bool, string) {
	v := r.PostFormValue(f.Name)
	if f.Required == true && v == "" {
		return v, false, fmt.Sprintf("%v is required.", f.Name)
	}
	for _, validator := range f.Validators {
		if valid, msg := validator.Validate(f, v); valid == false {
			return v, valid, msg
		}
	}
	return v, true, ""
}
