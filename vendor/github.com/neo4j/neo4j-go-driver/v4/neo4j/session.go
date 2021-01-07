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
	"context"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/internal/retry"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

// TransactionWork represents a unit of work that will be executed against the provided
// transaction
type TransactionWork func(tx Transaction) (interface{}, error)

// Session represents a logical connection (which is not tied to a physical connection)
// to the server
type Session interface {
	// LastBookmark returns the bookmark received following the last successfully completed transaction.
	// If no bookmark was received or if this transaction was rolled back, the bookmark value will not be changed.
	LastBookmark() string
	// BeginTransaction starts a new explicit transaction on this session
	BeginTransaction(configurers ...func(*TransactionConfig)) (Transaction, error)
	// ReadTransaction executes the given unit of work in a AccessModeRead transaction with
	// retry logic in place
	ReadTransaction(work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error)
	// WriteTransaction executes the given unit of work in a AccessModeWrite transaction with
	// retry logic in place
	WriteTransaction(work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error)
	// Run executes an auto-commit statement and returns a result
	Run(cypher string, params map[string]interface{}, configurers ...func(*TransactionConfig)) (Result, error)
	// Close closes any open resources and marks this session as unusable
	Close() error
}

// SessionConfig is used to configure a new session, its zero value uses safe defaults.
type SessionConfig struct {
	// AccessMode used when using Session.Run and explicit transactions. Used to route query to
	// to read or write servers when running in a cluster. Session.ReadTransaction and Session.WriteTransaction
	// does not rely on this mode.
	AccessMode AccessMode
	// Bookmarks are the initial bookmarks used to ensure that the executing server is at least up
	// to date to the point represented by the latest of the provided bookmarks. After running commands
	// on the session the bookmark can be retrieved with Session.LastBookmark. All commands executing
	// within the same session will automatically use the bookmark from the previous command in the
	// session.
	Bookmarks []string
	// DatabaseName contains the name of the database that the commands in the session will execute on.
	DatabaseName string
	// FetchSize defines how many records to pull from server in each batch.
	// From Bolt protocol v4 (Neo4j 4+) records can be fetched in batches as compared to fetching
	// all in previous versions.
	//
	// If FetchSize is set to FetchDefault, the driver decides the appropriate size. If set to a positive value
	// that size is used if the underlying protocol supports it otherwise it is ignored.
	//
	// To turn off fetching in batches and always fetch everything, set FetchSize to FetchAll.
	// If a single large result is to be retrieved this is the most performant setting.
	FetchSize int
}

// Turns off fetching records in batches.
const FetchAll = -1

// Lets the driver decide fetch size
const FetchDefault = 0

// Connection pool as seen by the session.
type sessionPool interface {
	Borrow(ctx context.Context, serverNames []string, wait bool) (db.Connection, error)
	Return(c db.Connection)
	CleanUp()
}

type session struct {
	config       *Config
	defaultMode  db.AccessMode
	bookmarks    []string
	databaseName string
	pool         sessionPool
	router       sessionRouter
	txExplicit   *transaction
	txAuto       *autoTransaction
	sleep        func(d time.Duration)
	now          func() time.Time
	logId        string
	log          log.Logger
	throttleTime time.Duration
	fetchSize    int
}

// Remove empty string bookmarks to check for "bad" callers
// To avoid allocating, first check if this is a problem
func cleanupBookmarks(bookmarks []string) []string {
	hasBad := false
	for _, b := range bookmarks {
		if len(b) == 0 {
			hasBad = true
			break
		}
	}

	if !hasBad {
		return bookmarks
	}

	cleaned := make([]string, 0, len(bookmarks)-1)
	for _, b := range bookmarks {
		if len(b) > 0 {
			cleaned = append(cleaned, b)
		}
	}
	return cleaned
}

func newSession(config *Config, router sessionRouter, pool sessionPool,
	mode db.AccessMode, bookmarks []string, databaseName string, fetchSize int, logger log.Logger) *session {

	logId := log.NewId()
	logger.Debugf(log.Session, logId, "Created")

	return &session{
		config:       config,
		router:       router,
		pool:         pool,
		defaultMode:  mode,
		bookmarks:    cleanupBookmarks(bookmarks),
		databaseName: databaseName,
		sleep:        time.Sleep,
		now:          time.Now,
		log:          logger,
		logId:        logId,
		throttleTime: time.Second * 1,
		fetchSize:    fetchSize,
	}
}

func (s *session) LastBookmark() string {
	// Pick up bookmark from pending auto-commit if there is a bookmark on it
	if s.txAuto != nil {
		s.retrieveBookmarks(s.txAuto.conn)
	}

	// Report bookmark from previously closed connection or from initial set
	if len(s.bookmarks) > 0 {
		return s.bookmarks[len(s.bookmarks)-1]
	}

	return ""
}

func (s *session) BeginTransaction(configurers ...func(*TransactionConfig)) (Transaction, error) {
	// Guard for more than one transaction per session
	if s.txExplicit != nil {
		err := &UsageError{Message: "Session already has a pending transaction"}
		s.log.Error(log.Session, s.logId, err)
		return nil, err
	}

	if s.txAuto != nil {
		s.txAuto.done()
	}

	// Apply configuration functions
	config := TransactionConfig{Timeout: 0, Metadata: nil}
	for _, c := range configurers {
		c(&config)
	}

	// Get a connection from the pool. This could fail in clustered environment.
	conn, err := s.getConnection(s.defaultMode)
	if err != nil {
		return nil, err
	}

	// Begin transaction
	txHandle, err := conn.TxBegin(db.TxConfig{
		Mode:      s.defaultMode,
		Bookmarks: s.bookmarks,
		Timeout:   config.Timeout,
		Meta:      config.Metadata,
	})
	if err != nil {
		s.pool.Return(conn)
		return nil, err
	}

	// Create transaction wrapper
	s.txExplicit = &transaction{
		conn:      conn,
		fetchSize: s.fetchSize,
		txHandle:  txHandle,
		onClosed: func() {
			// On transaction closed (rollbacked or committed)
			s.retrieveBookmarks(conn)
			s.pool.Return(conn)
			s.txExplicit = nil
		},
	}

	return s.txExplicit, nil
}

func (s *session) runRetriable(
	mode db.AccessMode,
	work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error) {

	// Guard for more than one transaction per session
	if s.txExplicit != nil {
		err := &UsageError{Message: "Session already has a pending transaction"}
		return nil, err
	}

	if s.txAuto != nil {
		s.txAuto.done()
	}

	config := TransactionConfig{Timeout: 0, Metadata: nil}
	for _, c := range configurers {
		c(&config)
	}

	state := retry.State{
		MaxTransactionRetryTime: s.config.MaxTransactionRetryTime,
		Log:                     s.log,
		LogName:                 log.Session,
		LogId:                   s.logId,
		Now:                     s.now,
		Sleep:                   s.sleep,
		Throttle:                retry.Throttler(s.throttleTime),
		MaxDeadConnections:      s.config.MaxConnectionPoolSize,
		Router:                  s.router,
		DatabaseName:            s.databaseName,
	}
	for state.Continue() {
		// Establish new connection
		conn, err := s.getConnection(mode)
		if err != nil {
			state.OnFailure(conn, err, false)
			continue
		}

		// Begin transaction
		txHandle, err := conn.TxBegin(db.TxConfig{
			Mode:      mode,
			Bookmarks: s.bookmarks,
			Timeout:   config.Timeout,
			Meta:      config.Metadata,
		})
		if err != nil {
			state.OnFailure(conn, err, false)
			s.pool.Return(conn)
			continue
		}

		// Construct a transaction like thing for client to execute stuff on.
		// Evaluate the returned error from all the work for retryable, this means
		// that client can mess up the error handling.
		tx := retryableTransaction{conn: conn, fetchSize: s.fetchSize, txHandle: txHandle}
		x, err := work(&tx)
		if err != nil {
			state.OnFailure(conn, err, false)
			s.pool.Return(conn)
			continue
		}

		// Commit transaction
		err = conn.TxCommit(txHandle)
		if err != nil {
			state.OnFailure(conn, err, true)
			s.pool.Return(conn)
			continue
		}

		// Collect bookmark and return connection to pool
		s.retrieveBookmarks(conn)
		s.pool.Return(conn)

		// All well
		return x, nil
	}

	// When retries has occured wrap the error, the last error is always added but
	// cause is only set when the retry logic could detect something strange.
	if state.LastErrWasRetryable {
		err := newTransactionExecutionLimit(state.Errs, state.Causes)
		s.log.Error(log.Session, s.logId, err)
		return nil, err
	}
	// Wrap and log the error if it belongs to the driver
	err := wrapError(state.LastErr)
	switch err.(type) {
	case *UsageError, *ConnectivityError:
		s.log.Error(log.Session, s.logId, err)
	}
	return nil, err
}

func (s *session) ReadTransaction(
	work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error) {

	return s.runRetriable(db.ReadMode, work, configurers...)
}

func (s *session) WriteTransaction(
	work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error) {

	return s.runRetriable(db.WriteMode, work, configurers...)
}

func (s *session) getServers(mode db.AccessMode) ([]string, error) {
	if mode == db.ReadMode {
		return s.router.Readers(s.databaseName)
	} else {
		return s.router.Writers(s.databaseName)
	}
}

func (s *session) getConnection(mode db.AccessMode) (db.Connection, error) {
	servers, err := s.getServers(mode)
	if err != nil {
		return nil, wrapError(err)
	}

	var ctx context.Context
	if s.config.ConnectionAcquisitionTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), s.config.ConnectionAcquisitionTimeout)
		if cancel != nil {
			defer cancel()
		}
	} else {
		ctx = context.Background()
	}
	conn, err := s.pool.Borrow(ctx, servers, s.config.ConnectionAcquisitionTimeout != 0)
	if err != nil {
		return nil, wrapError(err)
	}

	// Select database on server
	if s.databaseName != db.DefaultDatabase {
		dbSelector, ok := conn.(db.DatabaseSelector)
		if !ok {
			s.pool.Return(conn)
			return nil, &UsageError{Message: "Database does not support multidatabase"}
		}
		dbSelector.SelectDatabase(s.databaseName)
	}

	return conn, nil
}

func (s *session) retrieveBookmarks(conn db.Connection) {
	if conn == nil {
		return
	}
	bookmark := conn.Bookmark()
	if len(bookmark) > 0 {
		s.bookmarks = []string{bookmark}
	}
}

func (s *session) Run(
	cypher string, params map[string]interface{}, configurers ...func(*TransactionConfig)) (Result, error) {

	if s.txExplicit != nil {
		err := &UsageError{Message: "Trying to run auto-commit transaction while in explicit transaction"}
		s.log.Error(log.Session, s.logId, err)
		return nil, err
	}

	if s.txAuto != nil {
		s.txAuto.done()
	}

	config := TransactionConfig{Timeout: 0, Metadata: nil}
	for _, c := range configurers {
		c(&config)
	}

	conn, err := s.getConnection(s.defaultMode)
	if err != nil {
		return nil, err
	}

	stream, err := conn.Run(
		db.Command{
			Cypher:    cypher,
			Params:    params,
			FetchSize: s.fetchSize,
		},
		db.TxConfig{
			Mode:      s.defaultMode,
			Bookmarks: s.bookmarks,
			Timeout:   config.Timeout,
			Meta:      config.Metadata,
		})
	if err != nil {
		s.pool.Return(conn)
		return nil, wrapError(err)
	}

	s.txAuto = &autoTransaction{
		conn: conn,
		res:  newResult(conn, stream, cypher, params),
		onClosed: func() {
			s.retrieveBookmarks(conn)
			s.pool.Return(conn)
			s.txAuto = nil
		},
	}

	return s.txAuto.res, nil
}

func (s *session) Close() error {
	var err error

	if s.txExplicit != nil {
		s.txExplicit.Close()
		err = &UsageError{Message: "Closing session with a pending transaction"}
		s.log.Warnf(log.Session, s.logId, err.Error())
	}

	if s.txAuto != nil {
		s.txAuto.discard()
	}

	s.log.Debugf(log.Session, s.logId, "Closed")

	// Schedule cleanups
	go func() {
		s.pool.CleanUp()
		s.router.CleanUp()
	}()
	return err
}

type sessionWithError struct {
	err error
}

func (s *sessionWithError) LastBookmark() string {
	return ""
}

func (s *sessionWithError) BeginTransaction(configurers ...func(*TransactionConfig)) (Transaction, error) {
	return nil, s.err
}
func (s *sessionWithError) ReadTransaction(work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error) {
	return nil, s.err
}
func (s *sessionWithError) WriteTransaction(work TransactionWork, configurers ...func(*TransactionConfig)) (interface{}, error) {
	return nil, s.err
}
func (s *sessionWithError) Run(cypher string, params map[string]interface{}, configurers ...func(*TransactionConfig)) (Result, error) {
	return nil, s.err
}
func (s *sessionWithError) Close() error {
	return s.err
}
