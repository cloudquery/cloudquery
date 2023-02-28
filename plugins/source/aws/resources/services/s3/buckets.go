package s3

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_buckets",
		Resolver:  fetchS3Buckets,
		Transform: transformers.TransformWithStruct(&models.WrappedBucket{}),
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveBucketARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			bucketEncryptionRules(),
			bucketLifecycles(),
			bucketGrants(),
			bucketCorsRules(),
			bucketWebsites(),
		},
	}
}

func fetchS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().S3
	response, err := svc.ListBuckets(ctx, nil, func(options *s3.Options) {
		options.Region = listBucketRegion(cl)
	})
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	buckets := make(chan types.Bucket)
	errs := make(chan error)
	for i := 0; i < fetchS3BucketsPoolSize; i++ {
		wg.Add(1)
		go fetchS3BucketsWorker(ctx, meta, buckets, errs, res, &wg)
	}
	go func() {
		defer close(buckets)
		for _, bucket := range response.Buckets {
			select {
			case <-ctx.Done():
				return
			case buckets <- bucket:
			}
		}
	}()
	done := make(chan struct{})
	go func() {
		for err = range errs {
			cl.Logger().Err(err).Msg("failed to fetch s3 bucket")
		}
		close(done)
	}()
	wg.Wait()
	close(errs)
	<-done

	return nil
}
