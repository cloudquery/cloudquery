package ses

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func configurationSetEventDestinations() *schema.Table {
	tableName := "aws_ses_configuration_set_event_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html`,
		Resolver:    fetchSesConfigurationSetEventDestinations,
		Transform:   transformers.TransformWithStruct(&types.EventDestination{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "configuration_set_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchSesConfigurationSetEventDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2

	s := parent.Item.(*sesv2.GetConfigurationSetOutput)

	output, err := svc.GetConfigurationSetEventDestinations(ctx,
		&sesv2.GetConfigurationSetEventDestinationsInput{
			ConfigurationSetName: s.ConfigurationSetName,
		},
		func(o *sesv2.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	res <- output.EventDestinations

	return nil
}
