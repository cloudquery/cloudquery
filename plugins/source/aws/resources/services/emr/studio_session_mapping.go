package emr

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func studioSessionMapping() *schema.Table {
	tableName := "aws_emr_studio_session_mapping"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_SessionMappingSummary.html`,
		Resolver:            fetchEmrStudioSessionMapping,
		PreResourceResolver: getSessionMapping,
		Transform:           transformers.TransformWithStruct(&types.SessionMappingDetail{}),
	}
}

func fetchEmrStudioSessionMapping(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*types.Studio)
	svc := cl.Services().Emr
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
	svc := cl.Services().Emr
	response, err := svc.GetStudioSessionMapping(ctx, &emr.GetStudioSessionMappingInput{
		StudioId:     resource.Item.(types.SessionMappingSummary).StudioId,
		IdentityType: resource.Item.(types.SessionMappingSummary).IdentityType,
		IdentityId:   resource.Item.(types.SessionMappingSummary).IdentityId,
	}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.SessionMapping
	return nil
}
