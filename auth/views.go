package auth

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/form"
	"git.ironlabs.com/greg/zenpager/session"
	"git.ironlabs.com/greg/zenpager/template"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err == nil {
		v := form.Validate(r, []string{"email", "password"})
		if v.Valid() == true {
			// authenticate user
			// check for next field and redirect to it
			// otherwise default with dashboard
			user := r.PostFormValue("email")
			session.SetCookieHandler(w, r, USER_KEY, user)
			http.Redirect(w, r, "/dashboard/", http.StatusFound)
		} else {
			//
			session.flash(w, r, v.Errors)
		}
	} else {
		fmt.Println("Issue processing form...")
	}
	template.Render(w, "auth/login.html", template.Context{})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteCookieHandler(w, r, USER_KEY)
	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
