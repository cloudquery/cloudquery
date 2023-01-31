package client

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type startWriteMsg struct {
	tables schema.Tables
	err    chan error
}

type endWriteMsg struct {
	tables schema.Tables
	err    chan error
}

type migrateMsg struct {
	tables schema.Tables
	err    chan error
}

type readMsg struct {
	table     *schema.Table
	source    string
	err       chan error
	resources chan []any
}

type closeMsg struct {
	err chan error
}

type Client struct {
	destination.UnimplementedManagedWriter
	destination.DefaultReverseTransformer
	logger  zerolog.Logger
	spec    specs.Destination
	csvSpec Spec
	metrics destination.Metrics
	writers map[string]*tableWriter

	startWriteChan chan *startWriteMsg
	writeChan      chan *destination.ClientResource
	endWriteChan   chan *endWriteMsg
	migrateChan    chan *migrateMsg
	readChan       chan *readMsg
	closeChan      chan *closeMsg

	wg *sync.WaitGroup
}

type tableWriter struct {
	writer *csv.Writer
	file   *os.File
	count  uint64
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("csv destination only supports append mode")
	}

	c := &Client{
		logger:         logger.With().Str("module", "csv-dest").Logger(),
		spec:           spec,
		startWriteChan: make(chan *startWriteMsg),
		writeChan:      make(chan *destination.ClientResource),
		endWriteChan:   make(chan *endWriteMsg),
		migrateChan:    make(chan *migrateMsg),
		readChan:       make(chan *readMsg),
		closeChan:      make(chan *closeMsg),
		writers:        make(map[string]*tableWriter),
		wg:             &sync.WaitGroup{},
	}

	if err := spec.UnmarshalSpec(&c.csvSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal csv spec: %w", err)
	}
	c.csvSpec.SetDefaults()

	if err := os.MkdirAll(c.csvSpec.Directory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	c.wg.Add(1)
	go func() {
		c.listen()
	}()

	return c, nil
}

func (c *Client) close() {
	for tableName, w := range c.writers {
		w.writer.Flush()
		if err := w.file.Close(); err != nil {
			c.logger.Error().Err(err).Str("table", tableName).Msg("failed to close file")
		}
		delete(c.writers, tableName)
	}
}

func (c *Client) Close(ctx context.Context) error {
	msg := &closeMsg{
		err: make(chan error),
	}
	c.closeChan <- msg
	return <-msg.err
}

func (c *Client) listen() {
	for {
		select {
		case msg := <-c.startWriteChan:
			msg.err <- c.startWrite(msg.tables)
		case msg := <-c.endWriteChan:
			c.endWrite(msg.tables)
			msg.err <- nil
		case resource := <-c.writeChan:
			c.writeResource(resource)
		case msg := <-c.migrateChan:
			msg.err <- c.migrate(msg.tables)
		case msg := <-c.readChan:
			msg.err <- c.read(msg.table, msg.source, msg.resources)
		case msg := <-c.closeChan:
			c.close()
			msg.err <- nil
		}
	}
}
