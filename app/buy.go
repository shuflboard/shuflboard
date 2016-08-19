package app

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"

	"crypto/tls"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/subscribe", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	ac := appengine.NewContext(r)

	// TODO: make sure user is logged in.
	u := user.Current(ac)

	if r.Method == "GET" {
		WriteStatic(w, "subscribe.html", TemplateInfo{
			User: u,
			Path: "/subscribe",
		})
		log.Infof(ac, "key: %s", stripe.Key)
		return
	}

	// Stripe requires TLS 1.2, AppEngine defaults to TLS 1.0.
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	token := r.FormValue("stripeToken")
	if token == "" {
		http.Error(w, "No token received from Stripe", http.StatusInternalServerError)
		return
	}
	customerParams := &stripe.CustomerParams{
		// TODO: replace with user ID
		Desc: "Customer for elizabeth.wilson@example.com",
		Plan: "alpha-playtest",
	}

	customerParams.SetSource(token)
	sc := client.New(stripe.Key, stripe.NewBackends(httpClient))
	c, err := sc.Customers.New(customerParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := sc.Subs.New(&stripe.SubParams{
		Customer: c.ID,
		Plan: "alpha-playtest",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Infof(ac, "subscription received: %s", s)
}
