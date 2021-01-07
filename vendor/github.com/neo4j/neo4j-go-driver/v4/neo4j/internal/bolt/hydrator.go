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
	"errors"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/internal/packstream"
)

var hydrationInvalidState = errors.New("Hydration state error")

type ignored struct{}
type success struct {
	fields        []string
	tfirst        int64
	qid           int64
	bookmark      string
	connectionId  string
	server        string
	db            string
	hasMore       bool
	tlast         int64
	qtype         db.StatementType
	counters      map[string]int
	plan          *db.Plan
	profile       *db.ProfiledPlan
	notifications []db.Notification
	num           uint32
}

func (s *success) String() string {
	str := fmt.Sprintf("%#v", s)
	if s.plan != nil {
		str += fmt.Sprintf(" \nplan: %#v", s.plan)
	}
	if s.profile != nil {
		str += fmt.Sprintf(" \nprofile: %#v", s.profile)
	}
	return str
}

func (s *success) summary() *db.Summary {
	return &db.Summary{
		Bookmark:      s.bookmark,
		StmntType:     s.qtype,
		Counters:      s.counters,
		TLast:         s.tlast,
		Plan:          s.plan,
		ProfiledPlan:  s.profile,
		Notifications: s.notifications,
	}
}

func (s *success) isResetResponse() bool {
	return s.num == 0
}

// Since it is reused
func (s *success) clear() {
	s.fields = nil
	s.bookmark = ""
	s.hasMore = false
	s.notifications = nil
	s.plan = nil
	s.profile = nil
	s.qid = -1
	s.server = ""
	s.connectionId = ""
	s.server = ""
	s.tfirst = 0
	s.tlast = 0
	s.db = ""
	s.qtype = db.StatementTypeUnknown
}

type hydrator struct {
	unpacker      packstream.Unpacker
	unp           *packstream.Unpacker
	err           error
	cachedIgnored ignored
	cachedSuccess success
}

func (h *hydrator) setErr(err error) {
	if h.err != nil {
		h.err = err
	}
}

func (h *hydrator) getErr() error {
	if h.unp.Err != nil {
		return h.unp.Err
	}
	return h.err
}

func (h *hydrator) assertLength(expect, actual uint32) {
	if expect != actual {
		h.setErr(errors.New(fmt.Sprintf(
			"Invalid length of struct, expected %d but was %d", expect, actual)))
	}
}

// hydrate hydrates a top-level struct message
func (h *hydrator) hydrate(buf []byte) (x interface{}, err error) {
	h.unp = &h.unpacker
	h.unp.Reset(buf)
	h.unp.Next()

	if h.unp.Curr != packstream.PackedStruct {
		return nil, errors.New(fmt.Sprintf("Expected struct"))
	}

	n := h.unp.Len()
	t := h.unp.StructTag()
	switch t {
	case msgSuccess:
		x = h.success(n)
	case msgIgnored:
		x = h.ignored(n)
	case msgFailure:
		x = h.failure(n)
	case msgRecord:
		x = h.record(n)
	default:
		return nil, errors.New(fmt.Sprintf("Unexpected tag at top level: %d", t))
	}
	err = h.getErr()
	return
}

func (h *hydrator) ignored(n uint32) *ignored {
	h.assertLength(0, n)
	if h.getErr() != nil {
		return nil
	}
	return &h.cachedIgnored
}

func (h *hydrator) failure(n uint32) *db.Neo4jError {
	h.assertLength(1, n)
	if h.getErr() != nil {
		return nil
	}
	dberr := db.Neo4jError{}
	h.unp.Next() // Detect map
	for maplen := h.unp.Len(); maplen > 0; maplen-- {
		h.unp.Next()
		key := h.unp.String()
		h.unp.Next()
		switch key {
		case "code":
			dberr.Code = h.unp.String()
		case "message":
			dberr.Msg = h.unp.String()
		default:
			// Do not fail on unknown value in map
			h.trash()
		}
	}
	return &dberr
}

func (h *hydrator) success(n uint32) *success {
	h.assertLength(1, n)
	if h.getErr() != nil {
		return nil
	}
	// Use cached success but clear it first
	succ := &h.cachedSuccess
	succ.clear()

	h.unp.Next() // Detect map
	n = h.unp.Len()
	succ.num = n
	for ; n > 0; n-- {
		// Key
		h.unp.Next()
		key := h.unp.String()
		// Value
		h.unp.Next()
		switch key {
		case "fields":
			succ.fields = h.strings()
		case "t_first":
			succ.tfirst = h.unp.Int()
		case "qid":
			succ.qid = h.unp.Int()
		case "bookmark":
			succ.bookmark = h.unp.String()
		case "connection_id":
			succ.connectionId = h.unp.String()
		case "server":
			succ.server = h.unp.String()
		case "has_more":
			succ.hasMore = h.unp.Bool()
		case "t_last":
			succ.tlast = h.unp.Int()
		case "type":
			switch h.unp.String() {
			case "r":
				succ.qtype = db.StatementTypeRead
			case "w":
				succ.qtype = db.StatementTypeWrite
			case "rw":
				succ.qtype = db.StatementTypeReadWrite
			case "s":
				succ.qtype = db.StatementTypeSchemaWrite
			}
		case "db":
			succ.db = h.unp.String()
		case "stats":
			succ.counters = h.successStats()
		case "plan":
			m := h.amap()
			succ.plan = parsePlan(m)
		case "profile":
			m := h.amap()
			succ.profile = parseProfile(m)
		case "notifications":
			l := h.array()
			succ.notifications = parseNotifications(l)
		default:
			// Unknown key, waste it
			h.trash()
		}
	}
	return succ
}

func (h *hydrator) successStats() map[string]int {
	n := h.unp.Len()
	if n == 0 {
		return nil
	}
	counts := make(map[string]int, n)
	for ; n > 0; n-- {
		h.unp.Next()
		key := h.unp.String()
		h.unp.Next()
		val := h.unp.Int()
		counts[key] = int(val)
	}
	return counts
}

/*
func (h *hydrator) successPlanOperator() {
}

func (h *hydrator) successPlan() *db.Plan {
	// Should be on a map
	n := h.unp.Len()
	if n == 0 {
		return nil
	}

	plan := &db.Plan{}
}
*/

func (h *hydrator) strings() []string {
	n := h.unp.Len()
	slice := make([]string, n)
	for i := range slice {
		h.unp.Next()
		slice[i] = h.unp.String()
	}
	return slice
}

func (h *hydrator) amap() map[string]interface{} {
	n := h.unp.Len()
	m := make(map[string]interface{}, n)
	for ; n > 0; n-- {
		h.unp.Next()
		key := h.unp.String()
		h.unp.Next()
		m[key] = h.value()
	}
	return m
}

func (h *hydrator) array() []interface{} {
	n := h.unp.Len()
	a := make([]interface{}, n)
	for i := range a {
		h.unp.Next()
		a[i] = h.value()
	}
	return a
}

func (h *hydrator) record(n uint32) *db.Record {
	h.assertLength(1, n)
	if h.getErr() != nil {
		return nil
	}
	rec := db.Record{}
	h.unp.Next() // Detect array
	n = h.unp.Len()
	rec.Values = make([]interface{}, n)
	for i := range rec.Values {
		h.unp.Next()
		rec.Values[i] = h.value()
	}
	return &rec
}

func (h *hydrator) value() interface{} {
	switch h.unp.Curr {
	case packstream.PackedInt:
		return h.unp.Int()
	case packstream.PackedFloat:
		return h.unp.Float()
	case packstream.PackedStr:
		return h.unp.String()
	case packstream.PackedStruct:
		t := h.unp.StructTag()
		n := h.unp.Len()
		switch t {
		case 'N':
			return h.node(n)
		case 'R':
			return h.relationship(n)
		case 'r':
			return h.relationnode(n)
		case 'P':
			return h.path(n)
		case 'X':
			return h.point2d(n)
		case 'Y':
			return h.point3d(n)
		case 'F':
			return h.dateTimeOffset(n)
		case 'f':
			return h.dateTimeNamedZone(n)
		case 'd':
			return h.localDateTime(n)
		case 'D':
			return h.date(n)
		case 'T':
			return h.time(n)
		case 't':
			return h.localTime(n)
		case 'E':
			return h.duration(n)
		default:
			h.err = errors.New(fmt.Sprintf("Unknown tag: %02x", t))
			return nil
		}
	case packstream.PackedByteArray:
		return h.unp.ByteArray()
	case packstream.PackedArray:
		return h.array()
	case packstream.PackedMap:
		return h.amap()
	case packstream.PackedNil:
		return nil
	case packstream.PackedTrue:
		return true
	case packstream.PackedFalse:
		return false
	default:
		h.setErr(hydrationInvalidState)
		return nil
	}
}

// Trashes current value
func (h *hydrator) trash() {
	// TODO Less consuming implementation
	h.value()
}

func (h *hydrator) node(num uint32) interface{} {
	h.assertLength(3, num)
	if h.getErr() != nil {
		return nil
	}
	n := dbtype.Node{}
	h.unp.Next()
	n.Id = h.unp.Int()
	h.unp.Next()
	n.Labels = h.strings()
	h.unp.Next()
	n.Props = h.amap()
	return n
}

func (h *hydrator) relationship(n uint32) interface{} {
	h.assertLength(5, n)
	if h.getErr() != nil {
		return nil
	}
	r := dbtype.Relationship{}
	h.unp.Next()
	r.Id = h.unp.Int()
	h.unp.Next()
	r.StartId = h.unp.Int()
	h.unp.Next()
	r.EndId = h.unp.Int()
	h.unp.Next()
	r.Type = h.unp.String()
	h.unp.Next()
	r.Props = h.amap()
	return r
}

func (h *hydrator) relationnode(n uint32) interface{} {
	h.assertLength(3, n)
	if h.getErr() != nil {
		return nil
	}
	r := relNode{}
	h.unp.Next()
	r.id = h.unp.Int()
	h.unp.Next()
	r.name = h.unp.String()
	h.unp.Next()
	r.props = h.amap()
	return &r
}

func (h *hydrator) path(n uint32) interface{} {
	h.assertLength(3, n)
	if h.getErr() != nil {
		return nil
	}
	// Array of nodes
	h.unp.Next()
	num := h.unp.Int()
	nodes := make([]dbtype.Node, num)
	for i := range nodes {
		h.unp.Next()
		node, ok := h.value().(dbtype.Node)
		if !ok {
			h.setErr(errors.New("Path hydrate error"))
			return nil
		}
		nodes[i] = node
	}
	// Array of relnodes
	h.unp.Next()
	num = h.unp.Int()
	rnodes := make([]*relNode, num)
	for i := range rnodes {
		h.unp.Next()
		rnode, ok := h.value().(*relNode)
		if !ok {
			h.setErr(errors.New("Path hydrate error"))
			return nil
		}
		rnodes[i] = rnode
	}
	// Array of indexes
	h.unp.Next()
	num = h.unp.Int()
	indexes := make([]int, num)
	for i := range indexes {
		h.unp.Next()
		indexes[i] = int(h.unp.Int())
	}

	if (len(indexes) & 0x01) == 1 {
		h.setErr(errors.New("Path hydrate error"))
		return nil
	}

	return buildPath(nodes, rnodes, indexes)
}

func (h *hydrator) point2d(n uint32) interface{} {
	p := dbtype.Point2D{}
	h.unp.Next()
	p.SpatialRefId = uint32(h.unp.Int())
	h.unp.Next()
	p.X = h.unp.Float()
	h.unp.Next()
	p.Y = h.unp.Float()
	return p
}

func (h *hydrator) point3d(n uint32) interface{} {
	p := dbtype.Point3D{}
	h.unp.Next()
	p.SpatialRefId = uint32(h.unp.Int())
	h.unp.Next()
	p.X = h.unp.Float()
	h.unp.Next()
	p.Y = h.unp.Float()
	h.unp.Next()
	p.Z = h.unp.Float()
	return p
}

func (h *hydrator) dateTimeOffset(n uint32) interface{} {
	h.unp.Next()
	secs := h.unp.Int()
	h.unp.Next()
	nans := h.unp.Int()
	h.unp.Next()
	offs := h.unp.Int()
	t := time.Unix(secs, nans).UTC()
	l := time.FixedZone("Offset", int(offs))
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), l)
}

func (h *hydrator) dateTimeNamedZone(n uint32) interface{} {
	h.unp.Next()
	secs := h.unp.Int()
	h.unp.Next()
	nans := h.unp.Int()
	h.unp.Next()
	zone := h.unp.String()
	t := time.Unix(secs, nans).UTC()
	l, err := time.LoadLocation(zone)
	if err != nil {
		h.setErr(err)
		return nil
	}
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), l)
}

func (h *hydrator) localDateTime(n uint32) interface{} {
	h.unp.Next()
	secs := h.unp.Int()
	h.unp.Next()
	nans := h.unp.Int()
	t := time.Unix(secs, nans).UTC()
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	return dbtype.LocalDateTime(t)
}

func (h *hydrator) date(n uint32) interface{} {
	h.unp.Next()
	days := h.unp.Int()
	secs := days * 86400
	return dbtype.Date(time.Unix(secs, 0).UTC())
}

func (h *hydrator) time(n uint32) interface{} {
	h.unp.Next()
	nans := h.unp.Int()
	h.unp.Next()
	offs := h.unp.Int()
	secs := nans / int64(time.Second)
	nans -= secs * int64(time.Second)
	l := time.FixedZone("Offset", int(offs))
	t := time.Date(0, 0, 0, 0, 0, int(secs), int(nans), l)
	return dbtype.Time(t)
}

func (h *hydrator) localTime(n uint32) interface{} {
	h.unp.Next()
	nans := h.unp.Int()
	secs := nans / int64(time.Second)
	nans -= secs * int64(time.Second)
	t := time.Date(0, 0, 0, 0, 0, int(secs), int(nans), time.Local)
	return dbtype.LocalTime(t)
}

func (h *hydrator) duration(n uint32) interface{} {
	h.unp.Next()
	mon := h.unp.Int()
	h.unp.Next()
	day := h.unp.Int()
	h.unp.Next()
	sec := h.unp.Int()
	h.unp.Next()
	nan := h.unp.Int()
	return dbtype.Duration{Months: mon, Days: day, Seconds: sec, Nanos: int(nan)}
}

func parseNotifications(notificationsx []interface{}) []db.Notification {
	var notifications []db.Notification
	if len(notificationsx) > 0 {
		notifications = make([]db.Notification, 0, len(notificationsx))
		for _, x := range notificationsx {
			notificationx, ok := x.(map[string]interface{})
			if ok {
				notifications = append(notifications, parseNotification(notificationx))
			}
		}
	}
	return notifications
}

func parsePlanOpIdArgsChildren(planx map[string]interface{}) (string, []string, map[string]interface{}, []interface{}) {
	operator, _ := planx["operatorType"].(string)
	identifiersx, _ := planx["identifiers"].([]interface{})
	arguments, _ := planx["args"].(map[string]interface{})

	identifiers := make([]string, len(identifiersx))
	for i, id := range identifiersx {
		identifiers[i], _ = id.(string)
	}

	childrenx, _ := planx["children"].([]interface{})

	return operator, identifiers, arguments, childrenx
}

func parsePlan(planx map[string]interface{}) *db.Plan {
	op, ids, args, childrenx := parsePlanOpIdArgsChildren(planx)
	plan := &db.Plan{
		Operator:    op,
		Arguments:   args,
		Identifiers: ids,
	}

	plan.Children = make([]db.Plan, 0, len(childrenx))
	for _, c := range childrenx {
		childPlanx, _ := c.(map[string]interface{})
		if len(childPlanx) > 0 {
			childPlan := parsePlan(childPlanx)
			if childPlan != nil {
				plan.Children = append(plan.Children, *childPlan)
			}
		}
	}

	return plan
}

func parseProfile(profilex map[string]interface{}) *db.ProfiledPlan {
	op, ids, args, childrenx := parsePlanOpIdArgsChildren(profilex)
	plan := &db.ProfiledPlan{
		Operator:    op,
		Arguments:   args,
		Identifiers: ids,
	}

	plan.DbHits, _ = profilex["dbHits"].(int64)
	plan.Records, _ = profilex["rows"].(int64)

	plan.Children = make([]db.ProfiledPlan, 0, len(childrenx))
	for _, c := range childrenx {
		childPlanx, _ := c.(map[string]interface{})
		if len(childPlanx) > 0 {
			childPlan := parseProfile(childPlanx)
			if childPlan != nil {
				plan.Children = append(plan.Children, *childPlan)
			}
		}
	}

	return plan
}

func parseNotification(m map[string]interface{}) db.Notification {
	n := db.Notification{}
	n.Code, _ = m["code"].(string)
	n.Description = m["description"].(string)
	n.Severity, _ = m["severity"].(string)
	n.Title, _ = m["title"].(string)
	posx, exists := m["position"].(map[string]interface{})
	if exists {
		pos := &db.InputPosition{}
		i, _ := posx["column"].(int64)
		pos.Column = int(i)
		i, _ = posx["line"].(int64)
		pos.Line = int(i)
		i, _ = posx["offset"].(int64)
		pos.Offset = int(i)
		n.Position = pos
	}

	return n
}
