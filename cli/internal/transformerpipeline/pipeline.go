package transformerpipeline

import (
	"context"
	"errors"
	"io"
	"sync/atomic"
	"time"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"golang.org/x/sync/errgroup"
)

var (
	ErrPipelineClosed = errors.New("pipeline is closed")
)

// TransformerPipeline runs a pipeline of transform clients.
//
// Ideally we'd just call the result of each transform to the next one, but transformations are not synchronous calls,
// so orchestration is needed. That's what this does: it hides the orchestration of the transform clients.
//
// Use it like this:
//
// - Construct a new TransformerPipeline with `New`. Give it a context and a slice of transform clients.
// - Register a callback for transformed records with `OnOutput`.
// - Start all transformers with `RunBlocking`.
// - Send records to the pipeline with `Send`.
// - When done, close the pipeline with `Close`. Otherwise, `RunBlocking` won't finish.
type TransformerPipeline struct {
	clientWrappers []clientWrapper
	eg             *errgroup.Group

	// without transformers, every sent record is delivered to the output
	passthrough bool
	sent        atomic.Int64
	delivered   atomic.Int64
	onDelivered chan struct{}
}

func New(ctx context.Context, transformClients []plugin.Plugin_TransformClient) (*TransformerPipeline, context.Context, error) {
	var (
		eg, gctx = errgroup.WithContext(ctx)
		tp       = &TransformerPipeline{
			clientWrappers: make([]clientWrapper, len(transformClients)),
			eg:             eg,
			passthrough:    len(transformClients) == 0,
			onDelivered:    make(chan struct{}, 1),
		}
	)

	// Make sure there's at least one transformer
	if len(transformClients) == 0 {
		tp.clientWrappers = append(tp.clientWrappers, clientWrapper{client: newIdentityTransformer()})
	}

	// Wrap the clients to add orchestration logic
	for i, client := range transformClients {
		tp.clientWrappers[i] = clientWrapper{i: i, client: client}
	}

	// Connect each client to the next one
	for i := 0; i < len(transformClients)-1; i++ {
		tp.clientWrappers[i].nextSendFn = tp.clientWrappers[i+1].client.Send
		tp.clientWrappers[i].nextClose = tp.clientWrappers[i+1].client.CloseSend
	}

	// The last client in the pipeline has nothing else to close
	tp.clientWrappers[len(tp.clientWrappers)-1].nextClose = func() error { return nil }

	// The last client sends to the output. This connection happens in `OnOutput`.

	return tp, gctx, nil
}

func (lp *TransformerPipeline) RunBlocking() error {
	for i := len(lp.clientWrappers) - 1; i >= 0; i-- {
		lp.eg.Go(lp.clientWrappers[i].startBlocking)
	}
	return lp.eg.Wait()
}

func (lp *TransformerPipeline) Send(data []byte) error {
	// Constructor makes sure that there is at least one "identity" transform client
	if lp.clientWrappers[len(lp.clientWrappers)-1].nextSendFn == nil {
		return errors.New("OnOutput must be registered before Send is called, otherwise what do I do with the transformed data?")
	}

	if lp.clientWrappers[0].isClosed.Load() {
		return ErrPipelineClosed
	}

	sendCh := make(chan error, 1)

	// Send can block forever (e.g. if grpc buffer is full), so we run it asynchronously
	// and check if pipeline is closed every second. The channel is buffered so the
	// goroutine can always hand back its result, even after Send gave up on waiting
	// because the pipeline closed.
	go func() {
		err := lp.clientWrappers[0].client.Send(&plugin.Transform_Request{Record: data})
		sendCh <- err
	}()

	for {
		select {
		case err := <-sendCh:
			if err == nil {
				lp.sent.Add(1)
			}
			return err
		case <-time.After(1 * time.Second): // Check if pipeline is closed every second
			if lp.clientWrappers[0].isClosed.Load() {
				return ErrPipelineClosed
			}
		}
	}
}

func (lp *TransformerPipeline) OnOutput(fn func([]byte) error) error {
	if fn == nil {
		return errors.New("argument to OnOutput cannot be nil")
	}
	lp.clientWrappers[len(lp.clientWrappers)-1].nextSendFn = func(req *plugin.Transform_Request) error {
		err := fn(req.Record)
		if err != nil {
			// Our undocumented convention is that destination errors are unrecoverable. Thus, at this
			// point we close the pipeline.
			lp.Close()
			return err
		}
		lp.delivered.Add(1)
		select {
		case lp.onDelivered <- struct{}{}:
		default:
		}
		return nil
	}
	return nil
}

// WaitDelivered returns once every record sent so far has been passed to the OnOutput callback.
// Only pipelines without transformers support it, as transformers may not output a record for each input.
func (lp *TransformerPipeline) WaitDelivered(ctx context.Context) error {
	if !lp.passthrough {
		return errors.New("cannot wait for records to pass through transformers")
	}
	target := lp.sent.Load()
	for lp.delivered.Load() < target {
		select {
		case <-lp.onDelivered:
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second): // Check if pipeline is closed every second
			if lp.clientWrappers[0].isClosed.Load() {
				return ErrPipelineClosed
			}
		}
	}
	return nil
}

func (lp *TransformerPipeline) Close() {
	// Close() can happen in any goroutine, and closing is not thread safe.
	// Instead of closing, we set a flag that we check on send/recv.
	lp.clientWrappers[0].isClosed.Store(true)
}

type clientWrapper struct {
	i          int
	client     plugin.Plugin_TransformClient
	nextSendFn func(*plugin.Transform_Request) error
	nextClose  func() error
	isClosed   atomic.Bool
}

func (s *clientWrapper) startBlocking() error {
	if s.nextSendFn == nil {
		return errors.New("nextSendFn is nil")
	}

	recvCh := make(chan *plugin.Transform_Request)
	errCh := make(chan error)

	// Recv can block forever (e.g. if transformer decides to), so
	// we run it asynchronously and check if pipeline is closed every second.
	// done is closed when this function returns, which is the only way the
	// goroutine below can abandon a pending hand-off instead of parking on it
	// for the rest of the process.
	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			if s.isClosed.Load() {
				return
			}

			data, err := s.client.Recv()
			if err != nil {
				select {
				case errCh <- err:
				case <-done:
					return
				}
				continue
			}

			select {
			case recvCh <- &plugin.Transform_Request{Record: data.Record}:
			case <-done:
				return
			}
		}
	}()

	for {
		select {
		case <-time.After(1 * time.Second): // Check if pipeline is closed every second
			if s.isClosed.Load() {
				return s.nextClose()
			}
		case req := <-recvCh: // Propagate records to next transformer
			if err := s.nextSendFn(req); err != nil {
				return err
			}
		case err := <-errCh:
			if err == io.EOF {
				return s.nextClose()
			}
			return err
		}
	}
}
