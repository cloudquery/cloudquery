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
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bolt

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

// Intermediate representation of part of path
type relNode struct {
	id    int64
	name  string
	props map[string]interface{}
}

// buildPath builds a path from Bolt representation
func buildPath(nodes []dbtype.Node, relNodes []*relNode, indexes []int) dbtype.Path {
	num := len(indexes) / 2
	if num == 0 {
		return dbtype.Path{}
	}
	rels := make([]dbtype.Relationship, 0, num)

	i := 0
	n1 := nodes[0]
	for num > 0 {
		relni := indexes[i]
		i++
		n2i := indexes[i]
		i++
		num--
		var reln *relNode
		var n1start bool
		if relni < 0 {
			reln = relNodes[(relni*-1)-1]
		} else {
			reln = relNodes[relni-1]
			n1start = true
		}
		n2 := nodes[n2i]

		rel := dbtype.Relationship{
			Id:    reln.id,
			Type:  reln.name,
			Props: reln.props,
		}
		if n1start {
			rel.StartId = n1.Id
			rel.EndId = n2.Id
		} else {
			rel.StartId = n2.Id
			rel.EndId = n1.Id
		}
		rels = append(rels, rel)
		n1 = n2
	}

	return dbtype.Path{Nodes: nodes, Relationships: rels}
}
