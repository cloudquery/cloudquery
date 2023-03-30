package amp

import (
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Workspaces() *schema.Table {
	tableName := "aws_amp_workspaces"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription`,
		Resolver:            fetchAmpWorkspaces,
		PreResourceResolver: describeWorkspace,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "amp"),
		Transform:           client.TransformWithStruct(&types.WorkspaceDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "alert_manager_definition",
				Type:     schema.TypeJSON,
				Resolver: describeAlertManagerDefinition,
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: describeLoggingConfiguration,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			RuleGroupsNamespaces(),
		},
	}
}
