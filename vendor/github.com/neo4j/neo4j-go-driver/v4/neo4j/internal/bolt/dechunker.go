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
	"encoding/binary"
	"io"
)

func dechunkMessage(rd io.Reader, msgBuf []byte) ([]byte, error) {
	sizeBuf := []byte{0x00, 0x00}
	off := 0

	for {
		_, err := io.ReadFull(rd, sizeBuf)
		if err != nil {
			return msgBuf, err
		}
		chunkSize := int(binary.BigEndian.Uint16(sizeBuf))
		if chunkSize == 0 {
			if off > 0 {
				return msgBuf[:off], nil
			}
			// Got a nop chunk
			continue
		}

		// Need to expand buffer
		if (off + chunkSize) > cap(msgBuf) {
			msgBuf = append(msgBuf, make([]byte, off+chunkSize)...)
		}
		// Read the chunk into buffer
		_, err = io.ReadFull(rd, msgBuf[off:(off+chunkSize)])
		if err != nil {
			return msgBuf, err
		}
		off += chunkSize
	}
}
