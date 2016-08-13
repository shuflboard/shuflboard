# Shuflboard

## Running locally

This is a go appengine app.  To try out locally, download the go appengine SDK
and then run:

```
$ (/path/to/go_appengine/dev_appserver.py app)
```

_Do not remove the `()`s, the dev appserver is poop and you have to kill -9 it
to shut it down, which kills the terminal it's running in, to boot._

## Deploying

Run:

```
$ /path/to/go_appengine/appcfg.py update app
```
