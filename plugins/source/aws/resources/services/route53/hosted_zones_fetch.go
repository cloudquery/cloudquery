package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func getRoute53tagsByResourceID(id string, set []types.ResourceTagSet) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.Tags
		}
	}
	return nil
}
