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
	"context"
	"errors"
	"sync"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

const missingWriterRetries = 100
const missingReaderRetries = 100

type databaseRouter struct {
	dueUnix int64
	table   *db.RoutingTable
}

// Thread safe
type Router struct {
	routerContext map[string]string
	pool          Pool
	dbRouters     map[string]*databaseRouter
	dbRoutersMut  sync.Mutex
	now           func() time.Time
	sleep         func(time.Duration)
	rootRouter    string
	getRouters    func() []string
	log           log.Logger
	logId         string
}

type Pool interface {
	Borrow(ctx context.Context, servers []string, wait bool) (db.Connection, error)
	Return(c db.Connection)
}

func New(rootRouter string, getRouters func() []string, routerContext map[string]string, pool Pool, logger log.Logger, logId string) *Router {
	r := &Router{
		rootRouter:    rootRouter,
		getRouters:    getRouters,
		routerContext: routerContext,
		pool:          pool,
		dbRouters:     make(map[string]*databaseRouter),
		now:           time.Now,
		sleep:         time.Sleep,
		log:           logger,
		logId:         logId,
	}
	r.log.Infof(log.Router, r.logId, "Created {context: %v}", routerContext)
	return r
}

func (r *Router) getTable(database string) (*db.RoutingTable, error) {
	now := r.now()

	r.dbRoutersMut.Lock()
	defer r.dbRoutersMut.Unlock()

	dbRouter := r.dbRouters[database]
	if dbRouter != nil && now.Unix() < dbRouter.dueUnix {
		return dbRouter.table, nil
	}

	var (
		table *db.RoutingTable
		err   error
	)

	// Try last known set of routers if there are any
	if dbRouter != nil && len(dbRouter.table.Routers) > 0 {
		routers := dbRouter.table.Routers
		r.log.Infof(log.Router, r.logId, "Reading routing table for '%s' from previously known routers: %v", database, routers)
		table, err = readTable(context.Background(), r.pool, database, routers, r.routerContext)
	}

	// Try initial router if no routers or failed
	if table == nil || err != nil {
		r.log.Infof(log.Router, r.logId, "Reading routing table from initial router: %s", r.rootRouter)
		table, err = readTable(context.Background(), r.pool, database, []string{r.rootRouter}, r.routerContext)
	}

	// Use hook to retrieve possibly different set of routers and retry
	if err != nil && r.getRouters != nil {
		routers := r.getRouters()
		r.log.Infof(log.Router, r.logId, "Reading routing table for '%s' from custom routers: %v", routers)
		table, err = readTable(context.Background(), r.pool, database, routers, r.routerContext)
	}

	if err != nil {
		r.log.Error(log.Router, r.logId, err)
		return nil, err
	}

	if table == nil {
		// Safe guard for logical error somewhere else
		err = errors.New("No error and no table")
		r.log.Error(log.Router, r.logId, err)
		return nil, err
	}

	// Store the routing table
	r.dbRouters[database] = &databaseRouter{
		table:   table,
		dueUnix: now.Add(time.Duration(table.TimeToLive) * time.Second).Unix(),
	}
	r.log.Debugf(log.Router, r.logId, "New routing table for '%s', TTL %d", database, table.TimeToLive)

	return table, nil
}

func (r *Router) Readers(database string) ([]string, error) {
	table, err := r.getTable(database)
	if err != nil {
		return nil, err
	}

	// During startup we can get tables without any readers
	retries := missingReaderRetries
	for len(table.Readers) == 0 {
		retries--
		if retries == 0 {
			break
		}
		r.log.Infof(log.Router, r.logId, "Invalidating routing table, no readers")
		r.Invalidate(database)
		r.sleep(100 * time.Millisecond)
		table, err = r.getTable(database)
		if err != nil {
			return nil, err
		}
	}
	if len(table.Readers) == 0 {
		return nil, wrapError(r.rootRouter, errors.New("No readers"))
	}

	return table.Readers, nil
}

func (r *Router) Writers(database string) ([]string, error) {
	table, err := r.getTable(database)
	if err != nil {
		return nil, err
	}

	// During election we can get tables without any writers
	retries := missingWriterRetries
	for len(table.Writers) == 0 {
		retries--
		if retries == 0 {
			break
		}
		r.log.Infof(log.Router, r.logId, "Invalidating routing table, no writers")
		r.Invalidate(database)
		r.sleep(100 * time.Millisecond)
		table, err = r.getTable(database)
		if err != nil {
			return nil, err
		}
	}
	if len(table.Writers) == 0 {
		return nil, wrapError(r.rootRouter, errors.New("No writers"))
	}

	return table.Writers, nil
}

func (r *Router) Context() map[string]string {
	return r.routerContext
}

func (r *Router) Invalidate(database string) {
	r.log.Infof(log.Router, r.logId, "Invalidating routing table for '%s'", database)
	r.dbRoutersMut.Lock()
	defer r.dbRoutersMut.Unlock()
	// Reset due time to the 70s, this will make next access refresh the routing table using
	// last set of routers instead of the original one.
	dbRouter := r.dbRouters[database]
	if dbRouter != nil {
		dbRouter.dueUnix = 0
	}
}

func (r *Router) CleanUp() {
	r.log.Debugf(log.Router, r.logId, "Cleaning up")
	now := r.now().Unix()
	r.dbRoutersMut.Lock()
	defer r.dbRoutersMut.Unlock()

	for dbName, dbRouter := range r.dbRouters {
		if now > dbRouter.dueUnix {
			delete(r.dbRouters, dbName)
		}
	}
}
