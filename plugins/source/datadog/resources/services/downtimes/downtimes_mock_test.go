package downtimes

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildDowntimesMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockDowntimesAPIClient(ctrl)
	services := client.DatadogServices{
		DowntimesAPI: m,
	}

	var d []datadogV1.Downtime
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	i64val := int64(123)
	i32val := int32(123)
	textVal := "test string"
	d[0].ActiveChild.Set(datadogV1.NewDowntimeChild())
	d[0].Canceled.Set(&i64val)
	d[0].End.Set(&i64val)
	d[0].MonitorId.Set(&i64val)
	d[0].ParentId.Set(&i64val)
	d[0].Recurrence.Set(datadogV1.NewDowntimeRecurrence())
	d[0].UpdaterId.Set(&i32val)
	d[0].Message.Set(&textVal)

	m.EXPECT().ListDowntimes(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestDowntimes(t *testing.T) {
	client.DatadogMockTestHelper(t, Downtimes(), buildDowntimesMock, client.TestOptions{})
}
