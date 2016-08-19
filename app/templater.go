package app

import(
	"html/template"
	"net/http"
	"path"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

var tmpl map[string]*template.Template;
var handlers map[string]string;

const About = "about.html"
const Index = "index.html"
const Subscribe = "subscribe.html"

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	const prefix = "static"
	const base = "static/base.html"

	tmpl = make(map[string]*template.Template)
	tmpl[Index] = template.Must(
		template.ParseFiles(path.Join(prefix, Index), base))
	tmpl[About] = template.Must(
		template.ParseFiles(path.Join(prefix, About), base))
	tmpl[Subscribe] = template.Must(
		template.ParseFiles(path.Join(prefix, Subscribe), base))
}

type TemplateInfo struct {
	User *user.User
	Path string
}

func WriteStatic(w http.ResponseWriter, s string, ti TemplateInfo) {
	if err := tmpl[s].ExecuteTemplate(w, "base", ti); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	WriteStatic(w, Index, TemplateInfo{
		User: u,
		Path: r.URL.Path,
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	WriteStatic(w, About, TemplateInfo{
		User: u,
		Path: r.URL.Path,
	})
}
