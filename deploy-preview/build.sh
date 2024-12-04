#!/usr/bin/env bash

docker run \
  -v "${PWD}/dist:/hugo/dist" \
  -v "${PWD}/${{ inputs.source_directory }}:/hugo/${{ inputs.website_directory }}" \
  --rm grafana/docs-base:latest \
  /bin/bash \
    -c 'HUGO_SSI=false hugo --environment=docs --destination=dist/ --baseURL= --minify'