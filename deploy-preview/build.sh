#!/usr/bin/env bash

set -euf -o pipefail

# All input comes from environment variables that are capitalized by convention in this script.
SOURCES="${SOURCES:-[]}"

if ! jq -e . <<<"${SOURCES}" >/dev/null; then
  echo "SOURCES environment variable is not valid JSON"
fi

tempfile="$(mktemp -t deploy-preview.XXX)"
# Pull out the relevant fields from the JSON because the grafana/docs-base image doesn't have jq.
index_files=()
while read -r index_file relative_prefix; do
  if [[ "${index_file}" != 'null' ]]; then
    index_files+=("${index_file}:${relative_prefix}")
  fi
done < <(jq -r '.[] | "\(.index_file) \(.relative_prefix)"' <<<"${SOURCES}")
cat <<EOSCRIPT >"${tempfile}"
#!/usr/bin/env bash

# Create an index file to redirect a project root to the correct versioned URL.
for pair in ${index_files[@]}; do
  IFS=':' read -r index_file relative_prefix <<<"\${pair}"
  dst="/hugo/\${index_file}"
  parent="\${dst%/*}"
  echo "Creating custom index: \${dst} -> \${relative_prefix}"

  title="\${dst%/*}"
  while [[ -n "\${parent}" ]]; do
    if [[ ! -f "\${parent}/_index.md" ]]; then
       cat > "\${parent}/_index.md" <<EOPARENT
---
title: \${title}
---

# \${title}

{{< section >}}
EOPARENT
    fi
    parent="\${parent%/*}"
  done

  cat > "\${dst}" <<EOINDEX
---
type: redirect
redirectURL: \${relative_prefix}
versioned: true
---
EOINDEX
  cat "\${dst}"
done

HUGO_SSI=false hugo --environment=docs --destination=dist/ --baseURL= --minify
EOSCRIPT
chmod +x "${tempfile}"

volumes=("--volume=${PWD}/dist:/hugo/dist" "--volume=${tempfile}:/entrypoint:z")
while read -r source_directory website_directory; do
  volumes+=("--volume=${PWD}/${source_directory}:/hugo/${website_directory}:z")
done < <(jq -r '.[] | "\(.source_directory) \(.website_directory)"' <<<"${SOURCES}")

IFS='' read -r cmd <<EOF
docker run \
  ${volumes[@]} \
  --rm grafana/docs-base:latest \
  /entrypoint
EOF

echo "${cmd}"
${cmd}
