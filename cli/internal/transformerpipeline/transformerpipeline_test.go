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

func TestPipelineCloseDoesNotLeakGoroutines(t *testing.T) {
	// Warmup: first use may spawn lazy goroutines, so measure the baseline
	// only after one full pipeline cycle has completed.
	runCycle := func() {
		transformers := []plugin.Plugin_TransformClient{newIdentityTransformer(), newIdentityTransformer()}
		pipeline, _, err := New(context.Background(), transformers)
		require.NoError(t, err)
		require.NoError(t, pipeline.OnOutput(func([]byte) error { return nil }))
		done := make(chan struct{})
		go func() {
			defer close(done)
			require.NoError(t, pipeline.Send([]byte("ping")))
			pipeline.Close()
		}()
		require.NoError(t, pipeline.RunBlocking())
		<-done
	}
	runCycle()

	before := runtime.NumGoroutine()
	for i := 0; i < 5; i++ {
		runCycle()
	}
	waitForGoroutines(t, before)
}

func TestSendUnblocksWhenPipelineClosesMidSend(t *testing.T) {
	blocking := newBlockingTransformer()
	transformers := []plugin.Plugin_TransformClient{blocking, newIdentityTransformer()}
	pipeline, _, err := New(context.Background(), transformers)
	require.NoError(t, err)
	require.NoError(t, pipeline.OnOutput(func([]byte) error { return nil }))

	before := runtime.NumGoroutine()

	runDone := make(chan error, 1)
	go func() { runDone <- pipeline.RunBlocking() }()

	sendDone := make(chan error, 1)
	go func() { sendDone <- pipeline.Send([]byte("stuck")) }()

	// Wait until the Send goroutine is parked inside the transformer's Send,
	// then close the pipeline underneath it.
	<-blocking.sendStarted
	pipeline.Close()

	select {
	case err := <-sendDone:
		require.ErrorIs(t, err, ErrPipelineClosed)
	case <-time.After(10 * time.Second):
		t.Fatal("Send did not return after pipeline close")
	}

	select {
	case err := <-runDone:
		require.NoError(t, err)
	case <-time.After(10 * time.Second):
		t.Fatal("RunBlocking did not return after pipeline close")
	}

	// Release the parked Send. Its goroutine must be able to deliver the
	// result and exit even though Send itself already returned.
	close(blocking.sendRelease)
	waitForGoroutines(t, before)
}

// blockingTransformer parks in Send until the test releases it, and in Recv
// until CloseSend is called. Used to exercise close-during-send behavior.
type blockingTransformer struct {
	sendStarted chan struct{}
	sendRelease chan struct{}
	recvDone    chan struct{}
	recvOnce    sync.Once
	sendOnce    sync.Once
}

func newBlockingTransformer() *blockingTransformer {
	return &blockingTransformer{
		sendStarted: make(chan struct{}),
		sendRelease: make(chan struct{}),
		recvDone:    make(chan struct{}),
	}
}

func (t *blockingTransformer) Send(_ *plugin.Transform_Request) error {
	t.sendOnce.Do(func() { close(t.sendStarted) })
	<-t.sendRelease
	return nil
}

func (t *blockingTransformer) Recv() (*plugin.Transform_Response, error) {
	<-t.recvDone
	return nil, io.EOF
}

func (t *blockingTransformer) CloseSend() error {
	t.recvOnce.Do(func() { close(t.recvDone) })
	return nil
}

// Must satisfy the Plugin_TransformClient interface
func (*blockingTransformer) Header() (metadata.MD, error) { return metadata.MD{}, nil }
func (*blockingTransformer) Trailer() metadata.MD         { return metadata.MD{} }
func (*blockingTransformer) Context() context.Context     { return nil }
func (*blockingTransformer) SendMsg(m any) error          { return nil }
func (*blockingTransformer) RecvMsg(m any) error          { return nil }

// waitForGoroutines polls until the goroutine count returns to baseline.
// require.Eventually cannot be used for this: it evaluates the condition in
// a fresh goroutine per tick, which counts itself and never settles.
func waitForGoroutines(t *testing.T, before int) {
	t.Helper()
	for deadline := time.Now().Add(10 * time.Second); time.Now().Before(deadline); {
		if runtime.NumGoroutine() <= before {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("goroutines leaked: before=%d now=%d", before, runtime.NumGoroutine())
}
