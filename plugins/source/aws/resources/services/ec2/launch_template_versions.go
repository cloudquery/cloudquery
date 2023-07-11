package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func launchTemplateVersions() *schema.Table {
	tableName := "aws_ec2_launch_template_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateVersion.html`,
		Resolver:    fetchEc2LaunchTemplateVersions,
		Transform:   transformers.TransformWithStruct(&types.LaunchTemplateVersion{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "version_number",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("VersionNumber"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchEc2LaunchTemplateVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := ec2.DescribeLaunchTemplateVersionsInput{
		LaunchTemplateId: parent.Item.(types.LaunchTemplate).LaunchTemplateId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	paginator := ec2.NewDescribeLaunchTemplateVersionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LaunchTemplateVersions
	}
	return nil
}
