package app

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
)

var tmpl map[string]*template.Template;

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	tmpl = make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(
		template.ParseFiles("static/index.html", "static/base.html"))
}

func WriteStatic(w http.ResponseWriter, s string, u *user.User) {
	if err := tmpl[s].ExecuteTemplate(w, "base", u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeIndex(w http.ResponseWriter, u *user.User) {
	WriteStatic(w, "index.html", u)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	writeIndex(w, u);
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
        }
	writeIndex(w, u)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusFound)
		return
	}

	url, err := user.LogoutURL(c, r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}
