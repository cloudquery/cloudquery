package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/hashstructure/v2"
)

func fetchRdsEngineVersions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds

	input := &rds.DescribeDBEngineVersionsInput{}

	p := rds.NewDescribeDBEngineVersionsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBEngineVersions
	}
	return nil
}
func calculateUniqueHash(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	hash, err := hashstructure.Hash(resource.Item, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, fmt.Sprint(hash))
}
