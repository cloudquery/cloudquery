package spaces

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func cors() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_space_cors",
		Resolver:  fetchSpacesCors,
		Transform: transformers.TransformWithStruct(&types.CORSRule{}),
		Columns:   []schema.Column{},
	}
}
