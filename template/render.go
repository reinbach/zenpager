package template

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	static_file := r.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateTemplateList(tmpl string) []string {
	tmpl_list := []string{
		fmt.Sprintf("%sbase.html", TEMPLATE_DIR),
	}

	// for each dir down from base template dir
	// check for a base.html file, if there is
	// one then add it to the list of templates
	// to be parsed
	bases := strings.Split(tmpl, "/")
	base := ""
	for _, b := range bases {
		base = filepath.Join(base, b)
		basefile, _ := filepath.Abs(filepath.Join(TEMPLATE_DIR, base,
			"base.html"))
		f, _ := filepath.Glob(basefile)
		if f != nil {
			tmpl_list = append(tmpl_list, f[0])
		}
	}
	tmpl_list = append(tmpl_list, filepath.Join(TEMPLATE_DIR, tmpl))
	return tmpl_list
}

func Render(w http.ResponseWriter, tmpl string, ctx Context) {
	ctx.Add("Static", STATIC_URL)
	tmpl_list := CreateTemplateList(tmpl)
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, ctx.Values)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
