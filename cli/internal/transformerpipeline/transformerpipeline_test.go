package transformerpipeline

import (
	"context"
	"io"
	"testing"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestTransformerPipelineDoesntChangeInputsWithTwoIdentityTransformers(t *testing.T) {
	var (
		inputs          = []string{"test data 1", "test data 2", "test data 3"}
		expectedOutputs = []string{"test data 1", "test data 2", "test data 3"}
		actualOutputs   = []string{}
		recordOutputs   = func(output []byte) error { actualOutputs = append(actualOutputs, string(output)); return nil }
	)

	transformers := []plugin.Plugin_TransformClient{newIdentityTransformer(), newIdentityTransformer()}

	pipeline, _, err := New(context.Background(), transformers)
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(recordOutputs))

	// `Send` and `Close` affect the buffer of the initial transformer, so
	// they should succeed even if `RunBlocking` hasn't run yet and will
	// take a while to start the `Recv` loops.
	//
	// In this case, the transformer implementations block `Send` with an
	// unbuffered channel, so this goroutine will block until the pipeline
	// starts.
	go func() {
		for _, input := range inputs {
			require.NoError(t, pipeline.Send([]byte(input)))
		}
		pipeline.Close()
	}()

	// Blocks until pipeline is closed and all messages passed through
	// so above goroutine must have finished after this line.
	require.NoError(t, pipeline.RunBlocking())

	require.Equal(t, expectedOutputs, actualOutputs)
}

func TestTransformerPipelineReversesInputs(t *testing.T) {
	var (
		inputs          = []string{"test data 1", "test data 2", "test data 3"}
		expectedOutputs = []string{"1 atad tset", "2 atad tset", "3 atad tset"}
		actualOutputs   = []string{}
		recordOutputs   = func(output []byte) error { actualOutputs = append(actualOutputs, string(output)); return nil }
	)

	transformers := []plugin.Plugin_TransformClient{newReverserTransformer()}

	pipeline, _, err := New(context.Background(), transformers)
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(recordOutputs))

	// `Send` and `Close` affect the buffer of the initial transformer, so
	// they should succeed even if `RunBlocking` hasn't run yet and will
	// take a while to start the `Recv` loops.
	//
	// In this case, the transformer implementations block `Send` with an
	// unbuffered channel, so this goroutine will block until the pipeline
	// starts.
	go func() {
		for _, input := range inputs {
			require.NoError(t, pipeline.Send([]byte(input)))
		}
		pipeline.Close()
	}()

	// Blocks until pipeline is closed and all messages passed through
	// so above goroutine must have finished after this line.
	require.NoError(t, pipeline.RunBlocking())

	require.Equal(t, expectedOutputs, actualOutputs)
}

func TestTransformerPipelineDoesntChangeInputsWithTwoReversers(t *testing.T) {
	var (
		inputs          = []string{"test data 1", "test data 2", "test data 3"}
		expectedOutputs = []string{"test data 1", "test data 2", "test data 3"} // Reversed twice!
		actualOutputs   = []string{}
		recordOutputs   = func(output []byte) error { actualOutputs = append(actualOutputs, string(output)); return nil }
	)

	transformers := []plugin.Plugin_TransformClient{newReverserTransformer(), newReverserTransformer()}

	pipeline, _, err := New(context.Background(), transformers)
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(recordOutputs))

	// `Send` and `Close` affect the buffer of the initial transformer, so
	// they should succeed even if `RunBlocking` hasn't run yet and will
	// take a while to start the `Recv` loops.
	//
	// In this case, the transformer implementations block `Send` with an
	// unbuffered channel, so this goroutine will block until the pipeline
	// starts.
	go func() {
		for _, input := range inputs {
			require.NoError(t, pipeline.Send([]byte(input)))
		}
		pipeline.Close()
	}()

	// Blocks until pipeline is closed and all messages passed through
	// so above goroutine must have finished after this line.
	require.NoError(t, pipeline.RunBlocking())

	require.Equal(t, expectedOutputs, actualOutputs)
}

// reverserTransformer is a transformer mock that reverses the bytes, as runes, of the data
type reverserTransformer struct {
	ch chan []byte
}

func newReverserTransformer() *reverserTransformer {
	return &reverserTransformer{
		ch: make(chan []byte),
	}
}

func (t *reverserTransformer) Send(req *plugin.Transform_Request) error {
	t.ch <- req.Record
	return nil
}

func (t *reverserTransformer) Recv() (*plugin.Transform_Response, error) {
	bs, ok := <-t.ch
	if !ok {
		return nil, io.EOF
	}
	reversed := []rune(string(bs))
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	reversedBytes := []byte(string(reversed))
	return &plugin.Transform_Response{Record: reversedBytes}, nil
}

// Close the channel!
func (t *reverserTransformer) CloseSend() error {
	close(t.ch)
	return nil
}

// Must satisfy the Plugin_TransformClient interface
func (reverserTransformer) Header() (metadata.MD, error) { return metadata.MD{}, nil }
func (reverserTransformer) Trailer() metadata.MD         { return metadata.MD{} }
func (reverserTransformer) Context() context.Context     { return nil }
func (reverserTransformer) SendMsg(m any) error          { return nil }
func (reverserTransformer) RecvMsg(m any) error          { return nil }
