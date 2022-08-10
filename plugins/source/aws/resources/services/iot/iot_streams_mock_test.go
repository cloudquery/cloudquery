package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotStreamsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	streams := iot.ListStreamsOutput{}
	err := faker.FakeData(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.NextToken = nil
	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streams, nil)

	streamOutput := iot.DescribeStreamOutput{}
	err = faker.FakeData(&streamOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeStream(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streamOutput, nil)

	return client.Services{
		IOT: m,
	}
}

func TestIotStreams(t *testing.T) {
	client.AwsMockTestHelper(t, IotStreams(), buildIotStreamsMock, client.TestOptions{})
}
