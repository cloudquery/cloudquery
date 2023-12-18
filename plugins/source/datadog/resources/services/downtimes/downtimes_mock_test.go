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
	c := datadogV1.NewDowntimeChild()
	c.AdditionalProperties = map[string]any{"key": "value"}
	d[0].ActiveChild.Set(c)
	d[0].Canceled.Set(&i64val)
	d[0].End.Set(&i64val)
	d[0].MonitorId.Set(&i64val)
	d[0].ParentId.Set(&i64val)
	r := datadogV1.NewDowntimeRecurrence()
	r.AdditionalProperties = map[string]any{"key": "value"}
	d[0].Recurrence.Set(r)
	d[0].UpdaterId.Set(&i32val)
	d[0].Message.Set(&textVal)
	d[0].AdditionalProperties = map[string]any{"key": "value"}

	m.EXPECT().ListDowntimes(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestDowntimes(t *testing.T) {
	client.DatadogMockTestHelper(t, Downtimes(), buildDowntimesMock, client.TestOptions{})
}
