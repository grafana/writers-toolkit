#!/usr/bin/env bash

set -euf -o pipefail

if [[ -n "${RUNNER_DEBUG+x}" ]]; then
  set -x
fi

# All input comes from environment variables that are capitalized by convention in this script.
SOURCES="${SOURCES:-[]}"

if ! jq -e . <<<"${SOURCES}" >/dev/null; then
  echo "SOURCES environment variable is not valid JSON"
fi

# Clone a repository to a specific directory.
function clone {
  local repo=$1
  local directory=$2

  if [[ -d "${directory}" ]]; then
    echo "Directory ${directory} already exists, skipping clone"

    return
  fi

  gh repo clone "grafana/${repo}" "${directory}"
}

# Check out all the source repositories defined in SOURCES.
function check_out_sources {
  index_files=()
  while read -r repo; do
    if [[ -z "${repo}" ]]; then
      continue
    fi

    clone "${repo}" "src/${repo}"

  done < <(jq -r '.[].repo' <<<"${SOURCES}")
}

# Create an entrypoint script for the container.
function prepare_entrypoint {
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
}

# Add mounts to the volumes variable.
function configure_volumes {
  volumes=("--volume=${PWD}/dist:/hugo/dist" "--volume=${tempfile}:/entrypoint:z")
  while read -r repo source_directory website_directory; do
    volumes+=("--volume=${PWD}/src/${repo}/${source_directory}:/hugo/${website_directory}:z")
  done < <(jq -r '.[] | "\(.repo) \(.source_directory) \(.website_directory)"' <<<"${SOURCES}")
}

# Prepare the command to run the container.
function prepare_cmd {
  cat <<EOF
docker run \
${volumes[@]} \
--rm grafana/docs-base:latest \
/entrypoint
EOF
}

check_out_sources
prepare_entrypoint
declare -a volumes; configure_volumes; readonly volumes
cmd=$(prepare_cmd)

echo "${cmd}"
${cmd}
