package app

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)


func init() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
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
	WriteStatic(w, "index.html", TemplateInfo{
		User: u,
		Path: "/subscribe",
	})
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
