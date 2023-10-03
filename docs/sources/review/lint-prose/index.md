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

### Before you begin

If you are installing Vale on Linux, it may be possible to use the package repositories for your Linux distribution. The manual installation documented below is the preferred way to install Vale on Linux.

{{% admonition type="note" %}}
The Vale Snap is out of date and not maintained. Don't install Vale using the Ubuntu Snap store.
{{% /admonition %}}

If you are installing Vale on macOS, make sure that [Homebrew](https://brew.sh/) is installed.

If you are installing Vale on Windows, make sure that [Chocolatey](https://chocolatey.org/install) is installed.

{{% admonition type="note" %}}
You can [download and manually install](https://vale.sh/docs/vale-cli/installation/#github-releases) Vale on Linux, macOS, or Windows. If you manually install Vale, you must configure your system to add Vale to your path or set the **Vale › Vale CLI: Path** in the Vale VS Code extension configuration.
{{% /admonition %}}

### Install and configure Vale in VS Code

1. Clone the [Writer's Toolkit](https://github.com/grafana/writers-toolkit/) repository.

   ```bash
   git clone git@github.com:grafana/writers-toolkit.git
   ```

1. Download and install [Vale](https://vale.sh/docs/vale-cli/installation/).

   {{% admonition type="note" %}}
   Verify that you are downloading the latest build of Vale for Linux.
   To find the latest builds, refer to [Releases - errata-ai/vale](https://github.com/errata-ai/vale/releases).
   {{% /admonition %}}

   {{< code >}}

   ```linux-cli
   wget https://github.com/errata-ai/vale/releases/download/v2.28.0/vale_2.29.0_Linux_64-bit.tar.gz
   mkdir bin && tar -xvzf vale_2.29.0_Linux_64-bit.tar.gz -C bin
   export PATH=./bin:"$PATH"
   ```

   ```macos
   brew install vale
   ```

   ```windows
   choco install vale
   ```

   {{< /code >}}

1. Create a `vale.ini` file in your home directory or in a working directory with the following contents:

   ```bash
   MinAlertLevel = suggestion
   StylesPath = /FULL_PATH_TO_REPO/writers-toolkit/vale
   [*.md]
   BasedOnStyles = Google, Grafana
   Google.Quotes = NO
   Google.Units = NO
   Google.WordList = NO
   TokenIgnores = (<http[^\n]+>+?), \*\*[^\n]+\*\*
   ```

   Replace `FULL_PATH_TO_REPO` with the full path to the cloned Writer's Toolkit repository. The path depends on where you cloned the git repository. For example:

   - On Linux, you could set StylesPath to `/home/<USERNAME>/git-repos/writers-toolkit/vale`
   - On macOS, you could set StylesPath to `/Users/<USERNAME>/git-repos/writers-toolkit/vale`
   - On Windows, you could set StylesPath to `C:\Users\<USERNAME>\git-repos\writers-toolkit\vale`

1. Install the [Vale VS Code extension](https://marketplace.visualstudio.com/items?itemName=chrischinchilla.vale-vscode) in VS Code.

   1. Start VS Code.
   1. Press Ctrl+P, paste the following command, and press Enter. Alternatively, click the **Extensions** icon, search for "Vale VS Code", open it, and click "Install".

   ```
   ext install ChrisChinchilla.vale-vscode
   ```

1. Configure the Vale VS Code extension.

   1. Press Ctrl+Shift+X or click the **Extensions** icon and select the Vale VS Code extension.
   1. Select the gear icon.
   1. Set **Vale › Vale CLI: Config** to the path to your `vale.ini` file. The path depends on where you created the `vale.ini` file. For example:
      - On Linux, that could be `/home/<USERNAME>/vale.ini`
      - On macOS, that could be `/Users/<USERNAME>/vale.ini`
      - On Windows, that could be `C:\Users\<USERNAME>\vale.ini`
   1. For manual installations on Linux, macOS, or Windows, set **Vale › Vale CLI: Path** to the path for the Vale executable. The path depends on where you unzipped the Vale executable. For example, that could be `/home/<USERNAME>/bin/vale` on Linux.

1. Restart VS Code.

Vale lints your current document every time you save your changes. The extension reports the linting results in two ways:

- In-line edit marks. You can hover your mouse cursor over the edit marks to view the Vale warning or error.
- A full report in the **PROBLEMS** tab. Each Vale warning or error in the report includes the line and column where the error occurs.

## Errata for Vale

When you write something that has an associated rule in one of the Vale linting files, an error is generated, such as:

`Use '%s' instead of '%s'.`

`Did you mean '%s' instead of '%s'?`

Most of these error messages and suggestions are self-explanatory and include preferred spellings or alternate words. However, the following rules require further explanation:

### Allows to

Common wording error. The linter suggests replacing "allows to" to with the grammatically correct "allows you to", since there is no use case for the phrase "allows to."
