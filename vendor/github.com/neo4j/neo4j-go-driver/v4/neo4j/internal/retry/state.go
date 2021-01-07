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

// Package retry handles retry operations.
package retry

import (
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

type Router interface {
	Invalidate(database string)
}

type CommitFailedDeadError struct {
	inner error
}

func (e *CommitFailedDeadError) Error() string {
	return fmt.Sprintf("Connection lost during commit: %s", e.inner)
}

type State struct {
	LastErrWasRetryable     bool
	LastErr                 error
	stop                    bool
	Errs                    []error
	Causes                  []string
	MaxTransactionRetryTime time.Duration
	Log                     log.Logger
	LogName                 string
	LogId                   string
	Now                     func() time.Time
	Sleep                   func(time.Duration)
	Throttle                Throttler
	MaxDeadConnections      int
	Router                  Router
	DatabaseName            string

	start      time.Time
	cause      string
	deadErrors int
	skipSleep  bool
}

func (s *State) OnFailure(conn db.Connection, err error, isCommitting bool) {
	s.LastErr = err
	s.cause = ""
	s.skipSleep = false

	// Check timeout
	if s.start.IsZero() {
		s.start = s.Now()
	}
	if s.Now().Sub(s.start) > s.MaxTransactionRetryTime {
		s.stop = true
		s.cause = "Timeout"
		return
	}

	// Reset after determined to evaluate this error
	s.LastErrWasRetryable = false

	// Failed to connect
	if conn == nil {
		s.LastErrWasRetryable = true
		s.cause = "No available connection"
		return
	}

	// Check if the connection died, if it died during commit it is not safe to retry.
	if !conn.IsAlive() {
		if isCommitting {
			s.stop = true
			// The error is most probably io.EOF so enrich the error
			// to make this error more recognizable.
			s.LastErr = &CommitFailedDeadError{inner: s.LastErr}
			return
		}

		s.deadErrors += 1
		s.stop = s.deadErrors > s.MaxDeadConnections
		s.LastErrWasRetryable = true
		s.cause = "Connection lost"
		s.skipSleep = true
		return
	}

	if dbErr, isDbErr := err.(*db.Neo4jError); isDbErr {
		if dbErr.IsRetriableCluster() {
			// Force routing tables to be updated before trying again
			s.Router.Invalidate(s.DatabaseName)
			s.cause = "Cluster error"
			s.LastErrWasRetryable = true
			return
		}

		if dbErr.IsRetriableTransient() {
			s.cause = "Transient error"
			s.LastErrWasRetryable = true
			return
		}
	}

	s.stop = true
}

func (s *State) Continue() bool {
	// No error happened yet
	if !s.stop && s.LastErr == nil {
		return true
	}

	// Track the error and the cause
	s.Errs = append(s.Errs, s.LastErr)
	if s.cause != "" {
		s.Causes = append(s.Causes, s.cause)
	}

	// Retry after optional sleep
	if !s.stop {
		if s.skipSleep {
			s.Log.Debugf(s.LogName, s.LogId, "Retrying transaction (%s): %s", s.cause, s.LastErr)
		} else {
			s.Throttle = s.Throttle.next()
			sleepTime := s.Throttle.delay()
			s.Log.Debugf(s.LogName, s.LogId,
				"Retrying transaction (%s): %s [after %s]", s.cause, s.LastErr, sleepTime)
			s.Sleep(sleepTime)
		}
		return true
	}

	return false
}
