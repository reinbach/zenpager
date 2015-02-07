package form

import (
	"fmt"
	"net/http"
)

type Field struct {
	Name       string
	Required   bool
	Validators []Validator
}

type Validators []Validator

type Validator interface {
	Validate(*Field, string) (bool, string)
}

type Email struct{}

func (e Email) Validate(f *Field, v string) (bool, string) {
	if v[len(v)-3:] != ".com" {
		return false, fmt.Sprintf("%v requires .com", f.Name)
	}
	return true, ""
}

func (f *Field) Validate(r *http.Request) (bool, string) {
	v := r.PostFormValue(f.Name)
	if f.Required == true && v == "" {
		return false, fmt.Sprintf("%v is required", f.Name)
	}
	for _, validator := range f.Validators {
		if valid, msg := validator.Validate(f, v); valid == false {
			return valid, msg
		}
	}
	return true, ""
}
