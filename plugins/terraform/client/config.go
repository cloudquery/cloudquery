package client

type Config struct {
	Config []BackendConfigBlock `yaml:"config"`
}

func (Config) Example() string {
	return `
config:
  - name: mylocal # local backend
    backend: local
    path: ./examples/terraform.tfstate
  - name: myremote # s3 backend
    backend: s3
    bucket: tf-states
    key: terraform.tfstate
    region: us-east-1
    role_arn: ""
`
}
