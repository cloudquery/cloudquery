package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VersionNumber"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchEc2LaunchTemplateVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := ec2.DescribeLaunchTemplateVersionsInput{
		LaunchTemplateId: parent.Item.(types.LaunchTemplate).LaunchTemplateId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	paginator := ec2.NewDescribeLaunchTemplateVersionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.LaunchTemplateVersions
	}
	return nil
}
