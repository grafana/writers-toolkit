# Technical documentation

This repository publishes technical documentation to the https://grafana.com website.
The [`publish-technical-documentation-next.yml`](../.github/workflows/publish-technical-documentation-next.yml) commits the contents of the [`/docs/sources`](./sources) directory to the `/content/docs/cloud-onboarding` directory in the [`grafana/website`](https://github.com/grafana/website) repository.

Pushes to the `main` branch that add, delete, or modify files in the `/docs/sources` directory trigger the workflow.
To manually run the workflow, refer to [Manually run the workflow](https://github.com/grafana/technical-documentation/blob/main/docs/sources/publish-technical-documentation/index.md#manually-run-the-workflow).

If the workflow fails unexpectedly, reach out to the `#docs` channel on Slack with a link to the source commit or pull request, and the workflow run.

Successful completion of the workflow means that the contents have been committed to the [`grafana/website`](https://github.com/grafana/website) repository.
There are subsequent build pipelines that perform the production build and deploy.
For insight into the full procedure, refer to the [Website dashboard](https://ops.grafana-ops.net/d/c88932b3-7302-4695-b5eb-7b3c9d501b94/website?orgId=1&refresh=1m).

## Local development

A local webserver serves from http://localhost:3002/docs/grafana-cloud/data-configuration/integrations/.
For more information about the webserver, refer to [Run a local documentation webserver | Writers' Toolkit documentation](https://grafana.com/docs/writers-toolkit/writing-guide/tooling-and-workflows/run-a-local-webserver/).

The supporting scripts and Makefiles are maintained in the [`grafana/writers-toolkit`](https://github.com/grafana/writers-toolkit) repository.
To update those files, run `make update`.
