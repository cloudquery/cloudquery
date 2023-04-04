package amp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Workspaces() *schema.Table {
	tableName := "aws_amp_workspaces"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription`,
		Resolver:            fetchAmpWorkspaces,
		PreResourceResolver: describeWorkspace,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "amp"),
		Transform:           transformers.TransformWithStruct(&types.WorkspaceDescription{}),
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
			ruleGroupsNamespaces(),
		},
	}
}

func fetchAmpWorkspaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Amp

	p := amp.NewListWorkspacesPaginator(svc, &amp.ListWorkspacesInput{MaxResults: aws.Int32(int32(1000))})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.Workspaces
	}

	return nil
}

func describeWorkspace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Amp

	out, err := svc.DescribeWorkspace(ctx,
		&amp.DescribeWorkspaceInput{WorkspaceId: resource.Item.(types.WorkspaceSummary).WorkspaceId},
	)
	if err != nil {
		return err
	}

	resource.SetItem(out.Workspace)

	return nil
}

func describeAlertManagerDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Amp

	out, err := svc.DescribeAlertManagerDefinition(ctx,
		&amp.DescribeAlertManagerDefinitionInput{WorkspaceId: resource.Item.(*types.WorkspaceDescription).WorkspaceId},
	)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, out.AlertManagerDefinition)
}

func describeLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Amp

	out, err := svc.DescribeLoggingConfiguration(ctx,
		&amp.DescribeLoggingConfigurationInput{WorkspaceId: resource.Item.(*types.WorkspaceDescription).WorkspaceId},
	)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, out.LoggingConfiguration)
}
