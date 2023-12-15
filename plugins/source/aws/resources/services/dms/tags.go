package dms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/thoas/go-funk"
)

// getTags fetches tags for the given resources and returns them in a map. arnPtrPath is the path to the ARN field in the resource. The returned map is keyed by ARN.
func getTags[T any](ctx context.Context, svc services.DatabasemigrationserviceClient, resources []T, arnPtrPath string, listTagsOpts ...func(*databasemigrationservice.Options)) (map[string]map[string]any, error) {
	listTagsForResourceInput := databasemigrationservice.ListTagsForResourceInput{}

	for _, item := range resources {
		listTagsForResourceInput.ResourceArnList = append(listTagsForResourceInput.ResourceArnList, *funk.Get(item, arnPtrPath).(*string))
	}
	listTagsForResourceOutput, err := svc.ListTagsForResource(ctx, &listTagsForResourceInput, listTagsOpts...)
	if err != nil {
		return nil, err
	}

	tags := make(map[string]map[string]any, len(listTagsForResourceOutput.TagList))
	for _, tag := range listTagsForResourceOutput.TagList {
		if tags[*tag.ResourceArn] == nil {
			tags[*tag.ResourceArn] = make(map[string]any)
		}
		tags[*tag.ResourceArn][*tag.Key] = *tag.Value
	}

	return tags, nil
}

// putTags puts tags in the given resources, into the "Tags" field. arnPtrPath is the path to the ARN field in the resource. The tags are fetched from the given map, which is keyed by ARN.
func putTags[T any](resources []T, tags map[string]map[string]any, arnPtrPath string) error {
	for i, item := range resources {
		arn := *funk.Get(item, arnPtrPath).(*string)
		if tags[arn] == nil {
			continue
		}
		if err := funk.Set(&resources[i], tags[arn], "Tags"); err != nil {
			return err
		}
	}
	return nil
}
