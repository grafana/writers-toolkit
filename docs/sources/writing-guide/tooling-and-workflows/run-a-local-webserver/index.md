---
title: Run a local documentation webserver with the make-docs script
menuTitle: Run a local documentation webserver
description: Run a local documentation web server with the make-docs script.
---

# Run a local documentation webserver with the make-docs script

The [`make-docs`](https://github.com/grafana/writers-toolkit/blob/main/scripts/make-docs) script can be used to mount multiple documentation sets in a single build.
It is used by the `make docs` Make target in project repositories and can also be run separately.

## Examples

The following examples assume:

1. Your working directory is the root of a project repository.
1. You have a copy of the `make-docs` script in that working directory.
1. The script is marked as executable.
1. You have a checkout of each of the project repositories used in the examples.

**Build Mimir, the `mimir-distributed` Helm chart, and GEM**: `./make-docs metrics`

**Build Loki and GEL**: `./make-docs logs`

**Build Tempo and GET**: `./make-docs traces`

**Build Grafana and Grafana Cloud:** `./make-docs grafana grafana-cloud`

**Build Grafana and Mimir as "next":** `./make-docs grafana:next mimir:next`

**Mount `v9.3.x` and "latest" versions of Grafana:**:

From the `grafana` repository:

```console
$ git worktree add v9.3.x origin/v9.3.x
$ export GRAFANA_REPO="$(pwd)"
```

From the `technical-documentation` repository:

```console
$ ./make-docs grafana "grafana:v9.3.x:${GRAFANA_REPO}/v9.3.x"
```

To remove the worktree created in the `grafana` repository, run the following command from the `grafana` repository:

```console
$ git worktree remove v9.3.x
```

## Reference

Each argument to the `make-docs` script is a project to be mounted into the local build.
Each argument has four fields separated by colons (`:`) and optional fields can be omitted.

`<PROJECT>[:VERSION[:REPOSITORY[:DIR]]]`

- `PROJECT`: is the name of the project as it appears in the website URL path.

  For Grafana, the `PROJECT` is `grafana` (https://grafana.com/docs/grafana/)
  For Grafana Cloud, the `PROJECT` is `grafana-cloud` (https://grafana.com/docs/grafana-cloud/).

  Pseudo projects mount multiple source directories.

  - `logs`: mounts Loki and GEL directories.
  - `metrics`: mounts Mimir, the `mimir-distributed` Helm chart, and GEM directories.
  - `traces`: mounts Tempo and GET directories.

  **Note:** pseudo projects do not support the `REPOSITORY` or `DIR` fields.

- `VERSION`: is the name of the version directory under which the project documentation should be mounted.

  The `VERSION` field is optional and defaults to `latest` for versioned projects and is empty for unversioned projects.

- `REPOSITORY` is the the name of the local directory within `REPOS_PATH` for that project.

  The `REPOSITORY` field is optional and defaults to the scripts internal mapping of project names to repository names.
  For most projects, this is the same as the project name.

- The `DIR` field is the directory path within the repository that is where the source documentation resides.

  The `DIR` field is optional and defaults to the scripts internal mapping of project names to documentation source directories.
  For most projects, this is the `docs/sources` directory.

#### REPOS_PATH

The `REPOS_PATH` environment variable is a colon (`:`) separated list of paths in which to look for project repositories.
Only directories within the paths specified in `REPOS_PATH` are checked for projects.

By default, the script determines the `REPOS_PATH` to be the parent directory of the `grafana/technical-documentation` repository.
If you keep all repositories in the same directory, you do not need to set `REPOS_PATH`.

With a directory structure similar to the following `tree` command:

```console
$ tree -L 1 -d ~/ext/grafana
/home/jdb/ext/grafana
├── agent
├── grafana
├── loki
├── mimir
├── phlare
├── technical-documentation
├── tempo
├── website
└── writers-toolkit

9 directories
```

The script sets `REPOS_PATH` to be `/home/jdb/ext/grafana`.
