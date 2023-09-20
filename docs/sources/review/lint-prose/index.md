---
description: How to lint prose with the Vale linter.
menuTitle: Lint prose
title: Lint prose with the Vale linter
weight: 300
aliases:
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/lint-prose/
  - /docs/writers-toolkit/review/lint-prose/
---

# Lint prose with the Vale linter

[Vale](https://github.com/errata-ai/vale) is a syntax-aware linter for prose built with speed and extensibility in mind.

Every project keeps technical documentation in the `docs/sources` directory.
Additionally, every project uses [GNU Make](https://www.gnu.org/software/make/) to perform tasks related to technical documentation.
To learn more about GNU Make, refer to [GNU Make Manual](https://www.gnu.org/software/make/manual/).

To see a list of targets and their descriptions, run `make` from the `docs/` directory.
The output is similar to the following:

```console
Usage:
  make <target>

Targets:
  help             Display this help.
  docs-rm          Remove the docs container.
  docs-pull        Pull documentation base image.
  make-docs        Fetch the latest make-docs script.
  docs             Serve documentation locally, which includes pulling the latest `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image. See also `docs-no-pull`.
  docs-no-pull     Serve documentation locally without pulling the `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image.
  docs-debug       Run Hugo web server with debugging enabled. TODO: support all SERVER_FLAGS defined in website Makefile.
  doc-validator    Run docs-validator on the entire docs folder.
  doc-validator/%  Run doc-validator on a specific path. To lint the path /docs/sources/administration, run 'make doc-validator/administration'.
  docs.mk          Fetch the latest version of this Makefile from Writers' Toolkit.
```

To lint prose with Vale, run `make vale` from the `docs/` directory.
Any linting errors are logged by the tool.

Additionally, some repositories run Vale as part of Continuous Integration (CI).
Repositories that run Vale in CI include:

- [Writers' Toolkit](https://github.com/grafana/writers-toolkit)

## Skip rules

To skip a rule, enclose the section with HTML comments that first disable, and then re-enable the specific Vale rule.

To disable the `Grafana.Quotes` rule:

```markdown
<!-- vale Grafana.Quotes = NO -->

The task title is required.
The task title succinctly describes the goal to be accomplished as the result of following the instruction.
The task title contains a verb and an object. For example: "Create a dashboard".

<!-- vale Grafana.Quotes = YES -->
```

## Use Vale in VS Code

You can use Vale to lint your current document in VS Code.

1. Clone the [Technical Documentation](https://github.com/grafana/technical-documentation) repository.

   ```bash
   git clone git@github.com:grafana/technical-documentation.git
   ```

1. Download and install [Vale](https://github.com/errata-ai/vale/releases).

   {{% admonition type="note" %}}
   Refer to the Linux installation steps at [GitHub Releases](https://vale.sh/docs/vale-cli/installation/#github-releases). Verify that you are downloading the latest build of Vale for Linux.
   {{% /admonition %}}

   {{< code >}}

   ```linux
   wget https://github.com/errata-ai/vale/releases/download/v2.28.0/vale_2.28.0_Linux_64-bit.tar.gz
   mkdir bin && tar -xvzf vale_2.28.0_Linux_64-bit.tar.gz -C bin
   export PATH=./bin:"$PATH"
   ```

   ```macos
   brew install vale
   ```

   {{< /code >}}

1. Create a `vale.ini` file in your home directory or in a working directory with the following contents:

   ```bash
   StylesPath = /FULL_PATH_TO_REPO/technical-documentation/linters/vale
   MinAlertLevel = suggestion

   [*.md]
   BasedOnStyles = Grafana
   ```

   Replace `FULL_PATH_TO_REPO` with the full path to the cloned Technical Documentation repository. For example, in Linux you could set StylesPath to `/home/username/git-repos/technical-documentation/linters/vale` and in macOS, you could set it to `/Users/username/git-repos/technical-documentation/linters/vale`. The path depends on where you cloned the git repository.

1. Install the [Vale VS Code extension](https://marketplace.visualstudio.com/items?itemName=chrischinchilla.vale-vscode) in VS Code.

   1. Start VS Code.
   1. Press Ctrl+P, paste the following command, and press Enter.

   ```
   ext install ChrisChinchilla.vale-vscode
   ```

1. Configure the Vale VS Code extension.

   1. Press Ctrl+Shift+X or click the **Extensions** icon and select the Vale VS Code extension.
   1. Select the gear icon.
   <!-- vale off -->
   1. Set **Vale › Vale CLI: Config** to the path to your `vale.ini` file. For example, on Linux that could be `/home/USERNAME/vale.ini` and on macOS, that could be `/Users/USERNAME/vale.ini`. The path depends on where you created the `vale.ini` file.
   <!-- vale on -->
   1. Set **Vale › Vale CLI: Path** to the path for the vale executable. For example, in Linux, that could be `/home/USERNAME/bin/vale` and on macOS, that could be `/usr/local/bin/vale`.

1. Restart VS Code.

Vale lints your current document every time you save your changes. The extension reports the linting results in two ways:

- In-line edit marks. You can hover your mouse cursor over the edit marks to view the vale warning or error.
1. A full report in the **PROBLEMS** tab.
