package wellarchitected

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Workloads() *schema.Table {
	name := "aws_wellarchitected_workloads"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Workload.html`,
		Transform: transformers.TransformWithStruct(new(types.Workload),
			transformers.WithPrimaryKeys("WorkloadArn"),
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("workload_")),
		),
		Multiplex:           client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:            fetchWorkloads,
		PreResourceResolver: getWorkload,
		Columns:             schema.ColumnList{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
		Relations:           nil,
	}
}

func fetchWorkloads(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected

	p := wellarchitected.NewListWorkloadsPaginator(service, new(wellarchitected.ListWorkloadsInput))
	for p.HasMorePages() {
		output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.WorkloadSummaries
	}

	return nil
}

func getWorkload(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	summary := resource.Item.(types.WorkloadSummary)

	out, err := service.GetWorkload(ctx,
		&wellarchitected.GetWorkloadInput{WorkloadId: summary.WorkloadId},
		func(o *wellarchitected.Options) { o.Region = cl.Region },
	)

	if err != nil {
		// at the very least we want the summary data to be filled in
		cl.Logger().Err(err).Str("table", resource.Table.Name).Msg("Failed to get")
	}

	// for err != nil basically
	if out != nil && out.Workload != nil {
		resource.SetItem(out.Workload)
	}
	return nil
}
