package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovePII(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{
			"msg AccountID 123456789 something",
			"msg AccountID xxxx something",
		},
		{
			"msg 123456789 something",
			"msg 1234xxxxxxxx something",
		},
		{
			"operation error S3: GetBucketLogging, https response error StatusCode: 404, RequestID: 3PQRRTJ1BAB82DWH, HostID: MZfZa61jiI+47KWhZjIl1N92GVKOMhslal/A2dcWMJwk7rZazsIflh1LNa3yIDqjrRpF1fF17/k=, api error NoSuchBucket: The specified bucket does not exist",
			"operation error S3: GetBucketLogging, https response error StatusCode: 404, RequestID: xxxx, HostID: xxxx, api error NoSuchBucket: The specified bucket does not exist",
		},
		{
			"AccessDenied: User: arn:aws:sts::123456789:assumed-role/some-role/i-012304405c679abcd is not authorized to perform: sts:AssumeRole on resource: arn:aws:iam::123456789:role/other-role\n\tstatus code: 403, request id: d2f12332-d1f2-12c5-1234-abc12345d123",
			"AccessDenied: User: arn:aws:xxxx is not authorized to perform: sts:AssumeRole on resource: arn:aws:xxxx code: 403, request id: xxxx",
		},
		{
			"IAM: GetUserPolicy - User: arn:aws:sts::123456789:assumed-role/some-role/i-012304405c679abcd is not authorized to perform: iam:GetUserPolicy on resource: user some_user",
			"IAM: GetUserPolicy - User: arn:aws:xxxx is not authorized to perform: iam:GetUserPolicy on resource: user xxxx",
		},
		{
			"operation error CloudWatch Logs: DescribeMetricFilters, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://logs.eu-central-1.amazonaws.com/\": dial tcp: lookup logs.eu-central-1.amazonaws.com on 192.168.1.1:53: read udp 192.168.1.2:5353->192.168.1.1:53: i/o timeout",
			"operation error CloudWatch Logs: DescribeMetricFilters, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://xxxx\": dial tcp: lookup xxxx: read udp xxxx->xxxx: i/o timeout",
		},
		{
			"EC2: DescribeImageAttribute - You are not authorized to perform this operation. Encoded authorization failure message: SOMEENCODEDMESSAGEWITHNUMBERS1234567ANDDASHANDUNDERSCORES-ABCDE_123123123_EXAMPLEMESSAGE",
			"EC2: DescribeImageAttribute - You are not authorized to perform this operation. Encoded authorization failure message: xxxx",
		},
		{
			"operation error Elastic Beanstalk: DescribeConfigurationOptions, https response error StatusCode: 400, RequestID: 3PQRRTJ1BAB82DWH, api error InvalidParameterValue: Access Denied: S3Bucket=some-bucket-1, S3Key=object_path/some_key.ext (Service: Amazon S3; Status Code: 403; Error Code: AccessDenied; Request ID: 3PQRRTJ1BAB82DWH; Proxy: null)",
			"operation error Elastic Beanstalk: DescribeConfigurationOptions, https response error StatusCode: 400, RequestID: xxxx, api error InvalidParameterValue: Access Denied: S3Bucket=xxxx, S3Key=xxxx (Service: Amazon S3; Status Code: 403; Error Code: AccessDenied; Request ID: xxxx; Proxy: null)",
		},
		{
			"operation error Direct Connect: DescribeVirtualInterfaces, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://logs.eu-central-1.amazonaws.com/\": dial tcp 177.72.244.112:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.",
			"operation error Direct Connect: DescribeVirtualInterfaces, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://xxxx\": dial tcp xxxx: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.",
		},
	}
	for i, tc := range cases {
		res := removePII([]Account{{ID: "123456789"}}, tc.Input)
		assert.Equalf(t, tc.Expected, res, "Case #%d", i+1)
	}
}
