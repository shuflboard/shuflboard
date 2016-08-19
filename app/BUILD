load("//skylark:go_appengine.bzl", "go_app")

go_app(
    name = "app",
    srcs = glob(["*.go"]),
    config = ":yaml",
    data = glob(["static/**"]),
)

genrule(
    name = "yaml",
    srcs = [
        "fragment.yaml",
        "@stripe_key_repo//:secret-key",
    ],
    outs = ["app.yaml"],
    cmd = """
KEY=$$(cat $(location @stripe_key_repo//:secret-key));
cp $(location fragment.yaml) $@;
cat >> $@ <<EOF

env_variables:
  STRIPE_SECRET_KEY: '$$KEY'
EOF
""",
)
