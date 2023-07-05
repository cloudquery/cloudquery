package monitors

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildMonitorsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockMonitorsAPIClient(ctrl)
	d := mocks.NewMockDowntimesAPIClient(ctrl)
	services := client.DatadogServices{
		MonitorsAPI:  m,
		DowntimesAPI: d,
	}

	var monitors []datadogV1.Monitor
	err := faker.FakeObject(&monitors)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	monitors[0].Deleted.Set(&now)
	priority := int64(123)
	monitors[0].Priority.Set(&priority)

	m.EXPECT().ListMonitors(gomock.Any()).Return(monitors, nil, nil)

	var dt []datadogV1.Downtime
	err = faker.FakeObject(&dt)
	if err != nil {
		t.Fatal(err)
	}
	i64val := int64(123)
	i32val := int32(123)
	textVal := "test string"
	dt[0].ActiveChild.Set(datadogV1.NewDowntimeChild())
	dt[0].Canceled.Set(&i64val)
	dt[0].End.Set(&i64val)
	dt[0].MonitorId.Set(&i64val)
	dt[0].ParentId.Set(&i64val)
	dt[0].Recurrence.Set(datadogV1.NewDowntimeRecurrence())
	dt[0].UpdaterId.Set(&i32val)
	dt[0].Message.Set(&textVal)

	d.EXPECT().ListMonitorDowntimes(gomock.Any(), gomock.Any()).Return(dt, nil, nil)

	return services
}

func TestMonitors(t *testing.T) {
	client.DatadogMockTestHelper(t, Monitors(), buildMonitorsMock, client.TestOptions{})
}
