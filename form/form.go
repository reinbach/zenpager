package form

import (
	"net/http"
)

type Validation struct {
	Errors map[string]interface{}
}

func (v *Validation) Valid() bool {
	if len(v.Errors) == 0 {
		return true
	}
	return false
}

func (v *Validation) AddError(field, msg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]interface{}, 1)
	}
	v.Errors[field] = msg
}

func NewValidation() *Validation {
	return &Validation{}
}

func Validate(r *http.Request, f []Field) *Validation {
	v := NewValidation()
	for _, field := range f {
		if valid, msg := field.Validate(r); valid == false {
			v.AddError(field.Name, msg)
		}
	}
	return v
}
