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

package router

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type ReadRoutingTableError struct {
	err    error
	server string
}

func (e *ReadRoutingTableError) Error() string {
	if e.err != nil || len(e.server) > 0 {
		return fmt.Sprintf("Unable to retrieve routing table from %s: %s", e.server, e.err)
	}
	return "Unable to retrieve routing table, no router provided"
}

func wrapError(server string, err error) error {
	// Preserve error originating from the database, wrap other errors
	_, isNeo4jErr := err.(*db.Neo4jError)
	if isNeo4jErr {
		return err
	}
	return &ReadRoutingTableError{server: server, err: err}
}
