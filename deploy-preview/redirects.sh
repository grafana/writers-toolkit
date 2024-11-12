#!/usr/bin/env bash

if ! command -v rg &>/dev/null; then
  echo "ERROR: required binary 'rg' is not in PATH. For installation, refer to https://github.com/BurntSushi/ripgrep#installation." >&2
  exit 2
fi

# find all files that have a meta refresh
rg -e 'http-equiv=refresh content="0; url=/' --glob '*.html' --no-heading --no-ignore --no-line-number --with-filename --fixed-strings -- dist |
# ensure that any lines not matched by sed script are treated as comments by nginx
sed -E 's/^/# /' |
# change into nginx rewrite format
sed -E 's/# dist([^:]+)(\/index.html|(\/[^\/]+.html)):.+<meta http-equiv=refresh content="0; url=(.+)"><\/noscript><\/head><\/html>/rewrite "^\1\3\/\?\$" "\4" permanent;/' |
# output to file for nginx include
cat >dist/redirects.conf
