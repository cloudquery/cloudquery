package transformerpipeline

import (
	"context"
	"io"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"google.golang.org/grpc/metadata"
)

// identityTransformer is a transformer mock that does nothing to the data
// it exists so that we can have at least one transformer in the pipeline
type identityTransformer struct {
	ch chan []byte
}

func newIdentityTransformer() *identityTransformer {
	return &identityTransformer{
		ch: make(chan []byte),
	}
}

func (t *identityTransformer) Send(req *plugin.Transform_Request) error {
	t.ch <- req.Record
	return nil
}

func (t *identityTransformer) Recv() (*plugin.Transform_Response, error) {
	bs, ok := <-t.ch
	if !ok {
		return nil, io.EOF
	}
	return &plugin.Transform_Response{Record: bs}, nil
}

// Close the channel!
func (t *identityTransformer) CloseSend() error {
	close(t.ch)
	return nil
}

// Must satisfy the Plugin_TransformClient interface
func (identityTransformer) Header() (metadata.MD, error) { return metadata.MD{}, nil }
func (identityTransformer) Trailer() metadata.MD         { return metadata.MD{} }
func (identityTransformer) Context() context.Context     { return nil }
func (identityTransformer) SendMsg(m any) error          { return nil }
func (identityTransformer) RecvMsg(m any) error          { return nil }
