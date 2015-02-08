package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/form"
	"git.ironlabs.com/greg/zenpager/session"
	"git.ironlabs.com/greg/zenpager/template"
)

type Fields []form.Field

var (
	Email = form.Field{
		Name:       "email",
		Required:   true,
		Validators: form.Validators{form.Email{}},
		Value:      "",
	}
	Password = form.Field{
		Name:     "password",
		Required: true,
		Value:    "",
	}
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	if r.Method == "POST" {
		fields := Fields{Email, Password}
		f := form.Validate(r, fields)
		if f.IsValid() == true {
			// authenticate user
			// check for next field and redirect to it
			// otherwise default with dashboard
			user := r.PostFormValue("email")
			session.SetCookie(w, r, USER_KEY, user)
			http.Redirect(w, r, "/dashboard/", http.StatusFound)
		} else {
			ctx.Add("Form", f)
			session.AddMessage(&c, w, r, "Form failed validation!")
		}
	}
	template.Render(c, w, r, "auth/login.html", ctx)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteCookie(w, r, USER_KEY)
	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
