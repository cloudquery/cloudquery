package wellarchitected

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

// lens info consists of types.LensSummary & types.Lens fields
type lens struct {
	*types.LensSummary
	// extra field from types.Lens
	ShareInvitationId *string
	// extra field from types.Lens
	Tags map[string]string
}

func Lenses() *schema.Table {
	name := "aws_wellarchitected_lenses"
	return &schema.Table{
		Name:                name,
		Description:         `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Lens.html`,
		Transform:           transformers.TransformWithStruct(new(lens), transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:            fetchLenses,
		PreResourceResolver: getLens,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("LensArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchLenses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected

	// we do fetch for all 3 types
	for _, lensType := range types.LensType("").Values() {
		p := wellarchitected.NewListLensesPaginator(service,
			&wellarchitected.ListLensesInput{
				LensStatus: types.LensStatusTypeAll,
				LensType:   lensType,
				MaxResults: 50,
			},
		)
		for p.HasMorePages() {
			output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- output.LensSummaries
		}
	}

	return nil
}

func getLens(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	summary := resource.Item.(types.LensSummary)
	l := &lens{LensSummary: &summary}

	input := &wellarchitected.GetLensInput{LensAlias: l.LensAlias, LensVersion: summary.LensVersion}
	if summary.LensType == types.LensTypeAwsOfficial {
		input.LensVersion = nil // official lenses don't support versions
	}
	out, err := service.GetLens(ctx, input, func(o *wellarchitected.Options) { o.Region = cl.Region })
	if err != nil {
		cl.Logger().Err(err).Str("table", resource.Table.Name).Msg("Failed to perform get, ignoring...")
		// At the very least we want the summary data to be filled in
		// so don't update the item (leave summary in place)
		return nil
	}

	l.ShareInvitationId = out.Lens.ShareInvitationId
	l.Tags = out.Lens.Tags
	resource.SetItem(l)
	return nil
}
