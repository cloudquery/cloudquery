package appconfig

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func configurationProfiles() *schema.Table {
	tableName := "aws_appconfig_configuration_profiles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfigurationProfile.html`,
		Resolver:            fetchConfigurationProfiles,
		PreResourceResolver: getConfigurationProfiles,
		Transform:           transformers.TransformWithStruct(&appconfig.GetConfigurationProfileOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "application_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveConfigProfileARN,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			hostedConfigurationVersions(),
		},
	}
}

func fetchConfigurationProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Application)
	config := appconfig.ListConfigurationProfilesInput{
		ApplicationId: r.Id,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppconfig).Appconfig

	paginator := appconfig.NewListConfigurationProfilesPaginator(svc, &config)
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

func getConfigurationProfiles(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppconfig).Appconfig
	configurationProfileSummary := resource.Item.(types.ConfigurationProfileSummary)
	input := appconfig.GetConfigurationProfileInput{
		ApplicationId:          configurationProfileSummary.ApplicationId,
		ConfigurationProfileId: configurationProfileSummary.Id,
	}
	output, err := svc.GetConfigurationProfile(ctx, &input, func(o *appconfig.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

// ARN format defined here: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsappconfig.html#awsappconfig-resources-for-iam-policies
// arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/configurationprofile/${ConfigurationProfileId}
func resolveConfigProfileARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	configurationProfile := resource.Item.(*appconfig.GetConfigurationProfileOutput)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.AppconfigService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("application/%s/configurationprofile/%s", aws.ToString(configurationProfile.ApplicationId), aws.ToString(configurationProfile.Id)),
	}.String())
}
