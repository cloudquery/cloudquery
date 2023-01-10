package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

func fetchEc2Images(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().Ec2
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// fetch ec2.Images owned by this account
		pag := ec2.NewDescribeImagesPaginator(svc, &ec2.DescribeImagesInput{
			Owners: []string{"self"},
		})
		for pag.HasMorePages() {
			resp, err := pag.NextPage(ctx)
			if err != nil {
				return err
			}
			res <- resp.Images
		}
		return nil
	})

	g.Go(func() error {
		// fetch ec2.Images that are shared with this account
		pag := ec2.NewDescribeImagesPaginator(svc, &ec2.DescribeImagesInput{
			ExecutableUsers: []string{"self"},
		})
		for pag.HasMorePages() {
			resp, err := pag.NextPage(ctx)
			if err != nil {
				return err
			}
			res <- resp.Images
		}
		return nil
	})

	return g.Wait()
}

func resolveImageArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Image)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "image/" + aws.ToString(item.ImageId),
	}
	return resource.Set(c.Name, a.String())
}
