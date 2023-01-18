package client

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	recordsBatchInput := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(c.pluginSpec.StreamARN),
	}

	for _, resource := range data {
		jsonObj := make(map[string]any, len(table.Columns))
		for i := range resource {
			jsonObj[table.Columns[i].Name] = resource[i]
		}
		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
			Data: b,
		})
	}

	resp, err := c.firehoseClient.PutRecordBatch(ctx, recordsBatchInput)

	if aws.ToInt32(resp.FailedPutCount) > int32(0) {
		// Handle partial success
		c.logger.Warn().Msg("partial success in writing to firehose")
	}
	return err
}
