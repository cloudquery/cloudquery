package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

const (
	MaxRecordSizeBytes = 1024000
	MaxBatchRecords    = 500
	MaxBatchSizeBytes  = 4194000
)

func (c *Client) Write(ctx context.Context, tables schema.Tables, record <-chan arrow.Record) error {
	parsedARN, err := arn.Parse(c.pluginSpec.StreamARN)
	if err != nil {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return err
	}
	arnResource := strings.Split(parsedARN.Resource, "/")
	if len(arnResource) != 2 {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return fmt.Errorf("invalid firehose stream ARN")
	}
	recordsBatchInput := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(arnResource[1]),
	}
	batchSize := 0

	for rec := range record {
		tableName, ok := rec.Schema().Metadata().GetValue(schema.MetadataTableName)
		if !ok {
			return fmt.Errorf("%q metadata key not found", schema.MetadataTableName)
		}

		table := tables.Get(tableName)
		if table == nil {
			return fmt.Errorf("table %s not found", tableName)
		}

		for row := int64(0); row < rec.NumRows(); row++ {
			jsonObj := make(map[string]any, rec.NumCols()+1)
			for i := range rec.Columns() {
				jsonObj[rec.ColumnName(i)] = rec.Column(i).GetOneForMarshal(int(row))
			}
			// Add table name to the json object
			// TODO: This should be added to the SDK so that it can be used for other plugins as well
			jsonObj["_cq_table_name"] = tableName
			b, err := json.Marshal(jsonObj)
			if err != nil {
				return err
			}
			dst := &bytes.Buffer{}
			err = json.Compact(dst, b)
			if err != nil {
				return err
			}
			if len(dst.Bytes()) > MaxRecordSizeBytes {
				c.logger.Warn().Msgf("skipping record because it is too large: %s", string(b))
				continue
			}

			// If adding this record would exceed the batch size, send the batch
			if len(dst.Bytes())+batchSize > min(c.spec.BatchSizeBytes, MaxBatchSizeBytes) {
				err := c.sendBatch(ctx, recordsBatchInput, 0)
				if err != nil {
					return err
				}
				recordsBatchInput.Records = nil
				batchSize = 0
			}

			recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
				Data: dst.Bytes(),
			})
			// Store a running total of the batch size
			batchSize += len(dst.Bytes())

			// Send the batch if it is full
			if len(recordsBatchInput.Records) == min(c.spec.BatchSize, MaxBatchRecords) {
				err := c.sendBatch(ctx, recordsBatchInput, 0)
				if err != nil {
					return err
				}
				// Reset the batch
				recordsBatchInput.Records = nil
				batchSize = 0
			}
		}
	}
	// Send the last batch
	return c.sendBatch(ctx, recordsBatchInput, 0)
}

func (c *Client) sendBatch(ctx context.Context, recordsBatchInput *firehose.PutRecordBatchInput, count int) error {
	if count == *c.pluginSpec.MaxRetries {
		return fmt.Errorf("max retries reached")
	}
	if recordsBatchInput == nil || len(recordsBatchInput.Records) == 0 {
		return nil
	}
	time.Sleep(time.Duration(count) * time.Second)
	resp, err := c.firehoseClient.PutRecordBatch(ctx, recordsBatchInput)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to write to firehose")
		return err
	}
	retryRecords := getFailedRecords(recordsBatchInput, resp)
	atomic.AddUint64(&c.metrics.Writes, uint64(len(recordsBatchInput.Records)-len(retryRecords.Records)))
	return c.sendBatch(ctx, retryRecords, count+1)
}

func getFailedRecords(recordsBatchInput *firehose.PutRecordBatchInput, resp *firehose.PutRecordBatchOutput) *firehose.PutRecordBatchInput {
	retryRecords := &firehose.PutRecordBatchInput{
		DeliveryStreamName: recordsBatchInput.DeliveryStreamName,
	}
	for i, r := range resp.RequestResponses {
		if r.RecordId == nil {
			retryRecords.Records = append(retryRecords.Records, recordsBatchInput.Records[i])
		}
	}
	return retryRecords
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
