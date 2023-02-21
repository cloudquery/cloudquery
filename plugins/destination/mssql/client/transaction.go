package client

import (
	"context"
	"database/sql"
)

type txOp func(tx *sql.Tx) error

func (c *Client) doInTx(ctx context.Context, op txOp) (err error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		// also handles tx.Commit() err
		if err != nil {
			c.logErr(err)
			_ = tx.Rollback()
		}
	}()

	return op(tx)
}
