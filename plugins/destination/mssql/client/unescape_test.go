package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnescape(t *testing.T) {
	type testCase struct {
		have, want string
	}
	for _, tc := range []testCase{
		{
			have: `{\n  \"Version\" : \"2012-10-17\",\n  \"Statement\" : [ {\n    \"Sid\" : \"\",\n    \"Effect\" : \"Allow\",\n    \"Principal\" : {\n      \"AWS\" : \"arn:aws:iam::12345:role/service-role/AmazonGrafanaServiceRole-XYZ\"\n    },\n    \"Action\" : [ \"ecr:CreateRepository\", \"ecr:BatchImportUpstreamImage\" ],\n    \"Resource\" : \"arn:aws:ecr:us-east-2:1234:repository//*\"\n  } ]\n}`,
			want: `{
  "Version" : "2012-10-17",
  "Statement" : [ {
    "Sid" : "",
    "Effect" : "Allow",
    "Principal" : {
      "AWS" : "arn:aws:iam::12345:role/service-role/AmazonGrafanaServiceRole-XYZ"
    },
    "Action" : [ "ecr:CreateRepository", "ecr:BatchImportUpstreamImage" ],
    "Resource" : "arn:aws:ecr:us-east-2:1234:repository//*"
  } ]
}`,
		},
		{
			have: `{"BillTo":"P\u0026T"}`,
			want: `{"BillTo":"P&T"}`,
		},
		{
			have: `{"BillTo":"P&T"}`,
			want: `{"BillTo":"P&T"}`,
		},
	} {
		require.Equal(t, tc.want, unescape(tc.have))
	}
}
