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

// AuthToken contains credentials to be sent over to the neo4j server.
type AuthToken struct {
	tokens map[string]interface{}
}

const keyScheme = "scheme"
const schemeNone = "none"
const schemeBasic = "basic"
const schemeKerberos = "kerberos"
const keyPrincipal = "principal"
const keyCredentials = "credentials"
const keyRealm = "realm"
const keyTicket = "ticket"

// NoAuth generates an empty authentication token
func NoAuth() AuthToken {
	return AuthToken{tokens: map[string]interface{}{
		keyScheme: schemeNone,
	}}
}

// BasicAuth generates a basic authentication token with provided username, password and realm
func BasicAuth(username string, password string, realm string) AuthToken {
	tokens := map[string]interface{}{
		keyScheme:      schemeBasic,
		keyPrincipal:   username,
		keyCredentials: password,
	}

	if realm != "" {
		tokens[keyRealm] = realm
	}

	return AuthToken{tokens: tokens}
}

// KerberosAuth generates a kerberos authentication token with provided base-64 encoded kerberos ticket
func KerberosAuth(ticket string) AuthToken {
	token := AuthToken{
		tokens: map[string]interface{}{
			keyScheme: schemeKerberos,
			keyTicket: ticket,
		},
	}

	return token
}

// CustomAuth generates a custom authentication token with provided parameters
func CustomAuth(scheme string, username string, password string, realm string, parameters map[string]interface{}) AuthToken {
	tokens := map[string]interface{}{
		keyScheme:      scheme,
		keyPrincipal:   username,
		keyCredentials: password,
	}

	if realm != "" {
		tokens[keyRealm] = realm
	}

	if parameters != nil {
		tokens["parameters"] = parameters
	}

	return AuthToken{tokens: tokens}
}
