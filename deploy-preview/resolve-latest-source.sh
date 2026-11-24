#!/usr/bin/env bash

# Find the highest semantic version directory in a given path
function find_latest_version {
  local base_path=$1
  local highest_version=""

  # Look for directories matching semantic version pattern (vX.Y.Z or vX.Y.x)
  while IFS= read -r dir; do
    if [[ -z "${highest_version}" ]]; then
      highest_version="${dir}"
    else
      # Compare versions - extract major, minor, patch
      local current_major current_minor current_patch
      local highest_major highest_minor highest_patch

      # Extract version numbers from current directory
      if [[ "${dir}" =~ ^v([0-9]+)\.([0-9]+)\.(x|[0-9]+)$ ]]; then
        current_major="${BASH_REMATCH[1]}"
        current_minor="${BASH_REMATCH[2]}"
        current_patch="${BASH_REMATCH[3]}"
        [[ "${current_patch}" == "x" ]] && current_patch="9999"
      else
        continue
      fi

      # Extract version numbers from highest so far
      if [[ "${highest_version}" =~ ^v([0-9]+)\.([0-9]+)\.(x|[0-9]+)$ ]]; then
        highest_major="${BASH_REMATCH[1]}"
        highest_minor="${BASH_REMATCH[2]}"
        highest_patch="${BASH_REMATCH[3]}"
        [[ "${highest_patch}" == "x" ]] && highest_patch="9999"
      fi

      # Compare versions
      if (( current_major > highest_major )) || \
          (( current_major == highest_major && current_minor > highest_minor )) || \
          (( current_major == highest_major && current_minor == highest_minor && current_patch > highest_patch )); then
        highest_version="${dir}"
      fi
    fi
  done < <(find "${base_path}" -maxdepth 1 -type d -name 'v*' -printf '%f\n' 2>/dev/null | sort)

  echo "${highest_version}"
}

# Replace <LATEST> placeholder with the actual latest version
function resolve_latest_placeholder {
  local path=$1
  local repo_base=$2

  if [[ "${path}" == *"<LATEST>"* ]]; then
    # Extract the parent directory path before <LATEST>
    local parent_path="${path%/<LATEST>*}"
    local remaining_path="${path#*<LATEST>}"

    # Find the latest version in the parent directory
    local full_parent_path="${repo_base}/${parent_path}"
    local latest_version
    latest_version="$(find_latest_version "${full_parent_path}")"

    if [[ -n "${latest_version}" ]]; then
      echo "${parent_path}/${latest_version}${remaining_path}"
    else
      echo "Warning: No version directories found in ${full_parent_path}, using path as-is" >&2
      echo "${path}"
    fi
  else
    echo "${path}"
  fi
}
