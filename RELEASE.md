# Releases

The Writers' Toolkit repository is a collection of different tools that can have different release mechanisms.

## GitHub Actions

The following directories contain GitHub Actions actions:

- [`add-to-docs-project`](./add-to-docs-project/)
- [`prettier`](./prettier/)
- [`publish-technical-documentation`](./publish-technical-documentation/)
- [`publish-technical-documentation-release`](./publish-technical-documentation-release/)
- [`update-make-docs`](./update-make-docs/)

You release each action by creating or updating Git tags.
The Git tag begins with the action directory, then a slash (`/`), and then the tag version.
For example, `publish-technical-documentation-release/v1.0.0` is the v1.0.0 release of the `publish-technical-documentation-release` action.

The actions follow [semantic versioning](https://semver.org/).

To create a tag, use the following command:

```console
git tag --annotate --force --sign --local-user=<GPG KEY ID> -m <MESSAGE> <TAG>
```

Where:

- _`<GPG KEY ID>`_ is the ID of the GPG key associated with your GitHub account.

  For more information refer to [Signing commits](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits).

- _`<MESSAGE>`_ is a short message explaining the change to the action.

  This is typically similar to a commit title explaining the "what" of the change.

- _`TAG`_ is the name of the tag using the previously explained naming convention.

Each release has three tags:

- Major version: `<ACTION>/v<MAJOR>`
- Major and minor version: `<ACTION>/v<MAJOR>.<MINOR>`
- Major, minor, and patch version: `<ACTION>/v<MAJOR>.<MINOR>.<PATCH>`

The major version tag should point to the same commit as the latest major and minor version tag.
The major and minor version tag should point to the same commit as the major, minor, and patch version tag.

The `--force` flag provided to the `git tag` command overwrites existing tags which you need to be able to move the major version and major and minor version tags when you release a new major, minor, and patch version.

## Vale package

To release a new version of the Writers' Toolkit Vale package:

1. Create a new tag for the release with the form `vale/<VERSION>`.

   The Vale package follows [semantic versioning](https://semver.org).

1. Using the checked out tag, create an archive of the [Vale package directory](./vale/Grafana).

   ```console
   cd vale && make Grafana.zip
   ```

1. Create a GitHub release and upload the [`vale/Grafana.zip`](./vale/Grafana.zip) archive.
