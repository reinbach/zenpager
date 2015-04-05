package template

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/utils"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	static_file := r.URL.Path[len(STATIC_ROOT):]
	static_dir := utils.GetAbsDir("template", STATIC_ROOT)
	if len(static_file) != 0 {
		f, err := http.Dir(static_dir).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}

func UpdateTemplateList(tmpls []string) []string {
	d := utils.GetAbsDir("template", TEMPLATE_DIR)
	for i, v := range tmpls {
		tmpls[i] = filepath.Join(d, v)
	}
	return tmpls
}

func Render(c web.C, w http.ResponseWriter, r *http.Request, tmpls []string, ctx *Context) {
	ctx.Add("Static", STATIC_URL)
	ctx.GetMessages(c, w, r)

	tmpl_list := UpdateTemplateList(tmpls)
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Println("template parsing error: ", err)
	}
	err = t.Execute(w, ctx.Values)
	if err != nil {
		log.Println("template executing error: ", err)
	}
}
