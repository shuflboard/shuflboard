# Shuflboard

## Running locally

This is a go appengine app.  To try out locally, download the go appengine SDK,
set the Stripe API key environment variable (if you want to test out payments),
and then run:

```
$ export GO_SDK=/path/to/go_appengine
$ STRIPE_SECRET_KEY=sk_test_... ./runner.sh
```

The Stripe secret key can be found on the dashboard under [Your account ->
account settings -> API keys](https://dashboard.stripe.com/account/apikeys). Use
the one for the test instance that starts with "sk_test_".

## Deploying

Make sure the Stripe environment variable is set to the production secret key in
`app/app.yaml`. Run:

```
$ /path/to/go_appengine/appcfg.py update app
```
