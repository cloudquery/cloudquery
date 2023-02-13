package guardduty

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Detectors() *schema.Table {
	return &schema.Table{
		Name:                "aws_guardduty_detectors",
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html`,
		Resolver:            fetchGuarddutyDetectors,
		PreResourceResolver: getDetector,
		Transform:           transformers.TransformWithStruct(&models.DetectorWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer("guardduty"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGuarddutyARN(),
			},
			{
				Name: "id",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			DetectorMembers(),
		},
	}
}
