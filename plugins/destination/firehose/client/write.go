package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) Write(ctx context.Context, messages <-chan message.WriteMessage) error {
	parsedARN, err := arn.Parse(c.spec.StreamARN)
	if err != nil {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return err
	}
	arnResource := strings.Split(parsedARN.Resource, "/")
	if len(arnResource) != 2 {
		c.logger.Error().Err(err).Msg("invalid firehose stream ARN")
		return errors.New("invalid firehose stream ARN")
	}
	recordsBatchInput := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(arnResource[1]),
	}
	batchSize := 0

	for m := range messages {
		switch m := m.(type) {
		case *message.WriteDeleteStale:
			c.logger.Warn().Str("table", m.TableName).Msg("DeleteStale not implemented")
			continue
		case *message.WriteMigrateTable:
			c.logger.Warn().Str("table", m.Table.Name).Msg("Migrate not implemented")
			continue
		case *message.WriteDeleteRecord:
			c.logger.Warn().Str("table", m.TableName).Msg("DeleteRecord not implemented")
			continue
		case *message.WriteInsert:
		// ok, handle outside of switch
		default:
			return fmt.Errorf("unsupported message type: %T", m)
		}
		ins := m.(*message.WriteInsert)
		table, rec := ins.GetTable(), ins.Record

		for row := int64(0); row < rec.NumRows(); row++ {
			jsonObj := make(map[string]any, rec.NumCols()+1)
			for i := range rec.Columns() {
				jsonObj[rec.ColumnName(i)] = rec.Column(i).GetOneForMarshal(int(row))
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
			if len(dst.Bytes()) > c.spec.MaxRecordSizeBytes {
				c.logger.Warn().Msgf("skipping record because it is too large: %s", string(b))
				continue
			}

			// If adding this record would exceed the batch size, send the batch
			if len(dst.Bytes())+batchSize > c.spec.MaxBatchSizeBytes {
				err := c.sendBatch(ctx, recordsBatchInput, 0)
				if err != nil {
					return err
				}
				recordsBatchInput.Records = recordsBatchInput.Records[:0]
				batchSize = 0
			}

			recordsBatchInput.Records = append(recordsBatchInput.Records, types.Record{
				Data: dst.Bytes(),
			})
			// Store a running total of the batch size
			batchSize += len(dst.Bytes())

			// Send the batch if it is full
			if len(recordsBatchInput.Records) >= c.spec.MaxBatchRecords {
				err := c.sendBatch(ctx, recordsBatchInput, 0)
				if err != nil {
					return err
				}
				// Reset the batch
				recordsBatchInput.Records = recordsBatchInput.Records[:0]
				batchSize = 0
			}
		}
	}
	// Send the last batch
	return c.sendBatch(ctx, recordsBatchInput, 0)
}

func (c *Client) sendBatch(ctx context.Context, recordsBatchInput *firehose.PutRecordBatchInput, count int) error {
	if count >= c.spec.MaxRetries {
		return errors.New("max retries reached")
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
