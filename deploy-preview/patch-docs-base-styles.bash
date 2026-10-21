#!/usr/bin/env bash

set -euf -o pipefail

function usage {
  cat <<EOF
Patch the grafana/docs-base styles partial so learning-path previews load the
production learn microsite CSS.

The HUGO_ENV=docs branch of layouts/partials/styles.html references
https://grafana.com/web/tw-styles.min.css, which does not exist (production
serves tw-styles.css), and omits themes.min.css. Both are baked into the
docs-base image, so the preview patches them at build time until the upstream
styles.html is fixed and a new docs-base image is published.

Upstream bug: grafana/website layouts/partials/styles.html.

Usage:
  patch-docs-base-styles.bash <hugo-root>
Examples:
  patch-docs-base-styles.bash /hugo
EOF
}

if [[ $# -ne 1 ]]; then
  usage
  exit 2
fi

case "${1}" in
  -h | --help) usage && exit 0 ;;
esac

readonly hugo_root="${1}"
readonly styles="${hugo_root}/layouts/partials/styles.html"

if [[ ! -f "${styles}" ]]; then
  echo "Error: ${styles} not found; cannot patch learn microsite styles." >&2
  exit 1
fi

sed -i 's#/web/tw-styles.min.css#/web/tw-styles.css#g' "${styles}"

if ! grep -q '/web/themes.min.css' "${styles}"; then
  awk '
    { print }
    /grafana\.com\/web\/tw-styles\.css/ && !inserted {
      print "  <link href=\"https://grafana.com/web/themes.min.css\" rel=\"stylesheet\">"
      inserted = 1
    }
  ' "${styles}" > "${styles}.tmp"
  mv "${styles}.tmp" "${styles}"
fi

echo "Patched ${styles}:"
grep -nE '/web/(tw-styles|themes)[^"]*' "${styles}" || true
