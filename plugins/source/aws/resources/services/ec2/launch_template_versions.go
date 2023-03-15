package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LaunchTemplateVersions() *schema.Table {
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
