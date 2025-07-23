---
aliases:
  - /docs/writers-toolkit/review/doc-validator/
  - /docs/writers-toolkit/review/doc-validator/errata/
  - /docs/writers-toolkit/review/lint-prose/
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/lint-prose/
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/validate-technical-documentation/
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/validate-technical-documentation/errata/
date: "2023-06-16T09:36:21+01:00"
description: How to lint prose for Grafana Labs style with the Vale linter.
menuTitle: Lint prose
review_date: "2024-05-28"
title: Lint prose with the Vale linter
weight: 300
---

# Lint prose with the Vale linter

[Vale](https://github.com/errata-ai/vale) is a syntax-aware linter for prose built with speed and extensibility in mind.

To lint prose with Vale locally, configure Vale in your text editor or run `make vale` from the `docs/` directory.
To use Vale in Visual Studio Code, refer to [Use Vale in Visual Studio Code](#use-vale-in-visual-studio-code).

Some repositories run Vale as part of Continuous Integration (CI).
Repositories that run Vale in CI include:

- [`grafana/beyla`](https://github.com/grafana/beyla/blob/main/.github/workflows/vale.yml)
- [`grafana/grafana`](https://github.com/grafana/grafana/blob/main/.github/workflows/documentation-ci.yml)
- [`grafana/website`](https://github.com/grafana/website/blob/master/.github/workflows/vale.yml)
- [`grafana/writers-toolkit`](https://github.com/grafana/writers-toolkit/blob/main/.github/workflows/validate-documentation.yml#L78-L88)

## Lint from the command line

Repositories with the `make docs` target also include a `make vale` target that lets you lint documentation from the command line.
It runs Vale against all the documentation in the `docs/sources/` directory.

To use the `make vale` target:

1. Change to the `docs/` directory.
1. Run `make vale`

## Skip CI

You may want to skip linting for pull requests that move files that have linting issues or otherwise defer linting the content.

To skip Vale linting in CI, include `<-- vale = NO -->` in your pull request description.

## Skip rules

To skip a rule, enclose the section with HTML comments that first disable, and then re-enable the specific Vale rule.
Include the specific rule name in the comment, for example `Grafana.We` or `Grafana.GoogleEllipses`.
The following example shows how to disable the `Grafana.GooglePassive` rule:

```markdown
<!-- vale Grafana.GooglePassive = NO -->

- [Deprecated content](#deprecation-example): Features that have been deprecated, but still need to be documented for some time.
- [Configuration options](#configuration-options-example): Features that have several ways they can be configured.

<!-- vale Grafana.GooglePassive = YES -->
```

## Use Vale in Visual Studio Code

You can use Vale to lint your current document in Visual Studio Code.

### Before you begin

If you are installing Vale on Linux, you may be able to install Vale from the package repositories for your Linux distribution.
However, the following manual installation instructions are the preferred way to install Vale on Linux.

{{< admonition type="note" >}}
The Vale Snap is out of date and not maintained.
Don't install Vale using the Ubuntu Snap store.
{{< /admonition >}}

If you are installing Vale on macOS, first install [Homebrew](https://brew.sh/).

If you are installing Vale on Windows, first install [Chocolatey](https://chocolatey.org/install).

{{< admonition type="note" >}}
You can [download and manually install](https://vale.sh/docs/vale-cli/installation/#github-releases) Vale on Linux, macOS, or Windows.

If you manually install Vale, you must configure your system to add Vale to your path or set the **Vale › Vale CLI: Path** in the Vale Visual Studio Code extension configuration.
{{< /admonition >}}

### Install and configure Vale in Visual Studio Code

1. Clone the [Writers' Toolkit](https://github.com/grafana/writers-toolkit/) repository.

   ```bash
   git clone git@github.com:grafana/writers-toolkit.git
   ```

   If you have previously cloned the repository, run `git pull` on the `main` branch.

1. Download and install [Vale](https://vale.sh/docs/vale-cli/installation/).

   {{< admonition type="note" >}}
   Verify that you are downloading the most recent build of Vale for Linux.
   To find the most recent build, refer to [Releases - errata-ai/vale](https://github.com/errata-ai/vale/releases).
   {{< /admonition >}}

   {{< code >}}

   ```linux-cli
   wget https://github.com/errata-ai/vale/releases/download/v3.1.0/vale_3.1.0_Linux_64-bit.tar.gz
   mkdir bin && tar -xvzf vale_3.1.0_Linux_64-bit.tar.gz -C bin
   export PATH=./bin:"$PATH"
   ```

   ```macos
   brew install vale
   ```

   ```windows
   choco install vale
   ```

   {{< /code >}}

1. Create a `.vale.ini` file in your home directory or in a working directory with the following contents:

   ```ini
   MinAlertLevel = suggestion
   Packages = <PATH TO WRITERS TOOLKIT REPOSITORY>/vale/Grafana, https://github.com/errata-ai/Hugo/releases/download/v0.2.0/Hugo.zip

   [*]
   BasedOnStyles = Grafana
   ```

   Replace _`<PATH TO WRITERS TOOLKIT REPOSITORY>`_ with the full path to your checkout of the Writer's Toolkit repository.
   The path depends on where you cloned the Git repository. For example:

   - On Linux, you could set StylesPath to `/home/<USERNAME>/git-repos/writers-toolkit/vale`
   - On macOS, you could set StylesPath to `/Users/<USERNAME>/git-repos/writers-toolkit/vale`
   - On Windows, you could set StylesPath to `C:\Users\<USERNAME>\git-repos\writers-toolkit\vale`

1. Install the configured Vale packages.

   From your home directory, run `vale sync` to install the packages you defined in your `.vale.ini` configuration file.

   {{< admonition type="note" >}}
   Every time you pull changes in your Writers' Toolkit repository, you need to sync those changes with `vale sync`.
   {{< /admonition >}}

1. Install the [Vale Visual Studio Code extension](https://marketplace.visualstudio.com/items?itemName=chrischinchilla.vale-vscode) in Visual Studio Code.

   1. Start Visual Studio Code.
   1. Press Ctrl+P, paste the following command, and press Enter.
      Alternatively, click the **Extensions** icon, search for "Vale VSCode", open it, and click **Install**.

      ```
      ext install ChrisChinchilla.vale-vscode
      ```

1. Configure the Vale Visual Studio Code extension.

   1. Press Ctrl+Shift+X or click the **Extensions** icon and select the Vale Visual Studio Code extension.
   1. Select the gear icon.
   1. To use your own Vale configuration for all repositories, set **Vale › Vale CLI: Config** to the path to your `.vale.ini` file.
      The path depends on where you created the `.vale.ini` file. For example:
      - On Linux, that could be `/home/<USERNAME>/.vale.ini`
      - On macOS, that could be `/Users/<USERNAME>/.vale.ini`
      - On Windows, that could be `C:\Users\<USERNAME>\.vale.ini`
   1. For manual installations on Linux, macOS, or Windows, set **Vale › Vale CLI: Path** to the path for the Vale executable.
      The path depends on where you extracted the Vale executable.
      For example, that could be `/home/<USERNAME>/bin/vale` on Linux.

1. Restart Visual Studio Code.

Vale lints your current document every time you save your changes.
The extension reports the linting results in two ways:

- In-line edit marks.
  You can hover your mouse cursor over the edit marks to view the Vale warning or error.
- A full report in the **PROBLEMS** tab.
  Each Vale warning or error in the report includes the line and column where the error occurs.
