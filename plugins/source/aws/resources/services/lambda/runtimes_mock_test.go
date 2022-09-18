package lambda

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildLambdaRuntimesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)
	return client.Services{
		Lambda: m,
	}
}

func TestLambdaRuntimes(t *testing.T) {
	client.AwsMockTestHelper(t, Runtimes(), buildLambdaRuntimesMock, client.TestOptions{})
}
