package resources

import "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"

type ELBv1LoadBalancerWrapper struct {
	types.LoadBalancerDescription
	Tags       map[string]interface{}
	Attributes *types.LoadBalancerAttributes
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}
