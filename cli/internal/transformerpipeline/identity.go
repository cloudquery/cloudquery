package transformerpipeline

import (
	"context"
	"io"
	"sync"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"google.golang.org/grpc/metadata"
)

// identityTransformer is a transformer mock that does nothing to the data
// it exists so that we can have at least one transformer in the pipeline
type identityTransformer struct {
	ch      chan []byte
	closeCh chan struct{}
	once    sync.Once
}

func newIdentityTransformer() *identityTransformer {
	return &identityTransformer{
		ch:      make(chan []byte),
		closeCh: make(chan struct{}),
	}
}

func (t *identityTransformer) Send(req *plugin.Transform_Request) error {
	// ch is never closed, so this send can never panic. A send racing
	// CloseSend observes closeCh and returns instead of staying parked
	// on the unbuffered channel forever.
	select {
	case t.ch <- req.Record:
		return nil
	case <-t.closeCh:
		return ErrPipelineClosed
	}
}

func (t *identityTransformer) Recv() (*plugin.Transform_Response, error) {
	select {
	case bs := <-t.ch:
		return &plugin.Transform_Response{Record: bs}, nil
	case <-t.closeCh:
		// A sender parked mid-rendezvous when closure was signalled gets
		// one chance to deliver before we report EOF.
		select {
		case bs := <-t.ch:
			return &plugin.Transform_Response{Record: bs}, nil
		default:
		}
		return nil, io.EOF
	}
}

// CloseSend signals closure on closeCh. The data channel is deliberately
// left open: closing it would panic any Send parked on it.
func (t *identityTransformer) CloseSend() error {
	t.once.Do(func() { close(t.closeCh) })
	return nil
}

// Must satisfy the Plugin_TransformClient interface
func (*identityTransformer) Header() (metadata.MD, error) { return metadata.MD{}, nil }
func (*identityTransformer) Trailer() metadata.MD         { return metadata.MD{} }
func (*identityTransformer) Context() context.Context     { return nil }
func (*identityTransformer) SendMsg(m any) error          { return nil }
func (*identityTransformer) RecvMsg(m any) error          { return nil }
