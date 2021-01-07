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

package pool

import (
	"container/list"
	"sync/atomic"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

// Represents a server with a number of connections that either is in use (borrowed) or
// is ready for use.
// Not thread safe
type server struct {
	idle            list.List
	busy            list.List
	failedConnectAt time.Time
	roundRobin      uint32
}

var sharedRoundRobin uint32

const rememberFailedConnectDuration = 3 * time.Minute

// Returns a idle connection if any
func (s *server) getIdle() db.Connection {
	// Remove from idle list and add to busy list
	e := s.idle.Front()
	if e != nil {
		c := s.idle.Remove(e)
		s.busy.PushFront(c)
		// Update round-robin counter every time we give away a connection and keep track
		// of our own round-robin index
		s.roundRobin = atomic.AddUint32(&sharedRoundRobin, 1)
		return c.(db.Connection)
	}
	return nil
}

func (s *server) notifyFailedConnect(now time.Time) {
	s.failedConnectAt = now
}

func (s *server) notifySuccesfulConnect() {
	s.failedConnectAt = time.Time{}
}

func (s *server) hasFailedConnect(now time.Time) bool {
	if s.failedConnectAt.IsZero() {
		return false
	}
	return now.Sub(s.failedConnectAt) < rememberFailedConnectDuration
}

const newConnectionPenalty = uint32(1 << 8)

// Calculates a penalty value for how this server compares to other servers
// when there is more than one server to choose from. The lower penalty the better choice.
func (s *server) calculatePenalty(now time.Time) uint32 {
	penalty := uint32(0)

	// If a connect to the server has failed recently, add a penalty
	if s.hasFailedConnect(now) {
		penalty = 1 << 31
	}
	// The more busy connections, the higher penalty
	numBusy := uint32(s.busy.Len())
	if numBusy > 0xff {
		numBusy = 0xff
	}
	penalty |= numBusy << 16
	// If there are no idle connections, add a penalty as the cost of connect would
	// add to the transaction time
	if s.idle.Len() == 0 {
		penalty |= newConnectionPenalty
	}
	// Use last round-robin value as lowest priority penalty, so when all other is equal we will
	// make sure to spread usage among the servers. And yes it will wrap around once in a while
	// but since number of busy servers weights higher it will even out pretty fast.
	penalty |= (s.roundRobin & 0xff)

	return penalty
}

// Returns a busy connection, makes it idle
func (s *server) returnBusy(c db.Connection) {
	s.unregisterBusy(c)
	s.idle.PushFront(c)
}

// Number of idle connections
func (s server) numIdle() int {
	return s.idle.Len()
}

// Adds a connection to busy list
func (s *server) registerBusy(c db.Connection) {
	// Update round-robin to indicate when this server was last used.
	s.roundRobin = atomic.AddUint32(&sharedRoundRobin, 1)
	s.busy.PushFront(c)
}

func (s *server) unregisterBusy(c db.Connection) {
	found := false
	for e := s.busy.Front(); e != nil && !found; e = e.Next() {
		x := e.Value.(db.Connection)
		found = x == c
		if found {
			s.busy.Remove(e)
			return
		}
	}
}

func (s *server) size() int {
	return s.busy.Len() + s.idle.Len()
}

func (s *server) removeIdleOlderThan(now time.Time, maxAge time.Duration) {
	e := s.idle.Front()
	for e != nil {
		n := e.Next()
		c := e.Value.(db.Connection)

		age := now.Sub(c.Birthdate())
		if age >= maxAge {
			s.idle.Remove(e)
			go c.Close()
		}

		e = n
	}
}

func closeAndEmptyConnections(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		c := e.Value.(db.Connection)
		c.Close()
	}
	l.Init()
}

func (s *server) closeAll() {
	closeAndEmptyConnections(s.idle)
	// Closing the busy connections could mean here that we do close from another thread.
	closeAndEmptyConnections(s.busy)
}
