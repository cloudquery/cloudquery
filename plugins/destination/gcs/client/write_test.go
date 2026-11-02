package client

import (
	"errors"
	"io"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

type fakeObjectWriter struct {
	writeErr   error
	closeErr   error
	writeCalls int
	closeCalls int
}

func (w *fakeObjectWriter) Write(p []byte) (int, error) {
	w.writeCalls++
	if w.writeErr != nil {
		return 0, w.writeErr
	}
	return len(p), nil
}

func (w *fakeObjectWriter) Close() error {
	w.closeCalls++
	return w.closeErr
}

type fakeHandle struct {
	contentErr   error
	footerErr    error
	contentCalls int
}

func (h *fakeHandle) WriteContent([]arrow.RecordBatch) error {
	h.contentCalls++
	return h.contentErr
}

func (h *fakeHandle) WriteFooter() error { return h.footerErr }

type fakeFileType struct {
	headerErr error
	handle    types.Handle
}

func (f *fakeFileType) WriteHeader(io.Writer, *schema.Table) (types.Handle, error) {
	if f.headerErr != nil {
		return nil, f.headerErr
	}
	return f.handle, nil
}

func (f *fakeFileType) Read(types.ReaderAtSeeker, *schema.Table, chan<- arrow.RecordBatch) error {
	return nil
}

func testRecord(t *testing.T) arrow.RecordBatch {
	t.Helper()
	table := &schema.Table{
		Name:    "test",
		Columns: []schema.Column{{Name: "name", Type: arrow.BinaryTypes.String}},
	}
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*array.StringBuilder).Append("foo")
	return bldr.NewRecordBatch()
}

func insertsChannel(records ...arrow.RecordBatch) <-chan *message.WriteInsert {
	ch := make(chan *message.WriteInsert, len(records))
	for _, record := range records {
		ch <- &message.WriteInsert{Record: record}
	}
	close(ch)
	return ch
}

func TestWriteTableClosesWriterOnError(t *testing.T) {
	errHeader := errors.New("header failed")
	errContent := errors.New("content failed")
	errFooter := errors.New("footer failed")

	for _, tc := range []struct {
		name       string
		headerErr  error
		contentErr error
		footerErr  error
		wantErr    error
	}{
		{name: "write header error", headerErr: errHeader, wantErr: errHeader},
		{name: "write content error", contentErr: errContent, wantErr: errContent},
		{name: "write footer error", footerErr: errFooter, wantErr: errFooter},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := &fakeObjectWriter{}
			ft := &fakeFileType{
				headerErr: tc.headerErr,
				handle:    &fakeHandle{contentErr: tc.contentErr, footerErr: tc.footerErr},
			}

			err := writeTable(insertsChannel(testRecord(t)), ft, func(*schema.Table) objectWriter { return w })

			require.Equal(t, tc.wantErr, err)
			require.Equal(t, 1, w.closeCalls, "writer must be closed exactly once")
		})
	}
}

func TestWriteTableJoinsCloseError(t *testing.T) {
	errFooter := errors.New("footer failed")
	errClose := errors.New("close failed")

	w := &fakeObjectWriter{closeErr: errClose}
	ft := &fakeFileType{handle: &fakeHandle{footerErr: errFooter}}

	err := writeTable(insertsChannel(testRecord(t)), ft, func(*schema.Table) objectWriter { return w })

	require.ErrorIs(t, err, errFooter)
	require.ErrorIs(t, err, errClose)
	require.Equal(t, 1, w.closeCalls)
}

func TestWriteTableEmptyChannel(t *testing.T) {
	opened := 0
	ft := &fakeFileType{handle: &fakeHandle{}}

	err := writeTable(insertsChannel(), ft, func(*schema.Table) objectWriter {
		opened++
		return &fakeObjectWriter{}
	})

	require.NoError(t, err)
	require.Zero(t, opened, "no object should be created for an empty channel")
}

func TestWriteTableSuccess(t *testing.T) {
	w := &fakeObjectWriter{}
	handle := &fakeHandle{}
	ft := &fakeFileType{handle: handle}

	require.NoError(t, writeTable(insertsChannel(testRecord(t), testRecord(t)), ft, func(*schema.Table) objectWriter { return w }))
	require.Equal(t, 2, handle.contentCalls)
	require.Equal(t, 1, w.closeCalls)
}

func TestOnceCloserClosesUnderlyingWriterOnce(t *testing.T) {
	w := &fakeObjectWriter{closeErr: errors.New("close failed")}
	closer := newOnceCloser(w)

	require.EqualError(t, closer.Close(), "close failed")
	require.EqualError(t, closer.Close(), "close failed")
	require.Equal(t, 1, w.closeCalls, "underlying writer must be closed exactly once")
}
