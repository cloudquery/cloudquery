# CQ AWS integration tests - cloud stack

## Prerequisites
* Terraform

## Initialization
Export the current s3 bucket for tfstate file
```shell
export TF_BACKEND_BUCKET=cq-integration-tests-tf
export TF_BACKEND_KEY=terraform.tfstate
```

Init terraform
```shell
make init
```

## Apply resources 
Create / Update AWS resources
```shell
make apply
```

## Destroy
Destroy all AWS resources
```shell
make destroy
```

## Clean
Clean tf files
```shell
make clean
```