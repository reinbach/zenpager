package form

import (
	"net/http"
)

type Form struct {
	Errors []string
	Fields map[string]interface{}
}

func (f *Form) IsValid() bool {
	if len(f.Errors) == 0 {
		return true
	}
	return false
}

func (f *Form) AddError(msg string) {
	f.Errors = append(f.Errors, msg)
}

func (f *Form) AddField(field Field) {
	if f.Fields == nil {
		f.Fields = make(map[string]interface{}, 1)
	}
	f.Fields[field.Name] = field
}

func (f *Form) GetValue(field string) string {
	v := f.Fields[field].(Field)
	return v.GetValue()
}

func NewForm() *Form {
	return &Form{}
}

func Validate(r *http.Request, f []Field) *Form {
	form := NewForm()
	if err := r.ParseForm(); err == nil {
		for _, field := range f {
			value, valid, msg := field.Validate(r)
			if valid == false {
				form.AddError(msg)
			}
			field.Value = value
			form.AddField(field)
		}
	} else {
		form.AddError("Issue processing form data.")
	}
	return form
}
