package template

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
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

func TemplateDirAbs() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), TEMPLATE_DIR)
}

func CreateTemplateList(tmpl string) []string {
	d := TemplateDirAbs()
	tmpl_list := []string{
		filepath.Join(d, "base.html"),
	}

	// for each dir down from base template dir
	// check for a base.html file, if there is
	// one then add it to the list of templates
	// to be parsed
	bases := strings.Split(tmpl, "/")
	base := ""
	for _, b := range bases {
		base = filepath.Join(base, b)
		f, _ := filepath.Glob(filepath.Join(d, base, "base.html"))
		if f != nil {
			tmpl_list = append(tmpl_list, f[0])
		}
	}
	tmpl_list = append(tmpl_list, filepath.Join(d, tmpl))
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