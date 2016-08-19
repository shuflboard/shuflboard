#!/bin/bash

set -e

# Check for $GO_SDK.
if [ -z "$GO_SDK" ]; then
    echo "Set GO_SDK to the go_appengine directory"
    exit 1
fi

# Add the Strip secret key to app.yaml, if available.
cp fragment.yaml app/app.yaml
if [ -z "$STRIPE_SECRET_KEY" ]; then
    echo "Warning: STRIPE_SECRET_KEY not set, payments won't work."
else
    cat >> app/app.yaml <<EOF

env_variables:
  STRIPE_SECRET_KEY: '$STRIPE_SECRET_KEY'
EOF
fi

# Start the app server.
# Do not remove the `()`s, the dev appserver is poop and you have to kill -9 it
# to shut it down, which kills the terminal it's running in, to boot.
($GO_SDK/dev_appserver.py app)
