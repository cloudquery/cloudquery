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

func studioSessionMapping() *schema.Table {
	tableName := "aws_emr_studio_session_mappings"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_GetStudioSessionMapping.html`,
		Resolver:            fetchEmrStudioSessionMapping,
		PreResourceResolver: getSessionMapping,
		Transform:           transformers.TransformWithStruct(&types.SessionMappingDetail{}, transformers.WithPrimaryKeyComponents("IdentityType", "IdentityId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "studio_arn",
				Description:         "The Amazon Resource Name (ARN) of the EMR Studio.",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchEmrStudioSessionMapping(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*types.Studio)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListStudioSessionMappingsPaginator(svc, &emr.ListStudioSessionMappingsInput{
		StudioId: p.StudioId,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SessionMappings
	}
	return nil
}

func getSessionMapping(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	sms := resource.Item.(types.SessionMappingSummary)
	response, err := svc.GetStudioSessionMapping(ctx, &emr.GetStudioSessionMappingInput{
		StudioId:     sms.StudioId,
		IdentityType: sms.IdentityType,
		IdentityId:   sms.IdentityId,
	}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.SessionMapping
	return nil
}
