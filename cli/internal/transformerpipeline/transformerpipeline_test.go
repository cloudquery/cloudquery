package transformerpipeline

import (
	"context"
	"errors"
	"io"
	"runtime"
	"testing"
	"time"

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

func TestWaitDeliveredWithoutTransformers(t *testing.T) {
	var (
		delivered []string
		release   = make(chan struct{})
	)
	pipeline, _, err := New(context.Background(), nil)
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(func(output []byte) error {
		<-release
		delivered = append(delivered, string(output))
		return nil
	}))
	go func() { _ = pipeline.RunBlocking() }()
	defer pipeline.Close()

	sent := make(chan struct{})
	go func() {
		defer close(sent)
		for _, input := range []string{"a", "b", "c"} {
			require.NoError(t, pipeline.Send([]byte(input)))
		}
	}()

	waited := make(chan error)
	go func() {
		<-sent
		waited <- pipeline.WaitDelivered(context.Background())
	}()
	for range 3 {
		release <- struct{}{}
	}
	require.NoError(t, <-waited)
	require.Equal(t, []string{"a", "b", "c"}, delivered)
}

func TestWaitDeliveredWithTransformers(t *testing.T) {
	pipeline, _, err := New(context.Background(), []plugin.Plugin_TransformClient{newIdentityTransformer()})
	require.NoError(t, err)
	require.Error(t, pipeline.WaitDelivered(context.Background()))
}

type blockingSendTransformer struct {
	reverserTransformer
	unblock chan struct{}
}

func (b *blockingSendTransformer) Send(req *plugin.Transform_Request) error {
	<-b.unblock
	return nil
}

func TestSendOnClosedPipelineDoesNotBlockGoroutine(t *testing.T) {
	unblockSend := make(chan struct{})
	blockingClient := &blockingSendTransformer{
		unblock: unblockSend,
	}

	pipeline, _, err := New(context.Background(), []plugin.Plugin_TransformClient{blockingClient})
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(func([]byte) error { return nil }))

	initialGoroutines := runtime.NumGoroutine()

	sendErrCh := make(chan error, 1)
	go func() {
		sendErrCh <- pipeline.Send([]byte("data"))
	}()

	time.Sleep(100 * time.Millisecond)
	pipeline.Close()

	select {
	case err := <-sendErrCh:
		require.ErrorIs(t, err, ErrPipelineClosed)
	case <-time.After(3 * time.Second):
		t.Fatal("pipeline.Send did not return within expected duration")
	}

	close(unblockSend)

	var exited bool
	for range 200 {
		if runtime.NumGoroutine() <= initialGoroutines {
			exited = true
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	require.True(t, exited, "send goroutine leaked; initial goroutines: %d, current: %d", initialGoroutines, runtime.NumGoroutine())
}

type fixedRecordsTransformer struct {
	reverserTransformer
	records [][]byte
	idx     int
}

func (m *fixedRecordsTransformer) Recv() (*plugin.Transform_Response, error) {
	if m.idx >= len(m.records) {
		return nil, io.EOF
	}
	rec := m.records[m.idx]
	m.idx++
	return &plugin.Transform_Response{Record: rec}, nil
}

func TestStartBlockingDeliversFinalRecordBeforeEOF(t *testing.T) {
	// A client that returns 3 records then EOF, with a 1 ms nextSendFn.
	// With buffered recvCh, the final record was dropped intermittently because
	// errCh and recvCh raced in select.
	for run := range 100 {
		records := [][]byte{[]byte("rec1"), []byte("rec2"), []byte("rec3")}
		client := &fixedRecordsTransformer{records: records}
		var received [][]byte
		wrapper := clientWrapper{
			client: client,
			nextSendFn: func(req *plugin.Transform_Request) error {
				time.Sleep(1 * time.Millisecond)
				received = append(received, req.Record)
				return nil
			},
			nextClose: func() error { return nil },
		}

		err := wrapper.startBlocking()
		require.NoError(t, err)
		require.Equal(t, records, received, "failed on run %d", run)
	}
}

type continuousRecordsTransformer struct {
	reverserTransformer
	closed chan struct{}
}

func (c *continuousRecordsTransformer) Recv() (*plugin.Transform_Response, error) {
	select {
	case <-c.closed:
		return nil, io.EOF
	default:
		return &plugin.Transform_Response{Record: []byte("continuous")}, nil
	}
}

func TestStartBlockingReaderExitsAfterNextSendFnError(t *testing.T) {
	client := &continuousRecordsTransformer{
		closed: make(chan struct{}),
	}
	defer close(client.closed)

	initialGoroutines := runtime.NumGoroutine()

	sendErr := errors.New("nextSendFn failed")
	wrapper := clientWrapper{
		client: client,
		nextSendFn: func(req *plugin.Transform_Request) error {
			return sendErr
		},
		nextClose: func() error { return nil },
	}

	err := wrapper.startBlocking()
	require.ErrorIs(t, err, sendErr)

	// Verify the reader goroutine spawned by startBlocking has exited
	var exited bool
	for range 200 {
		if runtime.NumGoroutine() <= initialGoroutines {
			exited = true
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	require.True(t, exited, "reader goroutine leaked; initial goroutines: %d, current: %d", initialGoroutines, runtime.NumGoroutine())
}
