# Go AppEngine app rules.

_go_file_types = FileType([".go"])

def _impl(ctx):
    ctx.file_action(
        output = ctx.outputs.executable,
        content = """
#!/bin/bash
set -e

if [ -d "$0.runfiles" ]; then
    cd "$0.runfiles/com_shuflboard";
fi

# Do not remove the `()`s, the dev appserver is poop and you have to kill -9 it
# to shut it down, which kills the terminal it's running in, to boot.
({appserver} {name})
""".format(
            appserver = ctx.file._dev_appserver.path,
            name = ctx.label.name),
        executable = True,
    )

    runfiles = ctx.runfiles(
        files = ctx.files.srcs + ctx.files._appengine_sdk + [
            ctx.file.config, ctx.file._dev_appserver],
        collect_default = True,
    )
    return struct(runfiles = runfiles)

go_app = rule(
    implementation = _impl,
    attrs = {
        "config" : attr.label(allow_files = True, single_file = True),
        "srcs" : attr.label_list(allow_files = _go_file_types),
        "data" : attr.label_list(allow_files = True, cfg = DATA_CFG),
        "_appengine_sdk" : attr.label(
            default = Label("@go_appengine//:sdk")),
        "_dev_appserver" : attr.label(
            default = Label("@go_appengine//:dev_appserver.py"),
            allow_files = True,
            single_file = True),
    },
    executable = True,
)

def go_repositories():
    native.new_http_archive(
        name = "go_appengine",
        url = "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_darwin_amd64-1.9.40.zip",
        strip_prefix = "go_appengine",
        build_file_content = """
exports_files(["dev_appserver.py"])

filegroup(
    name = "sdk",
    srcs = glob(["**"], exclude = ["dev_appserver.py"]),
    visibility = ["//visibility:public"],
)
"""
    )
