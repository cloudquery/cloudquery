// +build mock

package lambda

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildLambdaRuntimesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)
	return client.Services{
		Lambda: m,
	}
}

func TestLambdaRuntimes(t *testing.T) {
	client.AwsMockTestHelper(t, LambdaRuntimes(), buildLambdaRuntimesMock, client.TestOptions{})
}
