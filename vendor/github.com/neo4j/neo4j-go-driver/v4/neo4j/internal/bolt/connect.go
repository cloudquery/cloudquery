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

// Package bolt contains implementations of the database functionality.
package bolt

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

type protocolVersion struct {
	major byte
	minor byte
}

// Supported versions in priority order
var versions = [4]protocolVersion{
	{major: 4, minor: 2},
	{major: 4, minor: 1},
	{major: 4, minor: 0},
	{major: 3, minor: 0},
}

// Negotiate version of bolt protocol.
// Returns instance of bolt protocol implmenting low-level abstract connection Connection interface.
func Connect(serverName string, conn net.Conn, auth map[string]interface{}, userAgent string, routingContext map[string]string, log log.Logger) (db.Connection, error) {
	// Perform Bolt handshake to negotiate version
	// Send handshake to server
	handshake := []byte{
		0x60, 0x60, 0xb0, 0x17, // Magic: GoGoBolt
		0x00, 0x00, versions[0].minor, versions[0].major,
		0x00, 0x00, versions[1].minor, versions[1].major,
		0x00, 0x00, versions[2].minor, versions[2].major,
		0x00, 0x00, versions[3].minor, versions[3].major,
	}
	_, err := conn.Write(handshake)
	if err != nil {
		return nil, err
	}

	// Receive accepted server version
	buf := make([]byte, 4)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	// Parse received version and construct the correct instance
	major := buf[3]
	minor := buf[2]
	switch major {
	case 3:
		// Handover rest of connection handshaking
		boltConn := NewBolt3(serverName, conn, log)
		err = boltConn.connect(auth, userAgent)
		if err != nil {
			return nil, err
		}
		return boltConn, nil
	case 4:
		// Handover rest of connection handshaking
		boltConn := NewBolt4(serverName, conn, log)
		err = boltConn.connect(int(minor), auth, userAgent, routingContext)
		if err != nil {
			return nil, err
		}
		return boltConn, nil
	case 0:
		err = errors.New(fmt.Sprintf("Server did not accept any of the requested Bolt versions (%#v)", versions))
	default:
		err = errors.New(fmt.Sprintf("Server responded with unsupported version %d.%d", major, minor))
	}

	return nil, err
}
