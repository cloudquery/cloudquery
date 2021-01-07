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
	"container/list"
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type stream struct {
	keys      []string
	fifo      list.List
	sum       *db.Summary
	err       error
	qid       int64
	fetchSize int
	key       int64
}

// Acts on buffered data, first return value indicates if buffering
// is active or not.
func (s *stream) bufferedNext() (bool, *db.Record, *db.Summary, error) {
	e := s.fifo.Front()
	if e != nil {
		s.fifo.Remove(e)
		return true, e.Value.(*db.Record), nil, nil
	}
	if s.err != nil {
		return true, nil, nil, s.err
	}
	if s.sum != nil {
		return true, nil, s.sum, nil
	}

	return false, nil, nil, nil
}

// Delayed error until fifo emptied
func (s *stream) Err() error {
	if s.fifo.Len() > 0 {
		return nil
	}
	return s.err
}

func (s *stream) push(rec *db.Record) {
	s.fifo.PushBack(rec)
}

type openstreams struct {
	curr         *stream
	num          int
	key          int64
	onAssertFail func(e error)
	onClose      func(s *stream) // Stream succesfully closed
	onEmpty      func()          // All streams closed succesfully or with an error
}

var (
	invalidStream          = errors.New("Invalid stream handle")
	assertDetachUncomplete = errors.New("Assert fail: Detaching incomplete stream")
	assertShoudBeNoCurrent = errors.New("Assert fail: Should be no curr stream")
	assertShoudBeCurrent   = errors.New("Assert fail: Should be curr stream")
)

// Adds a new open stream and sets it as current.
// There should NOT be a current stream .
func (o *openstreams) attach(s *stream) {
	if o.curr != nil {
		o.onAssertFail(assertShoudBeNoCurrent)
		return
	}
	// Track number of open streams and set the stream as current
	o.num++
	o.curr = s
	s.key = o.key
}

// Detaches the current stream from being current and
// removes it from set of open streams it is no longer open.
// The stream should be either in failed state or completed.
func (o *openstreams) detach(sum *db.Summary, err error) {
	if o.curr == nil {
		o.onAssertFail(assertShoudBeCurrent)
		return
	}

	if sum != nil {
		o.curr.sum = sum
		o.onClose(o.curr)
	} else if err != nil {
		o.curr.err = err
	} else {
		o.onAssertFail(assertDetachUncomplete)
		return
	}

	o.remove(o.curr)
	o.curr = nil
	if o.num <= 0 {
		o.onEmpty()
	}
}

// Streams can be paused when they have received a 'has_more' response from server
// Pauses the current stream
func (o *openstreams) pause() {
	o.curr = nil
}

// When resuming a stream a new PULL message needs to be sent.
func (o *openstreams) resume(s *stream) {
	if o.curr != nil {
		o.onAssertFail(assertShoudBeNoCurrent)
		return
	}
	o.curr = s
}

// Internal, "removes" the stream by setting it's corresponding entry to nil.
func (o *openstreams) remove(s *stream) {
	o.num--
	s.key = 0
}

func (o *openstreams) reset() {
	num := o.num
	o.num = 0
	if num > 0 {
		o.onEmpty()
	}
	o.curr = nil
	o.key = time.Now().UnixNano()
}

// Checks that the handle represents a stream but not necessarily a stream belonging
// to this set of open streams.
func (o *openstreams) getUnsafe(h db.StreamHandle) (*stream, error) {
	stream, ok := h.(*stream)
	if !ok || stream == nil {
		return nil, invalidStream
	}
	return stream, nil
}

func (o *openstreams) isSafe(s *stream) error {
	if s.key == o.key {
		return nil
	}
	return invalidStream
}
