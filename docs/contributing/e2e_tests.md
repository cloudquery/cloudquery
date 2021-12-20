# Running cq-provider-aws E2E Tests:


## Prerequisites:

1. Go
2. [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
3. [Terraform Prerequisites](https://learn.hashicorp.com/tutorials/terraform/aws-build)



## Running Tests

All of the tests in the `resources/integration_tests` directory use a wrapper to dynamically call terraform in order to provision resources at test time and then tear down the resources when the testing is complete. You do not have to deploy the entire infrastructure suite in order to run the tests, you can pick and choose exactly which tests you want to run

Using the Makefile in the root of the repository is the simplest way to start running your tests locally. 

(hint: You can include regular expressions in the test name to match multiple tests)

example of running multiple tests (`TestIntegrationCognitoUserPools` and `TestIntegrationCognitoIdentityPools`):
``` bash
make testName=^TestIntegrationCognito$ e2e-test-with-apply
```


### Running Read Only Tests


If you have another way of deploying the terraform defined in the `resources/integration_tests/infra` directory you can do that and then run the tests against the deployed resources like this:


``` bash
export TF_VAR_PREFIX=WhatEverValueYouSpecified
export TF_VAR_SUFFIX=WhatEverValueYouSpecified
make testName=^TestIntegrationCognito$ e2e-test-with-apply
```