// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"fmt"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
