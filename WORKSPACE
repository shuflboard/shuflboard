workspace(name = "com_shuflboard")

load('//skylark:go_appengine.bzl', 'go_repositories')
go_repositories()

# This is a hack around not having environment variables.
load('//skylark:stripe.bzl', 'gen_key')
gen_key(name = "stripe_key_repo")
