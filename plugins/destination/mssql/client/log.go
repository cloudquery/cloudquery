package client

import (
	"context"
	"errors"

	mssql "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/msdsn"
)

var _ mssql.ContextLogger = (*Client)(nil)

func (c *Client) Log(_ context.Context, _ msdsn.Log, msg string) {
	c.logger.Debug().Msg(msg)
}

func (c *Client) logErr(err error) {
	var dbErr mssql.Error
	if !errors.As(err, &dbErr) {
		return
	}
	c.logMSSQLError(dbErr)
}

func (c *Client) logMSSQLError(err mssql.Error) {
	if len(err.All) > 0 {
		for _, e := range err.All {
			c.logMSSQLError(e)
		}
		return
	}

	c.logger.Error().
		Int32("number", err.Number).
		Uint8("state", err.State).
		Uint8("class", err.Class).
		Str("server_name", err.ServerName).
		Str("proc_name", err.ProcName).
		Int32("line_no", err.LineNo).
		Msg(err.Message)
}
