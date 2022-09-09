package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
backends:
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

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"terraform",
		Version,
		[]*schema.Table{
			services.TFData(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
