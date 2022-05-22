package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource ledgers --config gen.hcl --output .
func Ledgers() *schema.Table {
	return &schema.Table{
		Name:         "aws_qldb_ledgers",
		Resolver:     fetchQldbLedgers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("qldb"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the pipeline.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveQldbLedgerTags,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the ledger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date_time",
				Description: "The date and time, in epoch time format, when the ledger was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deletion_protection",
				Description: "The flag that prevents a ledger from being deleted by any user",
				Type:        schema.TypeBool,
			},
			{
				Name:        "encryption_status",
				Description: "The current state of encryption at rest for the ledger",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionDescription.EncryptionStatus"),
			},
			{
				Name:          "kms_key_arn",
				Description:   "The Amazon Resource Name (ARN) of the customer managed KMS key that the ledger uses for encryption at rest",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("EncryptionDescription.KmsKeyArn"),
			},
			{
				Name:          "inaccessible_kms_key_date_time",
				Description:   "The date and time, in epoch time format, when the KMS key first became inaccessible, in the case of an error",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("EncryptionDescription.InaccessibleKmsKeyDateTime"),
			},
			{
				Name:        "name",
				Description: "The name of the ledger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "permissions_mode",
				Description: "The permissions mode of the ledger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current status of the ledger.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_qldb_ledger_journal_kinesis_streams",
				Description:   "Information about an Amazon QLDB journal stream, including the Amazon Resource Name (ARN), stream name, creation time, current status, and the parameters of the original stream creation request.",
				Resolver:      fetchQldbLedgerJournalKinesisStreams,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "ledger_cq_id",
						Description: "Unique CloudQuery ID of aws_qldb_ledgers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "stream_arn",
						Description: "The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KinesisConfiguration.StreamArn"),
					},
					{
						Name:        "aggregation_enabled",
						Description: "Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("KinesisConfiguration.AggregationEnabled"),
					},
					{
						Name:        "ledger_name",
						Description: "The name of the ledger.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The current state of the QLDB journal stream.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stream_id",
						Description: "The UUID (represented in Base62-encoded text) of the QLDB journal stream.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stream_name",
						Description: "The user-defined name of the QLDB journal stream.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the QLDB journal stream.",
						Type:        schema.TypeString,
					},
					{
						Name:        "creation_time",
						Description: "The date and time, in epoch time format, when the QLDB journal stream was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "error_cause",
						Description: "The error message that describes the reason that a stream has a status of IMPAIRED or FAILED",
						Type:        schema.TypeString,
					},
					{
						Name:        "exclusive_end_time",
						Description: "The exclusive date and time that specifies when the stream ends",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "inclusive_start_time",
						Description: "The inclusive start date and time from which to start streaming journal data.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:          "aws_qldb_ledger_journal_s3_exports",
				Description:   "Information about a journal export job, including the ledger name, export ID, creation time, current status, and the parameters of the original export creation request.",
				Resolver:      fetchQldbLedgerJournalS3Exports,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "ledger_cq_id",
						Description: "Unique CloudQuery ID of aws_qldb_ledgers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "exclusive_end_time",
						Description: "The exclusive end date and time for the range of journal contents that was specified in the original export request.  This member is required.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "export_creation_time",
						Description: "The date and time, in epoch time format, when the export job was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "export_id",
						Description: "The UUID (represented in Base62-encoded text) of the journal export job.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "inclusive_start_time",
						Description: "The inclusive start date and time for the range of journal contents that was specified in the original export request.  This member is required.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "ledger_name",
						Description: "The name of the ledger.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal export job to do the following:  * Write objects into your Amazon Simple Storage Service (Amazon S3) bucket.  * (Optional) Use your customer managed key in Key Management Service (KMS) for server-side encryption of your exported data.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "bucket",
						Description: "The Amazon S3 bucket name in which a journal export job writes the journal contents",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3ExportConfiguration.Bucket"),
					},
					{
						Name:        "object_encryption_type",
						Description: "The Amazon S3 object encryption type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3ExportConfiguration.EncryptionConfiguration.ObjectEncryptionType"),
					},
					{
						Name:        "kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of a symmetric key in Key Management Service (KMS)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3ExportConfiguration.EncryptionConfiguration.KmsKeyArn"),
					},
					{
						Name:        "prefix",
						Description: "The prefix for the Amazon S3 bucket in which a journal export job writes the journal contents",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3ExportConfiguration.Prefix"),
					},
					{
						Name:        "status",
						Description: "The current state of the journal export job.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "output_format",
						Description: "The output format of the exported journal data.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchQldbLedgers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().QLDB
	config := qldb.ListLedgersInput{}
	for {
		response, err := svc.ListLedgers(ctx, &config, func(options *qldb.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		ledgers := make([]*qldb.DescribeLedgerOutput, 0, len(response.Ledgers))
		for _, l := range response.Ledgers {
			response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name}, func(o *qldb.Options) {
				o.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			ledgers = append(ledgers, response)
		}
		res <- ledgers
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func ResolveQldbLedgerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ledger := resource.Item.(*qldb.DescribeLedgerOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().QLDB
	response, err := svc.ListTagsForResource(ctx, &qldb.ListTagsForResourceInput{
		ResourceArn: ledger.Arn,
	}, func(options *qldb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, response.Tags))
}
func fetchQldbLedgerJournalKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalKinesisStreamsForLedgerInput{
		LedgerName: ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().QLDB.ListJournalKinesisStreamsForLedger(ctx, config, func(options *qldb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- response.Streams
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchQldbLedgerJournalS3Exports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalS3ExportsForLedgerInput{
		Name:       ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().QLDB.ListJournalS3ExportsForLedger(ctx, config, func(options *qldb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- response.JournalS3Exports
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
