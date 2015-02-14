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
			user := User{
				Email:    f.GetValue("email"),
				Password: f.GetValue("password"),
			}
			if user.Login(c) {
				session.SetCookie(w, r, USER_KEY, user.Email)
				http.Redirect(w, r, "/dashboard/", http.StatusFound)
			} else {
				session.AddMessage(&c, w, r, "Invalid User")
			}
		}
		ctx.Add("Form", f)
	}
	template.Render(c, w, r, "auth/login.html", ctx)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteCookie(w, r, USER_KEY)
	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
