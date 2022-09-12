package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

var glueResources []*Resource = []*Resource{
	{
		Struct: &types.Crawler{},
	},
}

func GlueResources() []*Resource {
	return glueResources
}
