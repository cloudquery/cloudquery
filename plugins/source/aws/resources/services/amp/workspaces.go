package amp

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: describeAlertManagerDefinition,
			},
			{
				Name:     "logging_configuration",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: describeLoggingConfiguration,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			ruleGroupsNamespaces(),
		},
	}
}

func fetchAmpWorkspaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	p := amp.NewListWorkspacesPaginator(svc, &amp.ListWorkspacesInput{MaxResults: aws.Int32(int32(1000))})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *amp.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.Workspaces
	}

	return nil
}

func describeWorkspace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	out, err := svc.DescribeWorkspace(ctx,
		&amp.DescribeWorkspaceInput{WorkspaceId: resource.Item.(types.WorkspaceSummary).WorkspaceId},
		func(options *amp.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	resource.SetItem(out.Workspace)

	return nil
}

func describeAlertManagerDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	out, err := svc.DescribeAlertManagerDefinition(ctx,
		&amp.DescribeAlertManagerDefinitionInput{WorkspaceId: resource.Item.(*types.WorkspaceDescription).WorkspaceId},
		func(options *amp.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	return resource.Set(col.Name, out.AlertManagerDefinition)
}

func describeLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	out, err := svc.DescribeLoggingConfiguration(ctx,
		&amp.DescribeLoggingConfigurationInput{WorkspaceId: resource.Item.(*types.WorkspaceDescription).WorkspaceId},
		func(options *amp.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	return resource.Set(col.Name, out.LoggingConfiguration)
}
