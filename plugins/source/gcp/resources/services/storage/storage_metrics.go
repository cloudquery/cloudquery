package storage

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/api/monitoring/v3"
)

type storageSetter func(metric *storageMetric, value *monitoring.TypedValue)

type storageMetric struct {
	BucketName        string
	AclOperationCount int64
	ObjectCount       int64
	TotalSize         int64
}

const (
	queryACLCount = `
		fetch gcs_bucket
		| metric 'storage.googleapis.com/authz/acl_based_object_access_count'
		| group_by 1d,
			[value_acl_based_object_access_count_aggregate:
			   aggregate(value.acl_based_object_access_count)]
		| every 1d
		| group_by [resource.bucket_name],
			[value_acl_based_object_access_count_aggregate_aggregate:
       			aggregate(value_acl_based_object_access_count_aggregate)]`

	queryTotalObjects = `
		fetch gcs_bucket
		| metric 'storage.googleapis.com/storage/object_count'
		| group_by 1d, [value_object_count_mean: mean(value.object_count)]
		| every 1d`

	queryTotalBucketSize = `
		fetch gcs_bucket
		| metric 'storage.googleapis.com/storage/total_bytes'
		| group_by 1d, [value_total_bytes_mean: mean(value.total_bytes)]
		| every 1d`
)

func Metrics() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_metrics",
		Description: "storage metrics collecting by cloud monitoring service",
		Resolver:    fetchStorageMetrics,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "bucket_name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "bucket_name",
				Description: "Name of the bucket metric is associated with",
				Type:        schema.TypeString,
			},
			{
				Name:        "acl_operation_count",
				Description: "Usage of ACL operations count in 24 hour period",
				Type:        schema.TypeBigInt,
			},

			{
				Name:        "object_count",
				Description: "Total number of objects per bucket, grouped by storage class. This value is measured once per day, and there might be a delay after measuring before the value becomes available in Cloud Monitoring.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "total_size",
				Description: "Total size of all objects in the bucket (in bytes), grouped by storage class. This value is measured once per day, and there might be a delay after measuring before the value becomes available in Cloud Monitoring.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

func fetchStorageMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	metrics := make(map[string]*storageMetric)
	cl := meta.(*client.Client)
	if err := doTimeSeriesCall(ctx, cl, queryACLCount, func(metric *storageMetric, value *monitoring.TypedValue) {
		metric.AclOperationCount = cast.ToInt64(value.Int64Value)
	}, metrics); err != nil {
		return errors.WithStack(err)
	}
	if err := doTimeSeriesCall(ctx, cl, queryTotalObjects, func(metric *storageMetric, value *monitoring.TypedValue) {
		metric.ObjectCount = cast.ToInt64(value.DoubleValue)
	}, metrics); err != nil {
		return errors.WithStack(err)
	}

	if err := doTimeSeriesCall(ctx, cl, queryTotalBucketSize, func(metric *storageMetric, value *monitoring.TypedValue) {
		metric.TotalSize = cast.ToInt64(value.DoubleValue)
	}, metrics); err != nil {
		return errors.WithStack(err)
	}

	totalMetrics := make([]*storageMetric, 0, len(metrics))
	for _, m := range metrics {
		totalMetrics = append(totalMetrics, m)
	}
	res <- totalMetrics
	return nil
}

func doTimeSeriesCall(ctx context.Context, cl *client.Client, query string, setter storageSetter, metrics map[string]*storageMetric) error {
	call := cl.Services.Monitoring.Projects.TimeSeries.Query(fmt.Sprintf("projects/%s", cl.ProjectId), &monitoring.QueryTimeSeriesRequest{
		Query: query,
	})
	response, err := call.Do()
	if err != nil {
		return errors.WithStack(err)
	}

	if response.TimeSeriesData == nil {
		return nil
	}
	bucketIndex := getDescriptorIndex(response.TimeSeriesDescriptor.LabelDescriptors, "resource.bucket_name")
	if bucketIndex == -1 {
		return errors.WithStack(fmt.Errorf("failed to get bucket index for timeseries call"))
	}

	for _, data := range response.TimeSeriesData {
		bucketName := data.LabelValues[bucketIndex].StringValue
		bucket, ok := metrics[bucketName]
		if !ok {
			bucket = &storageMetric{BucketName: bucketName}
			metrics[bucketName] = bucket
		}
		// there should only be one point of data per bucket
		setter(bucket, data.PointData[0].Values[0])
	}
	return nil
}

func getDescriptorIndex(descriptors []*monitoring.LabelDescriptor, value string) int {
	for i, d := range descriptors {
		if d.Key == value {
			return i
		}
	}
	return -1
}
