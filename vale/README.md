# Vale

The [`.vale.ini`](.vale.ini) file in this directory is only for the container image.
For linting of this repository, the [`.vale.ini`](../.vale.ini) file in the root of the repository is used.
The two files should be mostly similar with the exception of the `Packages` and `StylesPath` configurations.
Perhaps in the future, these will be sourced from a single place.

The Grafana style extends the Google style and disables some rules.
The Google style is vendored in the repository.
To update the Google style, run `make sync`.

To build the container image, run `make grafana/vale`.
