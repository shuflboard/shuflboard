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

func WriteStatic(w http.ResponseWriter, s string, u *user.User) {
	if err := tmpl[s].ExecuteTemplate(w, "base", u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	WriteStatic(w, Index, u)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	WriteStatic(w, About, u)
}
