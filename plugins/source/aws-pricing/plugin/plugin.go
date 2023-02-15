package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws-pricing/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws-pricing/resources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cloudquery-aws-pricing",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
