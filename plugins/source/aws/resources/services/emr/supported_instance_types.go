package emr

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func supportedInstanceTypes() *schema.Table {
	tableName := "aws_emr_supported_instance_types"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_SupportedInstanceType.html`,
		Resolver:    fetchEmrSupportedInstanceTypes,
		Transform:   transformers.TransformWithStruct(&types.SupportedInstanceType{}, transformers.WithPrimaryKeyComponents("Type")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "release_label",
				Description:         "The release label of the EMR cluster.",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("release_label"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchEmrSupportedInstanceTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	p := parent.Item.(*emr.DescribeReleaseLabelOutput)
	paginator := emr.NewListSupportedInstanceTypesPaginator(svc, &emr.ListSupportedInstanceTypesInput{ReleaseLabel: p.ReleaseLabel})
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, instanceType := range response.SupportedInstanceTypes {
			res <- instanceType
		}
	}
	return nil
}
