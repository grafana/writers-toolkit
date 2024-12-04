#!/usr/bin/env bash

docker run \
  -v "${PWD}/dist:/hugo/dist" \
  -v "${PWD}/${source_directory}:/hugo/${website_directory}" \
  -e index_file \
  -e repo \
  -e website_directory \
  --rm grafana/docs-base:latest \
  /bin/bash \
    -c '
if [[ "${index_file}" == "true" ]]; then
  mkdir -p "/hugo/${website_directory}/content/docs/${repo}" && \
  cat > "/hugo/${website_directory}/content/docs/${repo}/_index.md" <<EOF
type: redirect
redirectURL: /docs/${repo}/latest/
versioned: true
EOF
fi
cat "/hugo/${website_directory}/content/docs/${repo}/_index.md"
HUGO_SSI=false hugo --environment=docs --destination=dist/ --baseURL= --minify
'
