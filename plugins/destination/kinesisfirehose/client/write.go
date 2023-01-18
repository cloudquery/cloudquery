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
