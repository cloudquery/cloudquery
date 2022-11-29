# Overview

CloudQuery can run locally but if you want to deploy in a remote non-ephemeral environment to fetch periodically and store the data in a managed PostgreSQL the current suggested way is to deploy on k8s (EKS or GKE) with our [helm-charts](https://github.com/cloudquery/helm-charts).

We also provide a Terraform module to spin-up the infrastructure - k8s clusters + database + deploy the helm chart via the Terraform helm-provider to provide a better deployment experience.

Follow the example in the [AWS Terraform module](https://github.com/cloudquery/terraform-aws-cloudquery) or the helm chart on how to deploy and configure CloudQuery to run periodically in the cloud.
