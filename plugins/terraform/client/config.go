package client

type Config struct {
	Config []BackendConfigBlock `hcl:"config,block"`
}

func (c Config) Example() string {
	return `configuration {

	// local backend
	config "mylocal" {
      backend = "local"
      path = "./examples/terraform.tfstate"
    }
	// s3 backend
    config "myremote" {
      backend = "s3"
      bucket = "tf-states"
      key    = "terraform.tfstate"
      region = "us-east-1"
      role_arn = ""
    }
}
`
}
