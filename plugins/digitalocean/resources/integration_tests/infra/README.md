# CQ DigitalOcean integration tests - cloud stack

## Prerequisites
* Terraform

## Authentication
Export the DigitalOcean credentials
```shell
export DIGITALOCEAN_TOKEN=
OR
export DIGITALOCEAN_ACCESS_TOKEN=
```


## Initialization
Export the current s3 bucket for tfstate file
```shell
export TF_BACKEND_BUCKET=cq-integration-tests-tf
export TF_BACKEND_KEY=do/terraform.tfstate
```

Init terraform
```shell
make init
```

## Apply resources 
Create / Update DigitalOcean resources
```shell
make apply
```

## Destroy
Destroy all DigitalOcean resources
```shell
make destroy
```

## Clean
Clean tf files
```shell
make clean
```