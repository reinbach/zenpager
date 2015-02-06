package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/form"
	"git.ironlabs.com/greg/zenpager/session"
	"git.ironlabs.com/greg/zenpager/template"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext(w, r)
	template.Render(w, "auth/login.html", ctx)
}

func Authenticate(c web.C, w http.ResponseWriter, r *http.Request) {
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
			session.AddMessage(w, r, "Form failed validation!")
		}
	} else {
		session.AddMessage(w, r, "Issue processing form.")
	}
	http.Redirect(w, r, Route("/login/"), http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteCookie(w, r, USER_KEY)
	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
