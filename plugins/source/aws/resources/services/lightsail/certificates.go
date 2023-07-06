package lightsail

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Certificates() *schema.Table {
	tableName := "aws_lightsail_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html`,
		Resolver:    fetchLightsailCertificates,
		Transform:   transformers.TransformWithStruct(&types.Certificate{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchLightsailCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := lightsail.GetCertificatesInput{
		IncludeCertificateDetails: true,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	response, err := svc.GetCertificates(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	for _, cer := range response.Certificates {
		res <- cer.CertificateDetail
	}
	return nil
}
