package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudhsmv2Backups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudhsmv2
	var input cloudhsmv2.DescribeBackupsInput
	paginator := cloudhsmv2.NewDescribeBackupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Backups
	}
	return nil
}

func resolveBackupArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Backup)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "hsm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "backup/" + aws.ToString(item.BackupId),
	}
	return resource.Set(c.Name, a.String())
}
