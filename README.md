# Shuflboard

## Running locally

This is a go appengine app that can be built with Bazel.  To try out locally,
[install Bazel](http://bazel.io/docs/install.html) and set the Stripe API key
environment variable, then run:

```
$ STRIPE_SECRET_KEY=sk_test_... bazel build //app
INFO: Found 1 target...
Target //app:app up-to-date:
  bazel-bin/app/app
INFO: Elapsed time: 8.682s, Critical Path: 2.53s
$ bazel-bin/app/app
```

The Stripe secret key can be found on the dashboard under [Your account ->
account settings -> API keys](https://dashboard.stripe.com/account/apikeys). Use
the one for the test instance that starts with "sk_test_".

Bazel will download and install the Go AppEngine SDK, so it will take a while to
run the first time you run it.

## Deploying

Make sure the Stripe environment variable is set to the production secret key in
`app/app.yaml`. Run:

```
$ /path/to/go_appengine/appcfg.py update app
```
