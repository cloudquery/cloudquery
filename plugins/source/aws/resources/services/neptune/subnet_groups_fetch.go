package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchNeptuneSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config neptune.DescribeDBSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Neptune
	for {
		response, err := svc.DescribeDBSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DBSubnetGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
