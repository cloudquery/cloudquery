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

package neo4j

import (
	"crypto/x509"
	"math"
	"net/url"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/log"
)

// A Config contains options that can be used to customize certain
// aspects of the driver
type Config struct {
	// RootCAs defines the set of certificate authorities that the driver trusts. If set
	// to nil, the driver uses hosts system certificates.
	//
	// The trusted certificates are used to validate connections for URI schemes 'bolt+s'
	// and 'neo4j+s'.
	RootCAs *x509.CertPool

	// Logging target the driver will send its log outputs
	//
	// Possible to use custom logger (implement neo4j.log.Logger interface) or
	// use neo4j.ConsoleLog.
	//
	// default: No Op Logger (neo4j.VoidLog)
	Log log.Logger
	// Resolver that would be used to resolve initial router address. This may
	// be useful if you want to provide more than one URL for initial router.
	// If not specified, the URL provided to NewDriver is used as the initial
	// router.
	//
	// default: nil
	AddressResolver ServerAddressResolver
	// Maximum amount of time a retriable operation would continue retrying. It
	// cannot be specified as a negative value.
	//
	// default: 30 * time.Second
	MaxTransactionRetryTime time.Duration
	// Maximum number of connections per URL to allow on this driver. It
	// cannot be specified as 0 and negative values are interpreted as
	// math.MaxInt32.
	//
	// default: 100
	MaxConnectionPoolSize int
	// Maximum connection life time on pooled connections. Values less than
	// or equal to 0 disables the lifetime check.
	//
	// default: 1 * time.Hour
	MaxConnectionLifetime time.Duration
	// Maximum amount of time to either acquire an idle connection from the pool
	// or create a new connection (when the pool is not full). Negative values
	// result in an infinite wait time where 0 value results in no timeout which
	// results in immediate failure when there are no available connections.
	//
	// default: 1 * time.Minute
	ConnectionAcquisitionTimeout time.Duration
	// Connect timeout that will be set on underlying sockets. Values less than
	// or equal to 0 results in no timeout being applied.
	//
	// default: 5 * time.Second
	SocketConnectTimeout time.Duration
	// Whether to enable TCP keep alive on underlying sockets.
	//
	// default: true
	SocketKeepalive bool
	// Optionally override the user agent string sent to Neo4j server.
	//
	// default: neo4j.UserAgent
	UserAgent string
}

func defaultConfig() *Config {
	return &Config{
		AddressResolver:              nil,
		MaxTransactionRetryTime:      30 * time.Second,
		MaxConnectionPoolSize:        100,
		MaxConnectionLifetime:        1 * time.Hour,
		ConnectionAcquisitionTimeout: 1 * time.Minute,
		SocketConnectTimeout:         5 * time.Second,
		SocketKeepalive:              true,
		RootCAs:                      nil,
		UserAgent:                    UserAgent,
	}
}

func validateAndNormaliseConfig(config *Config) error {
	// Max Transaction Retry Time
	if config.MaxTransactionRetryTime < 0 {
		return &UsageError{Message: "Maximum transaction retry time cannot be smaller than 0"}
	}

	// Max Connection Pool Size
	if config.MaxConnectionPoolSize == 0 {
		return &UsageError{Message: "Maximum connection pool cannot be 0"}
	}

	if config.MaxConnectionPoolSize < 0 {
		config.MaxConnectionPoolSize = math.MaxInt32
	}

	// Max Connection Lifetime
	if config.MaxConnectionLifetime < 0 {
		config.MaxConnectionLifetime = 0
	}

	// Connection Acquisition Timeout
	if config.ConnectionAcquisitionTimeout < 0 {
		config.ConnectionAcquisitionTimeout = -1
	}

	// Socket Connect Timeout
	if config.SocketConnectTimeout < 0 {
		config.SocketConnectTimeout = 0
	}

	return nil
}

// ServerAddress represents a host and port. Host can either be an IP address or a DNS name.
// Both IPv4 and IPv6 hosts are supported.
type ServerAddress interface {
	// Hostname returns the host portion of this ServerAddress.
	Hostname() string
	// Port returns the port portion of this ServerAddress.
	Port() string
}

// ServerAddressResolver is a function type that defines the resolver function used by the routing driver to
// resolve the initial address used to create the driver.
type ServerAddressResolver func(address ServerAddress) []ServerAddress

func newServerAddressURL(hostname string, port string) *url.URL {
	if hostname == "" {
		return nil
	}

	hostAndPort := hostname
	if port != "" {
		hostAndPort = hostAndPort + ":" + port
	}

	return &url.URL{Host: hostAndPort}
}

// NewServerAddress generates a ServerAddress with provided hostname and port information.
func NewServerAddress(hostname string, port string) ServerAddress {
	return newServerAddressURL(hostname, port)
}
