---
title: Test documentation changes
description: Test documentation changes by running a local documentation webserver
weight: 200
aliases:
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/run-a-local-webserver
  - /docs/writers-toolkit/review/run-a-local-webserver
---

# Test documentation changes

{{< docs/shared source="writers-toolkit" lookup="make-help.md" >}}

To run the local documentation webserver with the default configuration, run `make docs` from the `docs/` directory.

{{< admonition type="note" >}}
Running `make docs` from the root of a repository produces the output `make: Nothing to be done for 'docs'.` instead of running the documentation webserver.
To run the webserver, ensure that you are in the `docs/` directory.
{{< /admonition >}}

## Run with specific projects

Each project has a list of projects to build by default when running `make docs` that is defined by the `PROJECTS` variable in `docs/variables.mk`.
To override the default for a single invocation, provide the `PROJECTS` argument to `make docs` which is the name of the project as it appears in the website URL path.

For example:

- For Grafana, the `PROJECT` is `grafana` derived the URL `https://grafana.com/docs/grafana/`.
- For Grafana Cloud, the `PROJECT` is `grafana-cloud` derived from the URL `https://grafana.com/docs/grafana-cloud/`.

{{% admonition type="note" %}}
You must have the repository cloned locally for any projects specified in the space separated list to `PROJECTS` for the command to succeed.
{{% /admonition %}}

To specifically build only the Grafana documentation:

```bash
make docs PROJECTS='grafana'
```

To specifically build Grafana and Grafana Cloud documentation:

```bash
make docs PROJECTS='grafana grafana-cloud'
```

Let's say that you have forked the main project repository, so your local working directory name doesn't match the project name.
You can use the `PROJECTS` option to define the local cloned repository (see the Arguments section below).

```bash
make docs PROJECTS="tempo::tempo-doc-work"
```

The format is `<PROJECT>[:VERSION[:REPOSITORY[:DIR]]].`
The example mounts the `PROJECT` tempo, at the default `VERSION` latest, using the `REPOSITORY` `tempo-doc-work`, and the default `DIR` `docs/sources`.
This example builds the Tempo documentation from the local working directory, `tempo-doc-work`, instead of the standard `tempo` directory.

## Understand Hugo output from `make docs`

When you run `make docs`, Hugo (our static site generator) processes the Markdown files and outputs warnings and error messages.

![Hugo output for running make docs](/media/docs/writers-toolkit/screenshot-make-docs-output.png)

These messages are in the following format:

```text
WARN <DATE> <TIME> <LANGUAGE> REF_NOT_FOUND: Ref <RELREF ARGUMENT>: “<SOURCE FILE>:<LINE>:<COLUMN>”: <ERROR>
```

where:

- `SOURCE FILE` is the file with the broken `relref`
- `RELREF ARGUMENT` is the argument to the `relref` shortcode that is not working.
- `ERROR` is the reason `RELREF ARGUMENT` is not working.

When you save a file with an active local build, the page is rechecked. If the error messages is not repeated, then the issue is fixed.

### Example: Page not found

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For more information about linking, refer to [Links](https://grafana.com/docs/writers-toolkit/write/links/).

### Example: Rebuild failed due to missing shortcode

In this example, the rebuild fails because the file `contribute-documentation/_index.md` is missing a closing shortcode for `admonition` on line 152.

```
ERROR Rebuild failed: assemble: "/hugo/content/docs/writers-toolkit/contribute-documentation/_index.md:152:1": failed to extract shortcode: shortcode "admonition" must be closed or self-closed
```

## Reference

The `make docs` target uses the [`make-docs`](https://github.com/grafana/writers-toolkit/blob/main/scripts/make-docs) script to mount local documentation into the Hugo build.
It can also be run separately if special configuration is required.

### Examples

The following examples assume:

1. Your working directory is the root of a project repository.
1. You have a copy of the `make-docs` script in that working directory.
1. The script is executable.
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

#### Arguments

Each argument to the `make-docs` script is a project to be mounted into the local build.
Each argument has four fields separated by colons (`:`) and optional fields can be omitted.

`<PROJECT>[:VERSION[:REPOSITORY[:DIR]]]`

- `PROJECT`: is the name of the project as it appears in the website URL path.

  For Grafana, the `PROJECT` is `grafana` (`https://grafana.com/docs/grafana/`)
  For Grafana Cloud, the `PROJECT` is `grafana-cloud` (`https://grafana.com/docs/grafana-cloud/`).

  Pseudo projects mount multiple source directories.

  - `logs`: mounts Loki and GEL directories.
  - `metrics`: mounts Mimir, the `mimir-distributed` Helm chart, and GEM directories.
  - `traces`: mounts Tempo and GET directories.

  **Note:** pseudo projects do not support the `REPOSITORY` or `DIR` fields.

- `VERSION`: is the name of the version directory under which the project documentation should be mounted.

  The `VERSION` field is optional and defaults to `latest` for versioned projects and is empty for other projects.

- `REPOSITORY` is the the name of the local directory within `REPOS_PATH` for that project.

  The `REPOSITORY` field is optional and defaults to the scripts internal mapping of project names to repository names.
  For most projects, this is the same as the project name.

- The `DIR` field is the directory path within the repository that is where the source documentation resides.

  The `DIR` field is optional and defaults to the scripts internal mapping of project names to documentation source directories.
  For most projects, this is the `docs/sources` directory.

This example builds the Grafana documentation and the Tempo documentation from the local repository, `tempo-doc-work`.

```bash
make docs PROJECTS="grafana tempo::tempo-doc-work"
```

#### REPOS_PATH

The `REPOS_PATH` environment variable is a colon-separated list of paths in which to look for project repositories.
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

#### DEBUG

The `DEBUG` environment variable disables output filtering and enables extra debug logging to help with troubleshooting.

If you experience confusing behavior with the `make docs` procedure, report the problem via a GitHub issue or, for Grafana Labs employees, in the #docs Slack channel and provide the full command and output using `make docs DEBUG=true`.

## Stop running a local build

To stop the `make docs` command, press Commmand/Ctrl + C.

If this doesn't work, do one of the following:

- (Recommended) Open Docker Desktop, go to **Containers**, and stop all running containers or just the one for your local build.
- To remove all running containers, run `docker rm -f $(docker ps -q)`.
