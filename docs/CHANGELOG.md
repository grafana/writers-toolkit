# `make-docs` script changelog

<!-- Updates should conform to the guidelines in https://keepachangelog.com/en/1.1.0/ -->

[Semantic versioning](https://semver.org/) is used to help the reader identify the significance of changes.

## Unreleased

## 4.0.0 (2023-06-06)

### Removed

- `doc-validator/%` target.
  The behavior of the target was not as described.
  Instead, to limit `doc-validator` to only specific files, refer to https://grafana.com/docs/writers-toolkit/writing-guide/tooling-and-workflows/validate-technical-documentation/#run-on-specific-files.

## 3.0.0 (2023-05-18)

### Fixed

- Compatibility with the updated Make targets in the `website` repository.
  `docs` now runs this script itself, `server-docs` builds the site with the `docs` Hugo environment.

## 2.0.0 (2023-05-18)

### Added

- Support for the grafana-cloud/frontend-observability/faro-web-sdk project.
- Use of `doc-validator` v2.0.x which includes breaking changes to command line options.

### Fixed

- Source grafana-cloud project from website repository.

### Added

- Support for running the Vale linter with `make vale`.

## 1.2.1 (2023-05-05)

### Fixed

- Use `latest` tag of `grafana/vale` image by default instead of hardcoded older version.
- Fix mounting multiple projects broken by the changes in 1.0.1

## 1.2.0 (2023-05-05)

### Added

- Support for running the Vale linter with `make vale`.

### Fixed

## 1.1.0 (2023-05-05)

### Added

- Rewrite error output so it can be followed by text editors.

### Fixed

- Fix `docs-debug` container process port.

## 1.0.1 (2023-05-04)

### Fixed

- Ensure complete section hierarchy so that all projects have a visible menu.

## 1.0.0 (2023-05-04)

### Added

- Build multiple projects simultaneously if all projects are checked out locally.
- Run [`doc-validator`](https://github.com/grafana/technical-documentation/tree/main/tools/cmd/doc-validator) over projects.
- Redirect project root to mounted version.
  For example redirect `/docs/grafana/` to `/docs/grafana/latest/`.
- Support for Podman or Docker containers with `PODMAN` environment variable.
- Support for projects:
  - agent
  - enterprise-logs
  - enterprise-metrics
  - enterprise-traces
  - grafana
  - grafana-cloud
  - grafana-cloud/machine-learning
  - helm-charts/mimir-distributed
  - helm-charts/tempo-distributed
  - incident
  - loki
  - mimir
  - oncall
  - opentelemetry
  - phlare
  - plugins
  - slo
  - tempo
  - writers-toolkit
