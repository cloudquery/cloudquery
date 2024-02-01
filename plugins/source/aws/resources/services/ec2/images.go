package ec2

import (
	"context"
	"reflect"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
	"golang.org/x/sync/errgroup"
)

func Images() *schema.Table {
	tableName := "aws_ec2_images"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html`,
		Resolver:    fetchEc2Images,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform: transformers.TransformWithStruct(&types.Image{},
			transformers.WithTypeTransformer(
				func(field reflect.StructField) (arrow.DataType, error) {
					switch field.Name {
					case "CreationDate", "DeprecationTime": // based on docs these are timestamps
						return arrow.FixedWidthTypes.Timestamp_us, nil
					default:
						return transformers.DefaultTypeTransformer(field)
					}
				},
			),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveImageArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			imageAttributesLaunchPermissions(),
			imageAttributesLastLaunchTime(),
		},
	}
}

func fetchEc2Images(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	svc := cl.Services(client.AWSServiceEc2).Ec2
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// fetch ec2.Images owned by this account
		pag := ec2.NewDescribeImagesPaginator(svc, &ec2.DescribeImagesInput{
			Owners:     []string{"self"},
			MaxResults: aws.Int32(1000),
		})
		for pag.HasMorePages() {
			resp, err := pag.NextPage(ctx, func(options *ec2.Options) {
				options.Region = cl.Region
			})
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
			MaxResults:      aws.Int32(1000),
		})
		for pag.HasMorePages() {
			resp, err := pag.NextPage(ctx, func(options *ec2.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for _, image := range resp.Images {
				if aws.ToString(image.OwnerId) != cl.AccountID {
					res <- image
				}
			}
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
		AccountID: aws.ToString(item.OwnerId),
		Resource:  "image/" + aws.ToString(item.ImageId),
	}
	return resource.Set(c.Name, a.String())
}
