package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeDBInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().Rds
	for {
		response, err := svc.DescribeDBInstances(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DBInstances
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveRdsInstanceProcessorFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	processorFeatures := map[string]*string{}
	for _, t := range r.ProcessorFeatures {
		processorFeatures[*t.Name] = t.Value
	}
	return resource.Set(c.Name, processorFeatures)
}

func resolveRdsInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
