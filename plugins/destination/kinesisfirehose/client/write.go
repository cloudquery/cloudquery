package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

// TODO: Verify that metrics are being updated correctly
// TODO: Verify errors are propagated
// TODO: Clean up send batch and retry logic

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	parsedARN, err := arn.Parse(c.pluginSpec.StreamARN)
	if err != nil {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return err
	}
	resource := strings.Split(parsedARN.Resource, "/")
	if len(resource) != 2 {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return fmt.Errorf("invalid firehose stream ARN")
	}
	recordsBatchInput := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(resource[1]),
	}

	for resource := range res {
		table := tables.Get(resource.TableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", resource.TableName))
		}

		jsonObj := make(map[string]any, len(table.Columns)+1)
		for i := range resource.Data {
			jsonObj[table.Columns[i].Name] = resource.Data[i]
		}
		// Add table name to the json object
		// TODO: This should be added to the SDK so that it can be used for other plugins as well
		jsonObj["_cq_table_name"] = table.Name
		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		dst := &bytes.Buffer{}
		err = json.Compact(dst, b)
		if err != nil {
			return err
		}
		if len(dst.Bytes()) < 1024000 {
			recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
				Data: dst.Bytes(),
			})
		} else {
			// log that skipping because record is too large
			c.logger.Warn().Msgf("skipping record because it is too large: %s", string(b))
		}
		if len(recordsBatchInput.Records) == 500 {
			c.sendBatch(ctx, recordsBatchInput)
			atomic.AddUint64(&c.metrics.Writes, uint64(len(recordsBatchInput.Records)))
			recordsBatchInput.Records = nil
		}
	}

	if len(recordsBatchInput.Records) > 0 {
		c.sendBatch(ctx, recordsBatchInput)
		atomic.AddUint64(&c.metrics.Writes, uint64(len(recordsBatchInput.Records)))
	}

	return nil
}

func (c *Client) sendBatch(ctx context.Context, recordsBatchInput *firehose.PutRecordBatchInput) error {
	resp, err := c.firehoseClient.PutRecordBatch(ctx, recordsBatchInput)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to write to firehose")
		return err
	}
	if aws.ToInt32(resp.FailedPutCount) > int32(0) {
		// Handle partial success
		c.logger.Warn().Msg("partial success in writing to firehose")
		retryRecords := &firehose.PutRecordBatchInput{
			DeliveryStreamName: recordsBatchInput.DeliveryStreamName,
		}
		for i, r := range resp.RequestResponses {
			if r.RecordId == nil {
				retryRecords.Records = append(retryRecords.Records, recordsBatchInput.Records[i])
			}
		}
		return c.retryBatch(ctx, retryRecords, 1)
	} else {
		c.logger.Warn().Msgf("wrote: %d records", len(recordsBatchInput.Records))
	}

	return nil
}

func (c *Client) retryBatch(ctx context.Context, recordsBatchInput *firehose.PutRecordBatchInput, count int) error {
	if count == 5 {
		return fmt.Errorf("max retries reached")
	}

	time.Sleep(time.Duration(count) * time.Second)
	retryRecords := &firehose.PutRecordBatchInput{
		DeliveryStreamName: recordsBatchInput.DeliveryStreamName,
	}
	resp, err := c.firehoseClient.PutRecordBatch(ctx, recordsBatchInput)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to write to firehose")
		return err
	}
	if aws.ToInt32(resp.FailedPutCount) > int32(0) {
		for i, r := range resp.RequestResponses {
			if r.RecordId == nil {
				retryRecords.Records = append(retryRecords.Records, recordsBatchInput.Records[i])
			}
		}
		return c.retryBatch(ctx, retryRecords, count+1)
	}
	return nil

}
