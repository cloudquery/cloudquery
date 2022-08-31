package client

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	heroku "github.com/heroku/heroku-go/v5"
)

// Paginator implements the http.RoundTripper interface to intercept
// pagination information not supported by the official Heroku SDK. It
// injects this information into the request context so that resolvers
// can make additional calls, if necessary.
type Paginator struct {
	transport http.RoundTripper
}

var reNextRange = regexp.MustCompile(`(?P<field>\w+) (?P<firstID>[\w\[\]\-]*)\.\.(?P<lastID>[\w\[\]\-]*)`)

// NewPaginator returns a new paginator with the given RoundTripper
func NewPaginator(t http.RoundTripper) Paginator {
	return Paginator{
		transport: t,
	}
}

// RoundTrip is an implementation of the http.RoundTripper function
// that checks the status code and adds pagination information to the
// context accordingly. This is only done because the official SDK lacks
// support for pagination (see: https://github.com/heroku/heroku-go/issues/56)
func (p Paginator) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := p.transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	ctx := req.Context()
	nr := ctx.Value("nextRange")
	if nr != nil {
		v := nr.(*heroku.ListRange)
		if resp.StatusCode == http.StatusPartialContent {
			nextRange := resp.Header.Get("Next-Range")
			var listRange *heroku.ListRange
			listRange, err = parseNextRange(nextRange)
			if err != nil {
				return nil, err
			}

			// copy values to nextRange pointer
			v.Max = listRange.Max
			v.FirstID = listRange.FirstID
			v.LastID = listRange.LastID
			v.Descending = listRange.Descending
			v.Field = listRange.Field
		} else {
			// we indicate end of pagination with max = 0
			v.Max = 0
		}
	}
	return resp, err
}

// parseNextRange implements a parser for the Heroku Next-Range format that returns
// a *heroku.ListRange.
func parseNextRange(v string) (*heroku.ListRange, error) {
	parts := strings.Split(v, ";")
	if len(parts) < 2 {
		return nil, fmt.Errorf("got unexpected Next-Range header value: %q", v)
	}
	lr := new(heroku.ListRange)
	match := reNextRange.FindStringSubmatch(parts[0])
	for i, name := range reNextRange.SubexpNames() {
		switch name {
		case "field":
			lr.Field = match[i]
		case "firstID":
			lr.FirstID = match[i]
		case "lastID":
			lr.LastID = match[i]
		}
	}

	more := strings.Split(parts[1], ",")
	var err error
	for _, m := range more {
		m = strings.TrimSpace(m)
		switch {
		case strings.HasPrefix(m, "order="):
			lr.Descending = strings.TrimPrefix(m, "order=") == "desc"
		case strings.HasPrefix(m, "max="):
			lr.Max, err = strconv.Atoi(strings.TrimPrefix(m, "max="))
			if err != nil {
				return nil, fmt.Errorf("got unexpected Next-Range header max value: %q", v)
			}
		default:
			return nil, fmt.Errorf("unhandled Next-Range case: %v", m)
		}
	}
	return lr, nil
}
