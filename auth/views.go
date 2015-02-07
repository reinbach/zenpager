package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/form"
	"git.ironlabs.com/greg/zenpager/session"
	"git.ironlabs.com/greg/zenpager/template"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	if r.Method == "POST" {
		if err := r.ParseForm(); err == nil {
			v := form.Validate(r, []string{"email", "password"})
			if v.Valid() == true {
				// authenticate user
				// check for next field and redirect to it
				// otherwise default with dashboard
				user := r.PostFormValue("email")
				session.SetCookie(w, r, USER_KEY, user)
				http.Redirect(w, r, "/dashboard/", http.StatusFound)
			} else {
				session.AddMessage(&c, w, r, "Form failed validation!")
			}
		} else {
			session.AddMessage(&c, w, r, "Issue processing form.")
		}
	}
	template.Render(c, w, r, "auth/login.html", ctx)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteCookie(w, r, USER_KEY)
	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
