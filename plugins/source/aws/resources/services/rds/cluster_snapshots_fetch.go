package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds
	var input rds.DescribeDBClusterSnapshotsInput
	for {
		output, err := svc.DescribeDBClusterSnapshots(ctx, &input)
		if err != nil {
			return nil
		}
		res <- output.DBClusterSnapshots
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRDSClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	tags := map[string]*string{}
	for _, t := range s.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}

func resolveRDSClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().Rds
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = c.Region
		},
	)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if out.DBClusterSnapshotAttributesResult == nil {
		return nil
	}

	return resource.Set(column.Name, out.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes)
}
