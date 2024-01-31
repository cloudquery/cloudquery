package appconfig

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func hostedConfigurationVersions() *schema.Table {
	tableName := "aws_appconfig_hosted_configuration_versions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetHostedConfigurationVersion.html`,
		Resolver:            fetchHostedConfigurationVersions,
		PreResourceResolver: getHostedConfiguration,
		Transform:           transformers.TransformWithStruct(&appconfig.GetHostedConfigurationVersionOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "application_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveHostedConfigurationVersionARN,
				PrimaryKeyComponent: true,
			},
		},
		Relations: []*schema.Table{},
	}
}

func fetchHostedConfigurationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*appconfig.GetConfigurationProfileOutput)
	config := appconfig.ListHostedConfigurationVersionsInput{
		ApplicationId:          r.ApplicationId,
		ConfigurationProfileId: r.Id,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppconfig).Appconfig

	paginator := appconfig.NewListHostedConfigurationVersionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx, func(options *appconfig.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return err
		}
		res <- resp.Items
	}
	return nil
}

func getHostedConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppconfig).Appconfig
	hostedConfigurationVersionSummary := resource.Item.(types.HostedConfigurationVersionSummary)
	input := appconfig.GetHostedConfigurationVersionInput{
		ApplicationId:          hostedConfigurationVersionSummary.ApplicationId,
		ConfigurationProfileId: hostedConfigurationVersionSummary.ConfigurationProfileId,
		VersionNumber:          aws.Int32(hostedConfigurationVersionSummary.VersionNumber),
	}
	output, err := svc.GetHostedConfigurationVersion(ctx, &input, func(o *appconfig.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

// ARN format defined here: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsappconfig.html#awsappconfig-resources-for-iam-policies
// arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/configurationprofile/${ConfigurationProfileId}/hostedconfigurationversion/${VersionNumber}
func resolveHostedConfigurationVersionARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	hostedConfigurationVersion := resource.Item.(*appconfig.GetHostedConfigurationVersionOutput)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.AppconfigService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("application/%s/configurationprofile/%s/hostedconfigurationversion/%d", aws.ToString(hostedConfigurationVersion.ApplicationId), aws.ToString(hostedConfigurationVersion.ConfigurationProfileId), hostedConfigurationVersion.VersionNumber),
	}.String())
}
