# CQ Azure integration tests - cloud stack

## Prerequisites
* Terraform
* Azure Cli > 2.0.7

## Initialization
Export the current s3 bucket for tfstate file
```shell
export TF_BACKEND_BUCKET=cq-integration-tests-tf
export TF_BACKEND_KEY=azure/terraform.tfstate
```

Init terraform
```shell
make init
```

## Apply resources 
Create / Update Azure resources

Azure Terraform needs the following environment variables
```
export ARM_CLIENT_SECRET=
export ARM_CLIENT_ID=
export ARM_TENANT_ID=
export ARM_SUBSCRIPTION_ID=
```

```shell
make apply
```

## Destroy
Destroy all Azure resources
```shell
make destroy
```

## Clean
Clean tf files
```shell
make clean
```