# Contributing to CloudQuery AWS Provider

👍🎉 First off, thanks for taking the time to contribute! 🎉👍

## Getting Started

## Architecture and SDK Overview

## Adding a resource

## Running Tests

To Run Integration tests please run

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d  postgres:13.3
az login
# If you are team member and have access it is advised to set subscription to our test environment
az account set --subscription 78f26f10-0e60-4293-8a7e-122584ccb40d
go test -run=TestIntegration -tags=integration ./...
```

To Run integration test for specific table:

```bash
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...

# For example
go test -run="TestIntegration/azure_sql_managed_instances" -tags=integration ./...
```

## Adding new Terraform File Guidelines

There a few good rule of thumb to follow when creating new terraform resources that will be served as testing infrastructure:
* For every resource create it's own resource_group.
* Use `location = "East US"`.
* If possible make all resources private.
* Make sure to replace built-in plain text passwords with `random_password` generator
* For every compute/db try to use the smallest size to keep the cost low
* If autoscaling option is present, always turn it off

If you want to apply the terraform locally first before pushing it to CI and applying there use:

```
cd terraform/local
terraform apply -var="name" -target="myspecific resource"
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
```