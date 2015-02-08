package form

import (
	"net/http"
)

type Form struct {
	Errors map[string]interface{}
}

func (f *Form) IsValid() bool {
	if len(f.Errors) == 0 {
		return true
	}
	return false
}

func (f *Form) AddError(field, msg string) {
	if f.Errors == nil {
		f.Errors = make(map[string]interface{}, 1)
	}
	f.Errors[field] = msg
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
