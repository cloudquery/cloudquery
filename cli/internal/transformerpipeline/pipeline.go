package transformerpipeline

import (
	"context"
	"errors"
	"fmt"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
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
}

func New(ctx context.Context, transformClients []plugin.Plugin_TransformClient) (*TransformerPipeline, context.Context, error) {
	var (
		eg, gctx = errgroup.WithContext(ctx)
		tp       = &TransformerPipeline{clientWrappers: make([]clientWrapper, len(transformClients)), eg: eg}
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

	if lp.clientWrappers[0].closed {
		return fmt.Errorf("cannot send data to a closed pipeline")
	}

	return lp.clientWrappers[0].client.Send(&plugin.Transform_Request{Record: data})
}

func (lp *TransformerPipeline) OnOutput(fn func([]byte) error) error {
	if fn == nil {
		return errors.New("argument to OnOutput cannot be nil")
	}
	lp.clientWrappers[len(lp.clientWrappers)-1].nextSendFn = func(req *plugin.Transform_Request) error {
		return fn(req.Record)
	}
	return nil
}

func (lp *TransformerPipeline) Close() error {
	// Closing the pipeline happens on both source as well as destination close.
	// Not handling this will result in a close of closed channel panic.
	if lp.clientWrappers[0].closed {
		return nil
	}
	lp.clientWrappers[0].closed = true

	// Close the first transformer. The rest will follow gracefully, otherwise records will be lost.
	return lp.clientWrappers[0].client.CloseSend()
}

type clientWrapper struct {
	i          int
	client     plugin.Plugin_TransformClient
	nextSendFn func(*plugin.Transform_Request) error
	nextClose  func() error
	closed     bool
}

func (s clientWrapper) startBlocking() error {
	if s.nextSendFn == nil {
		return errors.New("nextSendFn is nil")
	}
	for {
		data, err := s.client.Recv()
		if err == io.EOF {
			err := s.nextClose()
			return err
		}
		if err != nil {
			return err
		}
		if err := s.nextSendFn(
			&plugin.Transform_Request{Record: data.Record},
		); err != nil {
			return err
		}
	}
}
