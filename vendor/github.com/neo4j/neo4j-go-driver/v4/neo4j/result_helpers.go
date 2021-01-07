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

import "fmt"

// Single returns one and only one record from the result stream. Any error passed in
// or reported while navigating the result stream is returned without any conversion.
// If the result stream contains zero or more than one records error is returned.
//   record, err := neo4j.Single(session.Run(...))
func Single(result Result, err error) (*Record, error) {
	if err != nil {
		return nil, err
	}
	return result.Single()
}

// Collect loops through the result stream, collects records into a slice and returns the
// resulting slice. Any error passed in or reported while navigating the result stream is
// returned without any conversion.
//   records, err := neo4j.Collect(session.Run(...))
func Collect(result Result, err error) ([]*Record, error) {
	if err != nil {
		return nil, err
	}
	return result.Collect()
}

// AsRecords passes any existing error or casts from to a slice of records.
// Use in combination with Collect and transactional functions:
//   records, err := neo4j.AsRecords(session.ReadTransaction(func (tx neo4j.Transaction) {
//       return neo4j.Collect(tx.Run(...))
//   }))
func AsRecords(from interface{}, err error) ([]*Record, error) {
	if err != nil {
		return nil, err
	}
	recs, ok := from.([]*Record)
	if !ok {
		return nil, &UsageError{
			Message: fmt.Sprintf("Expected type []*Record, not %T", from),
		}
	}
	return recs, nil
}

// AsRecord passes any existing error or casts from to a record.
// Use in combination with Single and transactional functions:
//   record, err := neo4j.AsRecord(session.ReadTransaction(func (tx neo4j.Transaction) {
//       return neo4j.Single(tx.Run(...))
//   }))
func AsRecord(from interface{}, err error) (*Record, error) {
	if err != nil {
		return nil, err
	}
	rec, ok := from.(*Record)
	if !ok {
		return nil, &UsageError{
			Message: fmt.Sprintf("Expected type *Record, not %T", from),
		}
	}
	return rec, nil
}
