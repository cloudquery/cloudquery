package wellarchitected

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
		Name:        name,
		Description: `https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/wellarchitected/types#Lens`,
		Transform: transformers.TransformWithStruct(new(lens),
			transformers.WithPrimaryKeys("LensArn"),
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("lens_")),
		),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(name, "wellarchitected"),
		Resolver:            fetchLenses,
		PreResourceResolver: getLens,
		Columns:             schema.ColumnList{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
		Relations:           nil,
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

	// we do fetch for all 3 types
	out, err := service.GetLens(ctx,
		&wellarchitected.GetLensInput{LensAlias: l.LensAlias, LensVersion: l.LensVersion},
		func(o *wellarchitected.Options) { o.Region = cl.Region },
	)
	if err != nil {
		// at the very least we want the summary data to be filled in
		cl.Logger().Err(err).Str("table", "aws_wellarchitected_lenses").Msg("Failed to get lens")
	}

	// for err != nil basically
	if out != nil && out.Lens != nil {
		l.ShareInvitationId = out.Lens.ShareInvitationId
		l.Tags = out.Lens.Tags
	}

	resource.SetItem(l)
	return nil
}
