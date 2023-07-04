package securityhub

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_securityhub_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html
` + "The `request_account_id` and `request_region` columns are added to show the account and region of where the request was made from." + `
This is useful when multi region and account aggregation is enabled.`,
		Resolver: fetchFindings,
		Transform: transformers.TransformWithStruct(&types.AwsSecurityFinding{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithPrimaryKeys("AwsAccountId", "Region", "CreatedAt", "UpdatedAt", "Description", "GeneratorId", "Id", "ProductArn", "SchemaVersion", "Title"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
		},
	}
}

func fetchFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	allConfigs := []tableoptions.CustomGetFindingsOpts{{}}
	noTableConfig := true
	if cl.Spec.TableOptions.SecurityHubFindings.GetFindingsOpts != nil {
		allConfigs = cl.Spec.TableOptions.SecurityHubFindings.GetFindingsOpts
		noTableConfig = false
	}

	svc := cl.Services().Securityhub
	var config securityhub.GetFindingsInput
	config.MaxResults = 100

	getFindings := func(svc services.SecurityhubClient, config securityhub.GetFindingsInput) error {
		p := securityhub.NewGetFindingsPaginator(svc, &config)
		for p.HasMorePages() {
			response, err := p.NextPage(ctx, func(o *securityhub.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- response.Findings
		}
		return nil
	}

	if !noTableConfig {
		for _, w := range allConfigs {
			config = w.GetFindingsInput
			if config.MaxResults == 0 {
				config.MaxResults = 100
			}
			err := getFindings(svc, config)
			if err != nil {
				return err
			}
		}
	} else {
		getFindings(svc, config)
	}
	return nil
}
