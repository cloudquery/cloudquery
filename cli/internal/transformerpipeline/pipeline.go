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
	cancel         context.CancelFunc
}

func New(ctx context.Context, transformClients []plugin.Plugin_TransformClient) (*TransformerPipeline, context.Context, error) {
	var (
		eg, gctx = errgroup.WithContext(ctx)
		tp       = &TransformerPipeline{clientWrappers: make([]clientWrapper, len(transformClients)), eg: eg}
	)

	// Create a cancellable context for pipeline-internal goroutines (e.g. the
	// Recv loop in startBlocking). This prevents goroutine leaks when Close()
	// is called: the Recv goroutine will notice the cancellation and exit
	// instead of blocking forever on dead channels.
	pipelineCtx, pipelineCancel := context.WithCancel(gctx)
	tp.cancel = pipelineCancel

	// Make sure there's at least one transformer
	if len(transformClients) == 0 {
		tp.clientWrappers = append(tp.clientWrappers, clientWrapper{client: newIdentityTransformer(), ctx: pipelineCtx})
	}

	// Wrap the clients to add orchestration logic
	for i, client := range transformClients {
		tp.clientWrappers[i] = clientWrapper{i: i, client: client, ctx: pipelineCtx}
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

	sendCh := make(chan error)

	// Send can block forever (e.g. if grpc buffer is full), so we run it asynchronously
	// and check if pipeline is closed every second.
	go func() {
		err := lp.clientWrappers[0].client.Send(&plugin.Transform_Request{Record: data})
		sendCh <- err
	}()

	// Use a ticker instead of time.After to avoid leaking timers. Each call to
	// time.After allocates a new timer that isn't GC'd until it fires, which
	// causes a memory leak in hot loops.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case err := <-sendCh:
			return err
		case <-ticker.C: // Check if pipeline is closed every second
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
		}
		return err
	}
	return nil
}

func (lp *TransformerPipeline) Close() {
	// Close() can happen in any goroutine, and closing is not thread safe.
	// Instead of closing, we set a flag that we check on send/recv.
	lp.clientWrappers[0].isClosed.Store(true)
	// Cancel the pipeline context so that Recv goroutines in startBlocking
	// can detect shutdown and exit cleanly instead of leaking.
	lp.cancel()
}

type clientWrapper struct {
	i          int
	client     plugin.Plugin_TransformClient
	nextSendFn func(*plugin.Transform_Request) error
	nextClose  func() error
	isClosed   atomic.Bool
	ctx        context.Context
}

func (s *clientWrapper) startBlocking() error {
	if s.nextSendFn == nil {
		return errors.New("nextSendFn is nil")
	}

	recvCh := make(chan *plugin.Transform_Request)
	errCh := make(chan error)

	// Recv can block forever (e.g. if transformer decides to), so
	// we run it asynchronously and check if pipeline is closed every second.
	//
	// The goroutine uses s.ctx to detect when the pipeline is closed. Without
	// this, the goroutine would be orphaned after startBlocking returns because
	// it would block forever trying to send to recvCh/errCh (whose only reader
	// has already exited).
	go func() {
		for {
			data, err := s.client.Recv()
			if err != nil {
				select {
				case errCh <- err:
				case <-s.ctx.Done():
					return
				}
			} else {
				select {
				case recvCh <- &plugin.Transform_Request{Record: data.Record}:
				case <-s.ctx.Done():
					return
				}
			}
		}
	}()

	// Use a ticker instead of time.After to avoid leaking timers.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: // Check if pipeline is closed every second
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
