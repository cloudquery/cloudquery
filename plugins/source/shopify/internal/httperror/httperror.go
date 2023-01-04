package httperror

import (
	"fmt"
	"net/http"
)

type Error struct {
	method string
	edge   string

	statusCode int
	status     string
	body       string
}

func New(statusCode int, method, edge, status, body string) Error {
	return Error{
		method:     method,
		edge:       edge,
		statusCode: statusCode,
		status:     status,
		body:       body,
	}
}

func (e Error) Error() string {
	if e.statusCode == http.StatusTooManyRequests {
		return fmt.Sprintf("HTTP error: Sending %s request to %s: %s <omitted>", e.method, e.edge, e.status)
	}

	return fmt.Sprintf("HTTP error: Sending %s request to %s: %s %q", e.method, e.edge, e.status, e.body)
}

func (e Error) Status() string {
	return e.status
}

func (e Error) StatusCode() int {
	return e.statusCode
}

func (e Error) String() string {
	return e.Error()
}

func (e Error) Temporary() bool {
	return e.statusCode == http.StatusTooManyRequests || e.statusCode == http.StatusInternalServerError
}

func (e Error) Body() string {
	return e.body
}
