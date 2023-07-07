package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EnrollmentStatuses() *schema.Table {
	tableName := "aws_computeoptimizer_enrollment_statuses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_GetEnrollmentStatus.html`,
		Resolver:    fetchEnrollmentStatus,
		Multiplex:   client.AccountMultiplex(tableName),
		Transform:   transformers.TransformWithStruct(&computeoptimizer.GetEnrollmentStatusOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchEnrollmentStatus(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Computeoptimizer

	output, err := svc.GetEnrollmentStatus(ctx, &computeoptimizer.GetEnrollmentStatusInput{}, func(options *computeoptimizer.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- output

	return nil
}
