package websecurityscanner

import (
	"context"
	"fmt"

	websecurityscanner "cloud.google.com/go/websecurityscanner/apiv1"
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

func fetchFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	parentItem := parent.Item.(*pb.ScanRun)

	gcpClient, err := websecurityscanner.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	findingTypes, err := getFindingTypes(ctx, gcpClient, parentItem.Name, c.CallOptions...)
	if err != nil {
		return err
	}

	for _, findingType := range findingTypes {
		it := gcpClient.ListFindings(ctx, &pb.ListFindingsRequest{
			Parent: parentItem.Name,
			Filter: fmt.Sprintf(`"finding_type="%s"`, findingType),
		}, c.CallOptions...)
		for {
			resp, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			res <- resp
		}
	}

	return nil
}

func getFindingTypes(ctx context.Context, gcpClient *websecurityscanner.Client, parentId string, callOptions ...gax.CallOption) ([]string, error) {
	response, err := gcpClient.ListFindingTypeStats(ctx, &pb.ListFindingTypeStatsRequest{
		Parent: parentId,
	}, callOptions...)

	if err != nil {
		return nil, err
	}

	findingTypes := make([]string, len(response.FindingTypeStats))

	for _, typestat := range response.FindingTypeStats {
		findingTypes = append(findingTypes, typestat.FindingType)
	}

	return findingTypes, nil
}
