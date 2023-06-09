# Overview

CloudQuery can run locally, but if you want to deploy in a remote non-ephemeral environment to sync periodically and store the data in a managed PostgreSQL, the recommended way is to deploy on Kubernetes (EKS or GKE) with our [helm-charts](https://github.com/cloudquery/helm-charts).

We also provide a [Terraform module](https://github.com/cloudquery/terraform-aws-cloudquery) to spin up the infrastructure: k8s clusters, database and deployment of the helm chart via the Terraform helm-provider to provide a better deployment experience. Follow the example in the [AWS Terraform module](https://github.com/cloudquery/terraform-aws-cloudquery) or the helm chart to see how to deploy and configure CloudQuery to run periodically in the cloud.

Apart from these options, we have also written some guides to help you get started on different platforms.

## Tutorials

 - [Airflow](/docs/deployment/airflow)
 - [Amazon ECS](/docs/deployment/ecs)
 - [Docker](/docs/deployment/docker)
 - [GitHub Actions](/docs/deployment/github-actions)
 - [Google Cloud Virtual Machine](/docs/deployment/google-cloud-vm)
 - [Kestra](/docs/deployment/kestra)
