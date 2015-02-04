package form

import (
	"fmt"
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

func Validate(r *http.Request, f []string) *Validation {
	var val string
	v := NewValidation()
	for _, field := range f {
		// field.Validate()
		val = r.PostFormValue(field)
		if val == "" {
			v.AddError(field, fmt.Sprintf("%v is required", field))
		}
	}
	return v
}
