package guardduty

import (
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DetectorMembers() *schema.Table {
	return &schema.Table{
		Name:        "aws_guardduty_detector_members",
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html`,
		Resolver:    fetchGuarddutyDetectorMembers,
		Transform:   transformers.TransformWithStruct(&types.Member{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("guardduty"),
		Columns: []schema.Column{
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "detector_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
