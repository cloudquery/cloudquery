# Terraform Test Environment

This folder contains all terraform files to create a full test environment

## Adding and testing new resource with Terraform

Most of the time you don't want to create the whole infrastructure but rather the specific resource you added

```
az login
# Run the following if you have multiple subscription
az account set --subscription XXXX
terraform plan -resource
# Or run all
```

## Opening a PR

The PR for the terraform files shouold go seperatly from the resource implementation.

This way we the github action will be able to plan & apply the resource in CloudQuery staging environment.
After the resource is available in CloudQuery staging environment your can run the integration tests action in your resource PR where it should pass.