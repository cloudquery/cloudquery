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

package db

import (
	"fmt"
	"reflect"
	"strings"
)

// Database server failed to fullfill request.
type Neo4jError struct {
	Code           string
	Msg            string
	parsed         bool
	classification string
	category       string
	title          string
}

func (e *Neo4jError) Error() string {
	return fmt.Sprintf("Neo4jError: %s (%s)", e.Code, e.Msg)
}

func (e *Neo4jError) Classification() string {
	e.parse()
	return e.classification
}

func (e *Neo4jError) Category() string {
	e.parse()
	return e.category
}

func (e *Neo4jError) Title() string {
	e.parse()
	return e.title
}

// parse parses code from Neo4j into usable parts.
// Code Neo.ClientError.General.ForbiddenReadOnlyDatabase is split into:
//   Classification: ClientError
//   Category: General
//   Title: ForbiddernReadOnlyDatabase
func (e *Neo4jError) parse() {
	if e.parsed {
		return
	}
	e.parsed = true
	parts := strings.Split(e.Code, ".")
	if len(parts) != 4 {
		return
	}
	e.classification = parts[1]
	e.category = parts[2]
	e.title = parts[3]
}

func (e *Neo4jError) IsAuthenticationFailed() bool {
	return e.Code == "Neo.ClientError.Security.Unauthorized"
}

func (e *Neo4jError) IsRetriableTransient() bool {
	e.parse()
	if e.classification != "TransientError" {
		return false
	}
	switch e.Code {
	// Happens when client aborts transaction, should not retry
	case "Neo.TransientError.Transaction.Terminated", "Neo.TransientError.Transaction.LockClientStopped":
		return false
	}
	return true
}

func (e *Neo4jError) IsRetriableCluster() bool {
	switch e.Code {
	case "Neo.ClientError.Cluster.NotALeader", "Neo.ClientError.General.ForbiddenOnReadOnlyDatabase":
		return true
	}
	return false
}

type RoutingNotSupportedError struct {
	Server string
}

func (e *RoutingNotSupportedError) Error() string {
	return fmt.Sprintf("%s does not support routing", e.Server)
}

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("Usage of type '%s' is not supported", e.Type.String())
}
