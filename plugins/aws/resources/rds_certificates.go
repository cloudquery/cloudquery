package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func RdsCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_certificates",
		Resolver:     fetchRdsCertificates,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "certificate_arn",
				Type: schema.TypeString,
			},
			{
				Name: "certificate_identifier",
				Type: schema.TypeString,
			},
			{
				Name: "certificate_type",
				Type: schema.TypeString,
			},
			{
				Name: "customer_override",
				Type: schema.TypeBool,
			},
			{
				Name: "customer_override_valid_till",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "thumbprint",
				Type: schema.TypeString,
			},
			{
				Name: "valid_from",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "valid_till",
				Type: schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config rds.DescribeCertificatesInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeCertificates(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
