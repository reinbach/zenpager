package form

import (
	"net/http"
)

type Form struct {
	Errors map[string]interface{}
}

func (v *Form) IsValid() bool {
	if len(v.Errors) == 0 {
		return true
	}
	return false
}

func (v *Form) AddError(field, msg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]interface{}, 1)
	}
	v.Errors[field] = msg
}

func NewForm() *Form {
	return &Form{}
}

func Validate(r *http.Request, f []Field) *Form {
	form := NewForm()
	if err := r.ParseForm(); err != nil {
		for _, field := range f {
			if valid, msg := field.Validate(r); valid == false {
				form.AddError(field.Name, msg)
			}
		}
	} else {
		form.AddError("all", "Issue processing form data.")
	}
	return form
}
