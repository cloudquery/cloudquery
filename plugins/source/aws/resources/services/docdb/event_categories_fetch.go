package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbEventCategories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeEventCategoriesInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
	}

	response, err := svc.DescribeEventCategories(ctx, input)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.EventCategoriesMapList
	return nil
}
