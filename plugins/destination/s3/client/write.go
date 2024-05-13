package client

import (
    "context"
    "encoding/json"
    "io"
    "regexp"
    "time"

    "github.com/apache/arrow/go/v16/arrow"
    "github.com/apache/arrow/go/v16/arrow/array"
    "github.com/apache/arrow/go/v16/arrow/memory"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    awstypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/cloudquery/filetypes/v4"
    "github.com/cloudquery/plugin-sdk/v4/message"
    "github.com/cloudquery/plugin-sdk/v4/schema"
    "github.com/cloudquery/plugin-sdk/v4/types"
    "github.com/google/uuid"
)

var reInvalidJSONKey = regexp.MustCompile(`\W`)

func (c *Client) createObject(ctx context.Context, table *schema.Table, objKey string) (*filetypes.Stream, error) {
    s, err := c.Client.StartStream(table, func(r io.Reader) error {
        params := &s3.PutObjectInput{
            Bucket:      aws.String(c.spec.Bucket),
            Key:         aws.String(objKey),
            Body:        r,
            ContentType: aws.String(c.spec.GetContentType()),
            ACL:         awstypes.ObjectCannedACL(c.spec.ACL),
        }

        sseConfiguration := c.spec.ServerSideEncryptionConfiguration
        if sseConfiguration != nil {
            params.SSEKMSKeyId = &sseConfiguration.SSEKMSKeyId
            params.ServerSideEncryption = sseConfiguration.ServerSideEncryption
        }

        _, err := c.uploader.Upload(ctx, params)
        return err
    })
    return s, err
}

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
    var s *filetypes.Stream

    for msg := range msgs {
        if s == nil {
            table := msg.GetTable()
            objKey := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC(), c.syncID)
            // if object was already initialized, use the same key
            // We don't need any locking here because all messages for the same table are processed sequentially
            if val, ok := c.initializedTables[table.Name]; ok {
                objKey = val
                delete(c.initializedTables, table.Name)
            }

            var err error
            s, err = c.createObject(ctx, table, objKey)
            if err != nil {
                return err
            }
        }

        if c.spec.Athena {
            var err error
            msg.Record, err = sanitizeRecordJSONKeys(msg.Record)
            if err != nil {
                _ = s.FinishWithError(err)
                return err
            }
        }

        if err := s.Write([]arrow.Record{msg.Record}); err != nil {
            _ = s.FinishWithError(err)
            return err
        }
    }

    return s.Finish()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
    return c.writer.Write(ctx, msgs)
}

func (c *Client) MigrateTable(ctx context.Context, ch <-chan *message.WriteMigrateTable) error {
    for msg := range ch {
        if !c.spec.GenerateEmptyObjects {
            continue
        }
        table := msg.GetTable()
        objKey := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC(), c.syncID)
        // We don't need any locking here because all messages for the same table are processed sequentially
        c.initializedTables[table.Name] = objKey
        s, err := c.createObject(ctx, table, objKey)
        if err != nil {
            return err
        }
        err = s.Finish()
        if err != nil {
            return err
        }
    }
    return nil
}

// sanitizeRecordJSONKeys replaces all invalid characters in JSON keys with underscores. This is required
// for compatibility with Athena.
func sanitizeRecordJSONKeys(record arrow.Record) (arrow.Record, error) {
    cols := make([]arrow.Array, record.NumCols())
    for i, col := range record.Columns() {
        if arrow.TypeEqual(col.DataType(), types.NewJSONType()) {
            b := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewJSONType()))
            for r := 0; r < int(record.NumRows()); r++ {
                if col.IsNull(r) {
                    b.AppendNull()
                    continue
                }
                data, err := sanitizeJSONRawMessage(col.GetOneForMarshal(r).(json.RawMessage))
                if err != nil {
                    return nil, err
                }
                b.Append(data)
            }
            cols[i] = b.NewArray()
            continue
        }
        cols[i] = col
    }
    return array.NewRecord(record.Schema(), cols, record.NumRows()), nil
}

func sanitizeJSONRawMessage(rawMessage json.RawMessage) (any, error) {
    var data any
    err := json.Unmarshal(rawMessage, &data)
    if err != nil {
        return nil, err
    }

    return sanitizeJSONKeysForObject(data), nil
}

func sanitizeJSONKeysForObject(data any) any {
    // we only care about objects that have keys: present either while nesting or in arrays
    switch data := data.(type) {
    case map[string]any:
        res := make(map[string]any, len(data))
        for k, v := range data {
            res[reInvalidJSONKey.ReplaceAllString(k, "_")] = sanitizeJSONKeysForObject(v)
        }
        return res
    case []any:
        for i, el := range data {
            data[i] = sanitizeJSONKeysForObject(el)
        }
        return data
    default:
        return data
    }
}
