Environment variables for running integration tests:

* `INTEGRATION_TESTS=1`
* `TF_APPLY_RESOURCES=1`
* `TF_VAR_PREFIX=test`
* `TF_VAR_SUFFIX=test`
* `KUBECONFIG=~/.kube/config`
* `KUBE_CONFIG_PATH=~/.kube/config`

Command to run them:
`go test -v -tags=integration ./...`


### Local Testing
- Start a minikube
- Start API server `kubectl proxy --port=80`
- Terraform apply resources