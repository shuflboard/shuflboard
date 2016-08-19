def _impl(repository_ctx):
    if not "STRIPE_SECRET_KEY" in repository_ctx.os.environ:
        fail("Environment variable STRIPE_SECRET_KEY must be set")

    key = repository_ctx.os.environ["STRIPE_SECRET_KEY"]
    repository_ctx.file("secret-key", key)
    repository_ctx.file("BUILD", """
exports_files(["secret-key"])
""")

gen_key = repository_rule(
    implementation = _impl,
    local = True,
)
