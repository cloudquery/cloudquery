package transformerpipeline

import (
	"testing"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/stretchr/testify/require"
)

func TestIdentityTransformer(t *testing.T) {
	transformer := newIdentityTransformer()

	// Test sending and receiving a message
	testRecord := []byte("test record")
	req := &plugin.Transform_Request{Record: testRecord}

	go func() {
		require.NoError(t, transformer.Send(req))
	}()

	resp, err := transformer.Recv()
	require.NoError(t, err)

	// Since it's an identityTransformer, the record should be the same
	require.Equal(t, testRecord, resp.Record, "Records should be the same")

	// Test closing the channel
	require.NoError(t, transformer.CloseSend())

	// Test channel is closed after call to CloseSend
	_, ok := <-transformer.ch
	require.False(t, ok, "Channel should be closed but it's not")
}
