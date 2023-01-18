package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticbeanstalk
	output, err := svc.DescribeApplications(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Applications
	return nil
}

func resolveElasticbeanstalkApplicationTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ApplicationDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticbeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.ApplicationArn,
	}, func(o *elasticbeanstalk.Options) {})
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tagsOutput.ResourceTags))
}
