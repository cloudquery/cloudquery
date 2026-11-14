package client

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestBuildDeleteCQIDsSQL(t *testing.T) {
	c := &Client{}
	ids := []uuid.UUID{
		uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	}

	require.Equal(t,
		`delete from "embeddings" where "_cq_id" in ('11111111-1111-1111-1111-111111111111','22222222-2222-2222-2222-222222222222')`,
		c.buildDeleteCQIDsSQL("embeddings", ids),
	)
	require.Equal(t,
		`delete from "emb""; drop table users; --" where "_cq_id" in ('11111111-1111-1111-1111-111111111111')`,
		c.buildDeleteCQIDsSQL(`emb"; drop table users; --`, ids[:1]),
	)
}
