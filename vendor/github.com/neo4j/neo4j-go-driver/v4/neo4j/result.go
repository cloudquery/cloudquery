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

type Result interface {
	// Keys returns the keys available on the result set.
	Keys() ([]string, error)
	// Next returns true only if there is a record to be processed.
	Next() bool
	// NextRecord returns true if there is a record to be processed, record parameter is set
	// to point to current record.
	NextRecord(record **Record) bool
	// Err returns the latest error that caused this Next to return false.
	Err() error
	// Record returns the current record.
	Record() *Record
	// Collect fetches all remaining records and returns them.
	Collect() ([]*Record, error)
	// Single returns one and only one record from the stream.
	// If the result stream contains zero or more than one records, error is returned.
	Single() (*Record, error)
	// Consume discards all remaining records and returns the summary information
	// about the statement execution.
	Consume() (ResultSummary, error)
}

type result struct {
	conn         db.Connection
	streamHandle db.StreamHandle
	cypher       string
	params       map[string]interface{}
	record       *Record
	summary      *db.Summary
	err          error
}

func newResult(conn db.Connection, str db.StreamHandle, cypher string, params map[string]interface{}) *result {
	return &result{
		conn:         conn,
		streamHandle: str,
		cypher:       cypher,
		params:       params,
	}
}

func (r *result) Keys() ([]string, error) {
	return r.conn.Keys(r.streamHandle)
}

func (r *result) Next() bool {
	r.record, r.summary, r.err = r.conn.Next(r.streamHandle)
	return r.record != nil
}

func (r *result) NextRecord(out **Record) bool {
	r.record, r.summary, r.err = r.conn.Next(r.streamHandle)
	if out != nil {
		*out = r.record
	}
	return r.record != nil
}

func (r *result) Record() *Record {
	return r.record
}

func (r *result) Err() error {
	return wrapError(r.err)
}

func (r *result) Collect() ([]*Record, error) {
	recs := make([]*Record, 0, 1024)
	for r.summary == nil && r.err == nil {
		r.record, r.summary, r.err = r.conn.Next(r.streamHandle)
		if r.record != nil {
			recs = append(recs, r.record)
		}
	}
	if r.err != nil {
		return nil, wrapError(r.err)
	}
	return recs, nil
}

func (r *result) buffer() {
	r.err = r.conn.Buffer(r.streamHandle)
}

func (r *result) Single() (*Record, error) {
	// Try retrieving the single record
	r.record, r.summary, r.err = r.conn.Next(r.streamHandle)
	if r.err != nil {
		return nil, wrapError(r.err)
	}
	if r.summary != nil {
		r.err = &UsageError{Message: "Result contains no more records"}
		return nil, r.err
	}

	// This is the potential single record
	single := r.record

	// Probe connection for more records
	r.record, r.summary, r.err = r.conn.Next(r.streamHandle)
	if r.record != nil {
		// There were more records, consume the stream since the user didn't
		// expect more records and should therefore not use them.
		r.summary, _ = r.conn.Consume(r.streamHandle)
		r.err = &UsageError{Message: "Result contains more than one record"}
		r.record = nil
		return nil, r.err
	}
	if r.err != nil {
		// Might be more records or not, anyway something is bad.
		// Both r.record and r.summary are nil at this point which is good.
		return nil, wrapError(r.err)
	}
	// We got the expected summary
	// r.record contains the single record and r.summary the summary.
	r.record = single
	return single, nil
}

func (r *result) toResultSummary() ResultSummary {
	return &resultSummary{
		sum:    r.summary,
		cypher: r.cypher,
		params: r.params,
	}
}

func (r *result) Consume() (ResultSummary, error) {
	// Already failed, reuse the internal error, might have been
	// set by Single to indicate some kind of usage error that "destroyed"
	// the result.
	if r.err != nil {
		return nil, wrapError(r.err)
	}

	r.record = nil
	r.summary, r.err = r.conn.Consume(r.streamHandle)
	if r.err != nil {
		return nil, wrapError(r.err)
	}
	return r.toResultSummary(), nil
}
