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

const MAX_RECORD_SIZE = 1024000

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
		if len(dst.Bytes()) < MAX_RECORD_SIZE {
			recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
				Data: dst.Bytes(),
			})
		} else {
			c.logger.Warn().Msgf("skipping record because it is too large: %s", string(b))
		}
		if len(recordsBatchInput.Records) == 500 {
			err := c.sendBatch(ctx, recordsBatchInput, 0)
			if err != nil {
				return err
			}
			recordsBatchInput.Records = nil
		}
	}

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
