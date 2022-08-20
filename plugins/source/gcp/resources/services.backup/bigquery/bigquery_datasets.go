package bigquery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/googleapi"
)

func BigqueryDatasets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigquery_datasets",
		Description: "dataset resources in the project",
		Resolver:    fetchBigqueryDatasets,

		Multiplex: client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_time",
				Description: "The time when this dataset was created, in milliseconds since the epoch",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "default_encryption_configuration_kms_key_name",
				Description: "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table The BigQuery Service Account associated with your project requires access to this encryption key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultEncryptionConfiguration.KmsKeyName"),
			},
			{
				Name:        "default_partition_expiration_ms",
				Description: "The default partition expiration for all partitioned tables in the dataset, in milliseconds Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones The storage in a partition will have an expiration time of its partition time plus this value Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table If you provide an explicit timePartitioningexpirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "default_table_expiration_ms",
				Description: "The default lifetime of all tables in the dataset, in milliseconds The minimum value is 3600000 milliseconds (one hour) Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones When the expirationTime for a given table is reached, that table will be deleted automatically If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "description",
				Description: "A user-friendly description of the dataset",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "A hash of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "friendly_name",
				Description: "A descriptive name for the dataset",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The fully-qualified unique name of the dataset in the format projectId:datasetId The dataset name without the project name is given in the datasetId field When creating a new dataset, leave this field blank, and instead specify the datasetId field",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "The resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "The labels associated with this dataset You can use these to organize and group your datasets You can set this property when inserting or updating a dataset See Creating and Updating Dataset Labels for more information",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "last_modified_time",
				Description: "The date when this dataset or any of its tables was last modified, in milliseconds since the epoch",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "location",
				Description: "The geographic location where the dataset should reside The default value is US See details at https://cloudgooglecom/bigquery/docs/locations",
				Type:        schema.TypeString,
			},
			{
				Name:          "satisfies_pzs",
				Description:   "Reserved for future use",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("SatisfiesPZS"),
				IgnoreInTests: true,
			},
			{
				Name:        "self_link",
				Description: "A URL that can be used to access the resource again You can use this URL in Get or Update requests to the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			BigqueryDatasetAccesses(),
			BigqueryDatasetTables(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBigqueryDatasets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.BigQuery.Datasets.
			List(c.ProjectId).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		for _, d := range output.Datasets {
			dataset, err := c.Services.BigQuery.Datasets.
				Get(c.ProjectId, d.DatasetReference.DatasetId).Do()
			if err != nil {
				return errors.WithStack(err)
			}
			res <- dataset
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func isAccessErrorToIgnore(err error, projectId string) bool {
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == http.StatusBadRequest &&
			len(gerr.Errors) > 0 &&
			gerr.Errors[0].Reason == "invalid" &&
			gerr.Errors[0].Message == fmt.Sprintf("The project %s has not enabled BigQuery.", projectId) {
			return true
		}
	}
	return false
}
