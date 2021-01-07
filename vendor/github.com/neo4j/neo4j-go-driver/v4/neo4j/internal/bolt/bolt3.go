/*
 * Copyright (c) 2002-2020 "Neo4j,"
 * Neo4j Sweden AB [http://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package bolt

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/internal/packstream"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

const (
	bolt3_ready        = iota // Ready for use
	bolt3_streaming           // Receiving result from auto commit query
	bolt3_pendingtx           // Transaction has been requested but not applied
	bolt3_tx                  // Transaction pending
	bolt3_streamingtx         // Receiving result from a query within a transaction
	bolt3_failed              // Recoverable error, needs reset
	bolt3_dead                // Non recoverable protocol or connection error
	bolt3_unauthorized        // Initial state, not sent hello message with authentication
)

type internalTx3 struct {
	mode      db.AccessMode
	bookmarks []string
	timeout   time.Duration
	txMeta    map[string]interface{}
}

func (i *internalTx3) toMeta() map[string]interface{} {
	meta := map[string]interface{}{}
	if i.mode == db.ReadMode {
		meta["mode"] = "r"
	}
	if len(i.bookmarks) > 0 {
		meta["bookmarks"] = i.bookmarks
	}
	ms := int(i.timeout.Nanoseconds() / 1e6)
	if ms > 0 {
		meta["tx_timeout"] = ms
	}
	if len(i.txMeta) > 0 {
		meta["tx_metadata"] = i.txMeta
	}
	return meta
}

type bolt3 struct {
	state         int
	txId          db.TxHandle
	currStream    *stream
	conn          net.Conn
	serverName    string
	out           *outgoing
	in            *incoming
	connId        string
	logId         string
	serverVersion string
	tfirst        int64        // Time that server started streaming
	pendingTx     *internalTx3 // Stashed away when tx started explcitly
	bookmark      string       // Last bookmark
	birthDate     time.Time
	log           log.Logger
	err           error // Last fatal error
}

func NewBolt3(serverName string, conn net.Conn, log log.Logger) *bolt3 {
	b := &bolt3{
		state:      bolt3_unauthorized,
		conn:       conn,
		serverName: serverName,
		in:         &incoming{buf: make([]byte, 4096)},
		birthDate:  time.Now(),
		log:        log,
	}
	b.out = &outgoing{
		chunker: newChunker(),
		packer:  &packstream.Packer{},
		onErr: func(err error) {
			if b.err == nil {
				b.err = err
			}
			b.state = bolt3_dead
		},
	}
	return b
}

func (b *bolt3) ServerName() string {
	return b.serverName
}

func (b *bolt3) ServerVersion() string {
	return b.serverVersion
}

// Sets b.err and b.state on failure
func (b *bolt3) receiveMsg() interface{} {
	msg, err := b.in.next(b.conn)
	if err != nil {
		b.err = err
		b.log.Error(log.Bolt3, b.logId, b.err)
		b.state = bolt3_dead
		return nil
	}
	return msg
}

// Receives a message that is assumed to be a success response or a failure in response
// to a sent command.
// Sets b.err and b.state on failure
func (b *bolt3) receiveSuccess() *success {
	switch v := b.receiveMsg().(type) {
	case *success:
		return v
	case *db.Neo4jError:
		b.state = bolt3_failed
		b.err = v
		if v.Classification() == "ClientError" {
			// These could include potentially large cypher statement, only log to debug
			b.log.Debugf(log.Bolt3, b.logId, "%s", v)
		} else {
			b.log.Error(log.Bolt3, b.logId, v)
		}
		return nil
	default:
		// Receive failed, state has been set
		if b.err != nil {
			return nil
		}
		// Unexpected message received
		b.state = bolt3_dead
		b.err = errors.New("Expected success or database error")
		b.log.Error(log.Bolt3, b.logId, b.err)
		return nil
	}
}

func (b *bolt3) connect(auth map[string]interface{}, userAgent string) error {
	if err := b.assertState(bolt3_unauthorized); err != nil {
		return err
	}

	hello := map[string]interface{}{
		"user_agent": userAgent,
	}
	// Merge authentication info into hello message
	for k, v := range auth {
		_, exists := hello[k]
		if exists {
			continue
		}
		hello[k] = v
	}

	// Send hello message and wait for confirmation
	b.out.appendHello(hello)
	if b.out.send(b.conn); b.err != nil {
		return b.err
	}

	succ := b.receiveSuccess()
	if b.err != nil {
		return b.err
	}

	b.connId = succ.connectionId
	b.logId = fmt.Sprintf("%s@%s", b.connId, b.serverName)
	b.serverVersion = succ.server

	// Transition into ready state
	b.state = bolt3_ready
	b.log.Infof(log.Bolt3, b.logId, "Connected")
	return nil
}

func (b *bolt3) TxBegin(txConfig db.TxConfig) (db.TxHandle, error) {
	// Ok, to begin transaction while streaming auto-commit, just empty the stream and continue.
	if b.state == bolt3_streaming {
		if err := b.bufferStream(); err != nil {
			return 0, err
		}
	}

	if err := b.assertState(bolt3_ready); err != nil {
		return 0, err
	}

	tx := &internalTx3{
		mode:      txConfig.Mode,
		bookmarks: txConfig.Bookmarks,
		timeout:   txConfig.Timeout,
		txMeta:    txConfig.Meta,
	}

	// If there are bookmarks, begin the transaction immediately to get the error from the
	// server early on. Requires a network roundtrip.
	if len(tx.bookmarks) > 0 {
		b.out.appendBegin(tx.toMeta())
		if b.out.send(b.conn); b.err != nil {
			return 0, b.err
		}
		if b.receiveSuccess(); b.err != nil {
			return 0, b.err
		}
		b.state = bolt3_tx
	} else {
		// Stash this into pending internal tx
		b.pendingTx = tx
		b.state = bolt3_pendingtx
	}
	b.txId = db.TxHandle(time.Now().Unix())
	return b.txId, nil
}

// Should NOT set b.err or change b.state as this is used to guard from
// misuse from clients that stick to their connections when they shouldn't.
func (b *bolt3) assertTxHandle(h1, h2 db.TxHandle) error {
	if h1 != h2 {
		err := errors.New("Invalid transaction handle")
		b.log.Error(log.Bolt3, b.logId, err)
		return err
	}
	return nil
}

// Should NOT set b.err or b.state since the connection is still valid
func (b *bolt3) assertState(allowed ...int) error {
	// Forward prior error instead, this former error is probably the
	// root cause of any state error. Like a call to Run with malformed
	// cypher causes an error and another call to Commit would cause the
	// state to be wrong. Do not log this.
	if b.err != nil {
		return b.err
	}
	for _, a := range allowed {
		if b.state == a {
			return nil
		}
	}
	err := errors.New(fmt.Sprintf("Invalid state %d, expected: %+v", b.state, allowed))
	b.log.Error(log.Bolt3, b.logId, err)
	return err
}

func (b *bolt3) TxCommit(txh db.TxHandle) error {
	if err := b.assertTxHandle(b.txId, txh); err != nil {
		return err
	}

	// Nothing to do, a transaction started but no commands were issued on it, server is unaware
	if b.state == bolt3_pendingtx {
		b.state = bolt3_ready
		return nil
	}

	// Consume pending stream if any to turn state from streamingtx to tx
	// Access to streams outside of tx boundary is not allowed, therefore we should discard
	// the stream (not buffer).
	if b.state == bolt3_streamingtx {
		if err := b.discardStream(); err != nil {
			return err
		}
	}

	// Should be in vanilla tx state now
	if err := b.assertState(bolt3_tx); err != nil {
		return err
	}

	// Send request to server to commit
	b.out.appendCommit()
	if b.out.send(b.conn); b.err != nil {
		return b.err
	}

	// Evaluate server response
	succ := b.receiveSuccess()
	if b.err != nil {
		return b.err
	}
	// Keep track of bookmark
	if len(succ.bookmark) > 0 {
		b.bookmark = succ.bookmark
	}

	// Transition into ready state
	b.state = bolt3_ready
	return nil
}

func (b *bolt3) TxRollback(txh db.TxHandle) error {
	if err := b.assertTxHandle(b.txId, txh); err != nil {
		return err
	}

	// Nothing to do, a transaction started but no commands were issued on it
	if b.state == bolt3_pendingtx {
		b.state = bolt3_ready
		return nil
	}

	// Can not send rollback while still streaming, consume to turn state into tx
	// Access to streams outside of tx boundary is not allowed, therefore we should discard
	// the stream (not buffer).
	if b.state == bolt3_streamingtx {
		if err := b.discardStream(); err != nil {
			return err
		}
	}

	// Should be in vanilla tx state now
	if err := b.assertState(bolt3_tx); err != nil {
		return err
	}

	// Send rollback request to server
	b.out.appendRollback()
	if b.out.send(b.conn); b.err != nil {
		return b.err
	}

	// Receive rollback confirmation
	if b.receiveSuccess(); b.err != nil {
		return b.err
	}

	b.state = bolt3_ready
	return nil
}

// Discards all records in current stream
func (b *bolt3) discardStream() error {
	if b.state != bolt3_streaming && b.state != bolt3_streamingtx {
		// Nothing to do
		return nil
	}

	var (
		sum *db.Summary
		err error
	)
	for sum == nil && err == nil {
		_, sum, err = b.receiveNext()
	}
	return err
}

// Collects all records in current stream
func (b *bolt3) bufferStream() error {
	if b.state != bolt3_streaming && b.state != bolt3_streamingtx {
		// Nothing to do
		return nil
	}

	n := 0
	var (
		sum *db.Summary
		err error
		rec *db.Record
	)
	for sum == nil && err == nil {
		rec, sum, err = b.receiveNext()
		if rec != nil {
			b.currStream.push(rec)
			n++
		}
	}

	if n > 0 {
		b.log.Warnf(log.Bolt3, b.logId, "Buffered %d records", n)
	}

	return err
}

func (b *bolt3) run(cypher string, params map[string]interface{}, tx *internalTx3) (*stream, error) {
	// If already streaming, finish current stream first
	if err := b.bufferStream(); err != nil {
		return nil, err
	}

	if err := b.assertState(bolt3_tx, bolt3_ready, bolt3_pendingtx); err != nil {
		return nil, err
	}

	var meta map[string]interface{}
	if tx != nil {
		meta = tx.toMeta()
	}

	// Append lazy begin transaction message
	if b.state == bolt3_pendingtx {
		b.out.appendBegin(meta)
		meta = nil
	}

	// Append run message
	b.out.appendRun(cypher, params, meta)

	// Append pull all message and send it along with other pending messages
	b.out.appendPullAll()
	if b.out.send(b.conn); b.err != nil {
		return nil, b.err
	}

	// Process server responses
	// Receive confirmation of transaction begin if it was started above
	if b.state == bolt3_pendingtx {
		if b.receiveSuccess(); b.err != nil {
			return nil, b.err
		}
		b.state = bolt3_tx
	}

	// Receive confirmation of run message
	succ := b.receiveSuccess()
	if b.err != nil {
		return nil, b.err
	}
	b.tfirst = succ.tfirst
	// Change state to streaming
	if b.state == bolt3_ready {
		b.state = bolt3_streaming
	} else {
		b.state = bolt3_streamingtx
	}

	b.currStream = &stream{keys: succ.fields}
	return b.currStream, nil
}

func (b *bolt3) Run(runCommand db.Command, txConfig db.TxConfig) (db.StreamHandle, error) {
	if err := b.assertState(bolt3_streaming, bolt3_ready); err != nil {
		return nil, err
	}

	tx := internalTx3{
		mode:      txConfig.Mode,
		bookmarks: txConfig.Bookmarks,
		timeout:   txConfig.Timeout,
		txMeta:    txConfig.Meta,
	}
	stream, err := b.run(runCommand.Cypher, runCommand.Params, &tx)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (b *bolt3) RunTx(txh db.TxHandle, runCommand db.Command) (db.StreamHandle, error) {
	if err := b.assertTxHandle(b.txId, txh); err != nil {
		return nil, err
	}

	stream, err := b.run(runCommand.Cypher, runCommand.Params, b.pendingTx)
	b.pendingTx = nil
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (b *bolt3) Keys(streamHandle db.StreamHandle) ([]string, error) {
	stream, ok := streamHandle.(*stream)
	if !ok {
		return nil, errors.New("Invalid stream handle")
	}
	// Don't care about if the stream is the current or even if it belongs to this connection.
	return stream.keys, nil
}

// Reads one record from the stream.
func (b *bolt3) Next(streamHandle db.StreamHandle) (*db.Record, *db.Summary, error) {
	stream, ok := streamHandle.(*stream)
	if !ok {
		return nil, nil, errors.New("Invalid stream handle")
	}

	// Buffered stream or someone elses stream, doesn't matter...
	buf, rec, sum, err := stream.bufferedNext()
	if buf {
		return rec, sum, err
	}

	// Nothing in the stream buffer, the stream must be the current
	// one to fetch on it otherwise something is wrong.
	if stream != b.currStream {
		return nil, nil, errors.New("Invalid stream handle")
	}

	return b.receiveNext()
}

func (b *bolt3) Consume(streamHandle db.StreamHandle) (*db.Summary, error) {
	stream, ok := streamHandle.(*stream)
	if !ok {
		return nil, errors.New("Invalid stream handle")
	}

	// If the stream isn't current, it should either already be complete
	// or have an error.
	if stream != b.currStream {
		return stream.sum, stream.err
	}

	// It is the current stream, it should not be complete but...
	if stream.err != nil || stream.sum != nil {
		return stream.sum, stream.err
	}

	b.discardStream()
	return stream.sum, stream.err
}

func (b *bolt3) Buffer(streamHandle db.StreamHandle) error {
	stream, ok := streamHandle.(*stream)
	if !ok {
		return errors.New("Invalid stream handle")
	}

	// If the stream isn't current, it should either already be complete
	// or have an error.
	if stream != b.currStream {
		return stream.Err()
	}

	// It is the current stream, it should not be complete but...
	if stream.err != nil || stream.sum != nil {
		return stream.Err()
	}

	b.bufferStream()
	return stream.Err()
}

// Reads one record from the network.
func (b *bolt3) receiveNext() (*db.Record, *db.Summary, error) {
	if err := b.assertState(bolt3_streaming, bolt3_streamingtx); err != nil {
		return nil, nil, err
	}

	res := b.receiveMsg()
	if b.err != nil {
		return nil, nil, b.err
	}

	switch x := res.(type) {
	case *db.Record:
		x.Keys = b.currStream.keys
		return x, nil, nil
	case *success:
		// End of stream, parse summary
		sum := x.summary()
		if sum == nil {
			b.state = bolt3_dead
			b.err = errors.New("Failed to parse summary")
			b.currStream.err = b.err
			b.currStream = nil
			b.log.Error(log.Bolt3, b.logId, b.err)
			return nil, nil, b.err
		}
		if b.state == bolt3_streamingtx {
			b.state = bolt3_tx
		} else {
			b.state = bolt3_ready
			// Keep bookmark for auto-commit tx
			if len(sum.Bookmark) > 0 {
				b.bookmark = sum.Bookmark
			}
		}
		b.currStream.sum = sum
		b.currStream = nil
		// Add some extras to the summary
		sum.ServerVersion = b.serverVersion
		sum.ServerName = b.serverName
		sum.TFirst = b.tfirst
		return nil, sum, nil
	case *db.Neo4jError:
		b.err = x
		b.currStream.err = b.err
		b.currStream = nil
		b.state = bolt3_failed
		if x.Classification() == "ClientError" {
			// These could include potentially large cypher statement, only log to debug
			b.log.Debugf(log.Bolt3, b.logId, "%s", x)
		} else {
			b.log.Error(log.Bolt3, b.logId, x)
		}
		return nil, nil, x
	default:
		b.state = bolt3_dead
		b.err = errors.New("Unknown response")
		b.currStream.err = b.err
		b.currStream = nil
		b.log.Error(log.Bolt3, b.logId, b.err)
		return nil, nil, b.err
	}
}

func (b *bolt3) Bookmark() string {
	return b.bookmark
}

func (b *bolt3) IsAlive() bool {
	return b.state != bolt3_dead
}

func (b *bolt3) Birthdate() time.Time {
	return b.birthDate
}

func (b *bolt3) Reset() {
	defer func() {
		// Reset internal state
		b.txId = 0
		b.currStream = nil
		b.bookmark = ""
		b.pendingTx = nil
		b.err = nil
	}()

	if b.state == bolt3_ready || b.state == bolt3_dead {
		// No need for reset
		return
	}

	// Discard any pending stream
	b.discardStream()

	if b.state == bolt3_ready || b.state == bolt3_dead {
		// No need for reset
		return
	}

	// Send the reset message to the server
	// Need to clear any pending error
	b.err = nil
	b.out.appendReset()
	if b.out.send(b.conn); b.err != nil {
		return
	}

	// Should receive x number of ignores until we get a success
	for {
		msg := b.receiveMsg()
		if b.err != nil {
			return
		}
		switch msg.(type) {
		case *ignored:
			// Command ignored
		case *success:
			// Reset confirmed
			b.state = bolt3_ready
			return
		default:
			b.state = bolt3_dead
			return
		}
	}
}

func (b *bolt3) GetRoutingTable(database string, context map[string]string) (*db.RoutingTable, error) {
	if err := b.assertState(bolt3_ready); err != nil {
		return nil, err
	}

	if database != db.DefaultDatabase {
		return nil, errors.New("Bolt 3 does not support routing to a specifiec database name")
	}

	// Only available when Neo4j is setup with clustering
	runCommand := db.Command{
		Cypher: "CALL dbms.cluster.routing.getRoutingTable($context)",
		Params: map[string]interface{}{"context": context},
	}
	txConfig := db.TxConfig{Mode: db.ReadMode}
	streamHandle, err := b.Run(runCommand, txConfig)
	if err != nil {
		// Give a better error
		dbError, isDbError := err.(*db.Neo4jError)
		if isDbError && dbError.Code == "Neo.ClientError.Procedure.ProcedureNotFound" {
			return nil, &db.RoutingNotSupportedError{Server: b.serverName}
		}
		return nil, err
	}

	rec, _, err := b.Next(streamHandle)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, errors.New("No routing table record")
	}
	// Just empty the stream, ignore the summary should leave the connecion in ready state
	b.Next(streamHandle)

	table := parseRoutingTableRecord(rec)
	if table == nil {
		return nil, errors.New("Unable to parse routing table")
	}

	return table, nil
}

// Beware, could be called on another thread when driver is closed.
func (b *bolt3) Close() {
	b.log.Infof(log.Bolt3, b.logId, "Close")
	if b.state != bolt3_dead {
		b.out.appendGoodbye()
		b.out.send(b.conn)
	}
	b.conn.Close()
	b.state = bolt3_dead
}
