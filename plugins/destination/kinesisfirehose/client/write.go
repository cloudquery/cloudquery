package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
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

	for _, resource := range data {
		jsonObj := make(map[string]any, len(table.Columns)+1)
		for i := range resource {
			jsonObj[table.Columns[i].Name] = resource[i]
		}
		// Add table name to the json object
		// TODO: This should be added to the SDK so that it can be used for other plugins as well
		jsonObj["_cq_table_name"] = table.Name
		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
			Data: b,
		})
		if len(recordsBatchInput.Records) == 1000 {
			c.sendBatch(ctx, recordsBatchInput)
			recordsBatchInput.Records = nil
		}
	}

	return c.sendBatch(ctx, recordsBatchInput)
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
	}
	return nil
}
