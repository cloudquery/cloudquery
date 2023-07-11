package apprunner

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

func resolveApprunnerTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		// AWS automatically generates untaggable resources with the following ARNs.
		//  - arn:aws:apprunner:<region>:01234567890:observabilityconfiguration/DefaultConfiguration/1/00000000000000000000000000000001
		//  - arn:aws:apprunner:<region>:01234567890:autoscalingconfiguration/DefaultConfiguration/1/00000000000000000000000000000001
		// So because they are untaggable we should just not even make the API call and end early
		if strings.Contains(aws.ToString(arn), "DefaultConfiguration/1/00000000000000000000000000000001") {
			return nil
		}
		cl := meta.(*client.Client)
		svc := cl.Services().Apprunner
		params := apprunner.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
