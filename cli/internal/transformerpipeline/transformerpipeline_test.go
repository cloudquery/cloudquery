package transformerpipeline

import (
	"context"
	"io"
	"runtime"
	"sync"
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

// blockingTransformer parks Send and Recv on channels driven by the test, so a call can
// be held open past the point its owner gave up on it. That is the state that turns an
// unbuffered hand-off channel into a permanent goroutine leak.
type blockingTransformer struct {
	sendEntered chan []byte
	sendAck     chan struct{}
	recv        chan *plugin.Transform_Response
	done        chan struct{}
	closeOnce   sync.Once
}

func newBlockingTransformer() *blockingTransformer {
	return &blockingTransformer{
		sendEntered: make(chan []byte, 1),
		sendAck:     make(chan struct{}),
		recv:        make(chan *plugin.Transform_Response),
		done:        make(chan struct{}),
	}
}

func (t *blockingTransformer) Send(req *plugin.Transform_Request) error {
	t.sendEntered <- req.Record
	<-t.sendAck
	return nil
}

func (t *blockingTransformer) Recv() (*plugin.Transform_Response, error) {
	select {
	case r := <-t.recv:
		return r, nil
	case <-t.done:
		return nil, io.EOF
	}
}

func (t *blockingTransformer) CloseSend() error {
	t.closeOnce.Do(func() { close(t.done) })
	return nil
}

func (t *blockingTransformer) releaseSend() { t.sendAck <- struct{}{} }

// Must satisfy the Plugin_TransformClient interface
func (blockingTransformer) Header() (metadata.MD, error) { return metadata.MD{}, nil }
func (blockingTransformer) Trailer() metadata.MD         { return metadata.MD{} }
func (blockingTransformer) Context() context.Context     { return nil }
func (blockingTransformer) SendMsg(m any) error          { return nil }
func (blockingTransformer) RecvMsg(m any) error          { return nil }

// settledGoroutines waits until the goroutine count stops moving and returns it. The
// pipeline polls its closed flag on a one second timer, so goroutines belonging to tests
// that already finished can still be unwinding when this test starts. It measures from the
// test goroutine itself so no helper goroutine is counted.
func settledGoroutines(t *testing.T) int {
	t.Helper()

	last := runtime.NumGoroutine()
	stable := 0
	for start := time.Now(); stable < 3 && time.Since(start) < 3*time.Second; {
		time.Sleep(20 * time.Millisecond)
		if n := runtime.NumGoroutine(); n == last {
			stable++
		} else {
			stable, last = 0, n
		}
	}
	return last
}

func requireGoroutinesDrained(t *testing.T, before int) {
	t.Helper()

	for start := time.Now(); time.Since(start) < 2*time.Second; {
		if runtime.NumGoroutine() <= before {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	require.LessOrEqual(t, runtime.NumGoroutine(), before, "pipeline goroutine leaked")
}

func TestSendDoesNotLeakGoroutineWhenPipelineClosesMidSend(t *testing.T) {
	before := settledGoroutines(t)

	transformer := newBlockingTransformer()
	pipeline, _, err := New(context.Background(), []plugin.Plugin_TransformClient{transformer})
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(func([]byte) error { return nil }))

	sendResult := make(chan error, 1)
	go func() { sendResult <- pipeline.Send([]byte("in flight")) }()

	<-transformer.sendEntered
	pipeline.Close()
	// Send stops waiting once it notices the pipeline closed, while client.Send is still parked.
	require.ErrorIs(t, <-sendResult, ErrPipelineClosed)

	// Now client.Send returns and hands its result to a channel nobody will ever read.
	transformer.releaseSend()

	requireGoroutinesDrained(t, before)
}

func TestStartBlockingDoesNotLeakRecvGoroutineWhenItReturns(t *testing.T) {
	before := settledGoroutines(t)

	transformer := newBlockingTransformer()
	wrapper := &clientWrapper{
		client:     transformer,
		nextSendFn: func(*plugin.Transform_Request) error { return nil },
		nextClose:  transformer.CloseSend,
	}
	wrapper.isClosed.Store(true)

	require.NoError(t, wrapper.startBlocking())

	// startBlocking has returned, so its Recv goroutine must be able to abandon the io.EOF
	// it is holding instead of waiting for a consumer that no longer exists.
	requireGoroutinesDrained(t, before)
}
