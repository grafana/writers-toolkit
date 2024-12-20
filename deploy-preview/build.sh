#!/usr/bin/env bash

docker run \
  -v "${PWD}/dist:/hugo/dist" \
  -v "${PWD}/${SOURCE_DIRECTORY}:/hugo/${WEBSITE_DIRECTORY}" \
  -e INDEX_FILE \
  -e REPO \
  -e WEBSITE_DIRECTORY \
  --rm grafana/docs-base:nightly \
  /bin/bash \
    -c '
if [[ "${INDEX_FILE}" == "true" ]]; then
  echo "Creating custom _index.md" && \
  cat > "/hugo/content/docs/${REPO}/_index.md" <<EOF
---
type: redirect
redirectURL: /docs/${REPO}/latest/
versioned: true
---
EOF
fi
cat "/hugo/content/docs/${REPO}/_index.md"
HUGO_SSI=false hugo --environment=docs --destination=dist/ --baseURL= --minify
'
