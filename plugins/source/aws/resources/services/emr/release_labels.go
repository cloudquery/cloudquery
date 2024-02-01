package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ReleaseLabels() *schema.Table {
	tableName := "aws_emr_release_labels"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeReleaseLabel.html`,
		Resolver:            fetchEmrReleaseLabels,
		PreResourceResolver: getReleaseLabel,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform: transformers.TransformWithStruct(
			&emr.DescribeReleaseLabelOutput{},
			transformers.WithPrimaryKeyComponents("ReleaseLabel"),
			transformers.WithSkipFields("ResultMetadata", "NextToken"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
		Relations: []*schema.Table{supportedInstanceTypes()},
	}
}

func fetchEmrReleaseLabels(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListReleaseLabelsPaginator(svc, &emr.ListReleaseLabelsInput{})
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ReleaseLabels
	}
	return nil
}

func getReleaseLabel(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	releaseLabel := resource.Item.(string)

	config := &emr.DescribeReleaseLabelInput{ReleaseLabel: &releaseLabel}
	result := &emr.DescribeReleaseLabelOutput{ReleaseLabel: &releaseLabel}

	// No paginator available
	for {
		response, err := svc.DescribeReleaseLabel(ctx, config, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		result.Applications = append(result.Applications, response.Applications...)
		result.AvailableOSReleases = append(result.AvailableOSReleases, response.AvailableOSReleases...)
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	resource.Item = result
	return nil
}
