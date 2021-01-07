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

// Package db defines generic database functionality.
package db

import (
	"time"
)

// Definitions of these should correspond to public API
type AccessMode int

const (
	WriteMode AccessMode = 0
	ReadMode  AccessMode = 1
)

type (
	TxHandle     uint64
	StreamHandle interface{}
)

type Command struct {
	Cypher    string
	Params    map[string]interface{}
	FetchSize int
}

type TxConfig struct {
	Mode      AccessMode
	Bookmarks []string
	Timeout   time.Duration
	Meta      map[string]interface{}
}

// Connection defines an abstract database server connection.
type Connection interface {
	TxBegin(txConfig TxConfig) (TxHandle, error)
	TxRollback(tx TxHandle) error
	TxCommit(tx TxHandle) error
	Run(cmd Command, txConfig TxConfig) (StreamHandle, error)
	RunTx(tx TxHandle, cmd Command) (StreamHandle, error)
	// Keys for the specified stream.
	Keys(streamHandle StreamHandle) ([]string, error)
	// Moves to next item in the stream.
	// If error is nil, either Record or Summary has a value, if Record is nil there are no more records.
	// If error is non nil, neither Record or Summary has a value.
	Next(streamHandle StreamHandle) (*Record, *Summary, error)
	// Discards all records on the stream and returns the summary otherwise it will return the error.
	Consume(streamHandle StreamHandle) (*Summary, error)
	// Buffers all records on the stream, records, summary and error will be received through call to Next
	// The Connection implementation should preserve/buffer streams automatically if needed when new
	// streams are created and the server doesn't support multiple streams. Use Buffer to force
	// buffering before calling Reset to get all records and the bookmark.
	Buffer(streamHandle StreamHandle) error
	// Returns bookmark from last committed transaction or last finished auto-commit transaction.
	// Note that if there is an ongoing auto-commit transaction (stream active) the bookmark
	// from that is not included, use Buffer or Consume to end the stream with a bookmark.
	// Empty string if no bookmark.
	Bookmark() string
	// Returns name of the remote server
	ServerName() string
	// Returns server version on pattern Neo4j/1.2.3
	ServerVersion() string
	// Returns true if the connection is fully functional.
	// Implementation of this should be passive, no pinging or similair since it might be
	// called rather frequently.
	IsAlive() bool
	// Returns the point in time when this connection was established.
	Birthdate() time.Time
	// Resets connection to same state as directly after a connect.
	// Active streams will be discarded and the bookmark will be lost.
	Reset()
	// Closes the database connection as well as any underlying connection.
	// The instance should not be used after being closed.
	Close()
	// Gets routing table for specified database name or the default database if
	// database equals DefaultDatabase. If the underlying connection does not support
	// multiple databases, DefaultDatabase should be used as database.
	GetRoutingTable(database string, context map[string]string) (*RoutingTable, error)
}

type RoutingTable struct {
	TimeToLive int
	Routers    []string
	Readers    []string
	Writers    []string
}

// Marker for using the default database instance.
const DefaultDatabase = ""

// If database server connection supports selecting which database instance on the server
// to connect to. Prior to Neo4j 4 there was only one database per server.
type DatabaseSelector interface {
	// Should be called immediately after Reset. Not allowed to call multiple times with different
	// databases without a reset inbetween.
	SelectDatabase(database string)
}
