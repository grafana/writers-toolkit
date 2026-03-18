---
aliases:
  - /docs/writers-toolkit/review/run-a-local-webserver/
  - /docs/writers-toolkit/review/test-documentation-changes/
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/run-a-local-webserver/
date: "2024-02-23T15:12:32+00:00"
description: Test documentation changes by running a local documentation web server
review_date: "2024-05-28"
title: Test documentation changes
weight: 200
---

# Test documentation changes

{{< docs/shared source="writers-toolkit" lookup="make-help.md" >}}

To run the local documentation web server, run `make docs` from the `docs/` directory.
If you are in the website repository, run `make docs` from the root of the repository instead.

{{< admonition type="note" >}}
Running `make docs` from the wrong directory, produces the output `make: Nothing to be done for 'docs'.` or `make: *** No rule to make target 'docs'.  Stop.`, and the local documentation web server is not started.
To run the local documentation web server, ensure that you are in the right directory.
{{< /admonition >}}

The output message of a successful build includes a URL that you can follow to view the changes to the documentation in the browser.
Refer to an [example of a successful build](#example-successful-build).

## Run with specific projects

Each repository has a list of projects to build by default when running `make docs`, defined by the `PROJECTS` variable in `docs/variables.mk`.
To override the defaults, provide the `PROJECTS` option to `make docs`.
The argument is a space-separated list of names of the projects.

The project name for a repository is the sub-directory of the `/docs/` directory in the website where the repository publishes documentation.
For example:

- For Grafana, the `PROJECT` is `grafana` derived the URL `https://grafana.com/docs/grafana/`.
- For Grafana Cloud, the `PROJECT` is `grafana-cloud` derived from the URL `https://grafana.com/docs/grafana-cloud/`.

{{< admonition type="note" >}}
You must have the repository cloned locally for any projects specified in the space separated list to `PROJECTS` for the command to succeed.
{{< /admonition >}}

To build only the Grafana documentation:

```bash
make docs PROJECTS=grafana
```

To build Grafana and Grafana Cloud documentation:

```bash
make docs 'PROJECTS=grafana grafana-cloud'
```

If your local repository name doesn't match the upstream repository name.
You can use the `PROJECTS` option to override the directory.

For example, if you have the Tempo repository `tempo` cloned into a directory called `tempo-doc-work`.

```bash
make docs PROJECTS=tempo::tempo-doc-work
```

The format is `<PROJECT>[:VERSION[:REPOSITORY[:DIRECTORY]]].`
The example mounts the `PROJECT` `tempo`, at the default `VERSION` (since the `VERSION` argument is empty), using the `REPOSITORY` `tempo-doc-work`, and the default documentation `DIRECTORY` `docs/sources`.
This example builds the Tempo documentation from the local working directory, `tempo-doc-work`, instead of the standard `tempo` directory.

## Understand Hugo output from `make docs`

When you run `make docs`, Hugo renders the Markdown files and outputs warnings and error messages.

![Hugo output for running make docs](/media/docs/writers-toolkit/screenshot-make-docs-output.png)

These messages are in the following format:

```text
<LEVEL> [LANGUAGE] <MESSAGE>
```

where:

- _`<LEVEL>`_ is one of `WARN` or `ERROR`
- _`<LANGUAGE>`_ may be present
- _`<MESSAGE>`_ is the issue

### Example: Successful build

If the `make docs` command runs correctly, the console prints a message similar to the following:

```
View documentation locally:
  http://localhost:3002/docs/grafana/latest/

Press Ctrl+C to stop the server
```

{{< admonition type="note" >}}
To test Grafana tutorial changes, the address is `http://localhost:3002/docs/grafana/latest/tutorials/`.
{{< /admonition >}}

### Example: Page not found

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For more information about linking, refer to [Links](https://grafana.com/docs/writers-toolkit/write/links/).

### Example: Rebuild failed due to missing shortcode

In this example, the rebuild fails because the file `contribute-documentation/_index.md` is missing a closing shortcode for `admonition` on line 152.

```
ERROR Rebuild failed: assemble: "/hugo/content/docs/writers-toolkit/contribute-documentation/_index.md:152:1": failed to extract shortcode: shortcode "admonition" must be closed or self-closed
```

### Extended usage

Refer to the following sections for examples of more complicated usage of `make docs`.

#### Mount documentation to a different version

```
make docs PROJECTS=grafana:next
```

#### Mount `v9.3.x` and the default version of Grafana documentation together

Run the following command from the root of the repository to add a worktree that contains the `v9.3.x` branch:

```bash
git worktree add v9.3.x origin/v9.3.x
```

Change to the `docs/` directory:

```bash
cd docs
```

Run `make docs` with both versions:

```
make docs 'PROJECTS=grafana grafana:v9.3.x:grafana'
```

To remove the worktree, run the following command from the root of the repository:

```bash
git worktree remove v9.3.x
```

### Arguments

Each argument to the `make-docs` script is a project to be mounted into the local build.
Each argument has four fields separated by colons (`:`), and optional fields can be omitted.

`<PROJECT>[:VERSION[:REPOSITORY[:DIRECTORY]]]`

- _`<PROJECT>`_: is the sub-directory of the `/docs/` directory in the website where the repository publishes documentation.

  For example:

  - For Grafana, _`<PROJECT>`_ is `grafana` (`https://grafana.com/docs/grafana/`)
  - For Grafana Cloud, _`<PROJECT>`_ is `grafana-cloud` (`https://grafana.com/docs/grafana-cloud/`).

  Pseudo projects mount multiple source directories.

  - `logs`: mounts Loki and Grafana Enterprise Logs (GEL) directories.
  - `metrics`: mounts Mimir, the `mimir-distributed` Helm chart, and Grafana Enterprise Metrics (GEM) directories.
  - `traces`: mounts Tempo and Grafana Enterprise Traces (GET) directories.

  {{< admonition type="note" >}}
  Pseudo projects don't support the _`<REPOSITORY>`_ or _`<DIR>`_ fields.
  {{< /admonition >}}

- _`<VERSION>`_: is the name of the version directory to mount the documentation in.

  The _`<VERSION>`_ field is optional and defaults to `latest` for versioned projects and is empty for other projects.

- _`<REPOSITORY>`_: is the name of the directory that the project is cloned to.

  The _`<REPOSITORY>`_ field is optional and defaults to the script's internal mapping of project names to repository names.
  For most projects, this is the same as the project name.

- _`<DIRECTORY>`_: is the directory path within the repository containing the technical documentation documentation.

  The _`<DIRECTORY>`_ field is optional and defaults to the script's internal mapping of project names to documentation source directories.
  For most projects, it's the `docs/sources` directory.

#### `REPOS_PATH`

The `REPOS_PATH` environment variable is a colon-separated list of paths in which to look for project repositories.
The script only checks for projects in the directories specified in `REPOS_PATH`.

By default, the script determines the `REPOS_PATH` to be the parent directory of the current project.
If you keep all repositories in the same directory, you don't need to set `REPOS_PATH`.

With a directory structure similar to the following output from the `tree -L 1 -d ~/ext/grafana` command:

```console
/home/username/ext/grafana
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

When you run `make docs` from the Grafana repository, the script sets `REPOS_PATH` to be `/home/username/ext/grafana`.

#### `DEBUG`

The `DEBUG` environment variable disables output filtering and enables extra debug logging to help with troubleshooting.

If you experience confusing behavior with the `make docs` procedure, report the problem via a GitHub issue or, for Grafana Labs employees, in the #docs Slack channel and provide the full command and output using `make docs DEBUG=true`.

## Stop running a local build

To stop the `make docs` command, press Commmand/Ctrl + C.

Sometimes an old build process can be running in another terminal.
If this is the case, when you run `make docs`, you see output similar to the following:

```console
docker: Error response from daemon: driver failed programming external connectivity on endpoint eloquent_nightingale (eb2c4546727b41bbc44354ac616a14404c57f30c312f6869b147c578ac5de6bf): Bind for 0.0.0.0:3002 failed: port is already allocated.
make: *** [docs] Error 125
```

To remove an old build process, do one of the following:

- Open Docker Desktop, go to **Containers**, and stop all running containers or just the one for your local build.
- To remove all running containers, run `docker rm -f $(docker ps -q)`.
