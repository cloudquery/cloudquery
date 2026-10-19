package client

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/plugins/destination/s3/v7/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestSupportsFlush(t *testing.T) {
	zero := int64(0)
	noTimeout := configtype.NewDuration(0)
	tests := []struct {
		name string
		spec *spec.Spec
		want bool
	}{
		{
			name: "no spec",
		},
		{
			name: "path with uuid",
			spec: &spec.Spec{Bucket: bucket, Region: region, Path: "dir/{{TABLE}}/{{UUID}}.json", FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON}},
			want: true,
		},
		{
			name: "default path",
			spec: &spec.Spec{
				Bucket: bucket, Region: region, Path: "dir",
				BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &noTimeout,
				FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON},
			},
			want: true,
		},
		{
			name: "no_rotate",
			spec: &spec.Spec{Bucket: bucket, Region: region, Path: "dir", NoRotate: true, FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON}},
		},
		{
			name: "path without uuid",
			spec: &spec.Spec{
				Bucket: bucket, Region: region, Path: "dir/{{TABLE}}.json",
				BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &noTimeout,
				FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b []byte
			if tc.spec != nil {
				var err error
				b, err = json.Marshal(tc.spec)
				require.NoError(t, err)
			}
			p := plugin.NewPlugin("s3", "development", New, plugin.WithFlushSupport())
			require.NoError(t, p.Init(context.Background(), b, plugin.NewClientOptions{NoConnection: true}))
			require.Equal(t, tc.want, p.SupportsFlush())
		})
	}
}

func TestFlush(t *testing.T) {
	zero := int64(0)
	noTimeout := configtype.NewDuration(0)
	for _, ft := range []filetypes.FormatType{
		filetypes.FormatTypeCSV,
		filetypes.FormatTypeJSON,
		filetypes.FormatTypeParquet,
	} {
		t.Run(string(ft), func(t *testing.T) {
			t.Parallel()
			testFlush(t, &spec.Spec{
				Bucket:         bucket,
				Region:         region,
				Path:           t.TempDir()[1:] + "/{{TABLE}}/{{UUID}}.{{FORMAT}}",
				BatchSize:      &zero,
				BatchSizeBytes: &zero,
				BatchTimeout:   &noTimeout,
				FileSpec:       filetypes.FileSpec{Format: ft},
			})
		})
	}
}

// testFlush checks that every flush uploads the open object, and that the next records go to a new object.
func testFlush(t *testing.T, s *spec.Spec) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var client *Client
	p := plugin.NewPlugin("s3", "development", func(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
		c, err := New(ctx, logger, spec, opts)
		if err != nil {
			return nil, err
		}
		client = c.(*Client)
		return c, nil
	}, plugin.WithFlushSupport())
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))
	t.Cleanup(func() { require.NoError(t, p.Close(context.Background())) })
	require.True(t, p.SupportsFlush())

	table := &schema.Table{
		Name:    fmt.Sprintf("cq_test_flush_%d", time.Now().UnixNano()),
		Columns: []schema.Column{{Name: "id", Type: arrow.PrimitiveTypes.Int64}},
	}
	msgs := make(chan message.WriteMessage)
	writeErr := make(chan error, 1)
	go func() { writeErr <- p.Write(ctx, msgs) }()

	send := func(msg message.WriteMessage) {
		select {
		case msgs <- msg:
		case err := <-writeErr:
			t.Fatalf("write returned early: %v", err)
		}
	}
	insert := func(ids ...int64) {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		defer bldr.Release()
		for _, id := range ids {
			bldr.Field(0).(*array.Int64Builder).Append(id)
		}
		send(&message.WriteInsert{Record: bldr.NewRecordBatch()})
	}
	flush := func() {
		msg := &message.WriteFlush{}
		send(msg)
		select {
		case err := <-msg.Result():
			require.NoError(t, err)
		case <-ctx.Done():
			t.Fatal("flush did not complete")
		}
	}
	prefix := path.Dir(client.spec.ReplacePathVariables(table.Name, "uuid", time.Now().UTC(), "")) + "/"
	objects := func() int {
		out, err := client.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{Bucket: aws.String(s.Bucket), Prefix: aws.String(prefix)})
		require.NoError(t, err)
		return len(out.Contents)
	}

	send(&message.WriteMigrateTable{Table: table})
	insert(1, 2, 3)
	flush()
	require.Equal(t, 1, objects())
	insert(4, 5)
	flush()
	require.Equal(t, 2, objects())
	close(msgs)
	require.NoError(t, <-writeErr)
	require.Equal(t, 2, objects())
}
