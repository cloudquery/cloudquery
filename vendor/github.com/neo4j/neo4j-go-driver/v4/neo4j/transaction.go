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

package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

// Transaction represents a transaction in the Neo4j database
type Transaction interface {
	// Run executes a statement on this transaction and returns a result
	Run(cypher string, params map[string]interface{}) (Result, error)
	// Commit commits the transaction
	Commit() error
	// Rollback rolls back the transaction
	Rollback() error
	// Close rolls back the actual transaction if it's not already committed/rolled back
	// and closes all resources associated with this transaction
	Close() error
}

// Transaction implementation when explicit transaction started
type transaction struct {
	conn      db.Connection
	fetchSize int
	txHandle  db.TxHandle
	done      bool
	err       error
	onClosed  func()
}

func (tx *transaction) Run(cypher string, params map[string]interface{}) (Result, error) {
	stream, err := tx.conn.RunTx(tx.txHandle, db.Command{Cypher: cypher, Params: params, FetchSize: tx.fetchSize})
	if err != nil {
		return nil, wrapError(err)
	}
	return newResult(tx.conn, stream, cypher, params), nil
}

func (tx *transaction) Commit() error {
	if tx.done {
		return tx.err
	}
	tx.err = tx.conn.TxCommit(tx.txHandle)
	tx.done = true
	tx.onClosed()
	return wrapError(tx.err)
}

func (tx *transaction) Rollback() error {
	if tx.done {
		return tx.err
	}
	tx.err = tx.conn.TxRollback(tx.txHandle)
	tx.done = true
	tx.onClosed()
	return wrapError(tx.err)
}

func (tx *transaction) Close() error {
	return tx.Rollback()
}

// Transaction implementation used as parameter to transactional functions
type retryableTransaction struct {
	conn      db.Connection
	fetchSize int
	txHandle  db.TxHandle
}

func (tx *retryableTransaction) Run(cypher string, params map[string]interface{}) (Result, error) {
	stream, err := tx.conn.RunTx(tx.txHandle, db.Command{Cypher: cypher, Params: params, FetchSize: tx.fetchSize})
	if err != nil {
		return nil, wrapError(err)
	}
	return newResult(tx.conn, stream, cypher, params), nil
}

func (tx *retryableTransaction) Commit() error {
	return &UsageError{Message: "Commit not allowed on retryable transaction"}
}

func (tx *retryableTransaction) Rollback() error {
	return &UsageError{Message: "Rollback not allowed on retryable transaction"}
}

func (tx *retryableTransaction) Close() error {
	return &UsageError{Message: "Close not allowed on retryable transaction"}
}

// Represents an auto commit transaction.
// Does not implement the Transaction interface.
type autoTransaction struct {
	conn     db.Connection
	res      *result
	closed   bool
	onClosed func()
}

func (tx *autoTransaction) done() {
	if !tx.closed {
		tx.res.buffer()
		tx.closed = true
		tx.onClosed()
	}
}

func (tx *autoTransaction) discard() {
	if !tx.closed {
		tx.res.Consume()
		tx.closed = true
		tx.onClosed()
	}
}
