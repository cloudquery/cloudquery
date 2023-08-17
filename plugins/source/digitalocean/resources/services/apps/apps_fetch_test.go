package apps_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/apps"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func TestApps(t *testing.T) {
	client.MockTestHelper(t, apps.Apps(), newMockApps, client.TestOptions{})
}

func newMockApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	var (
		ctxInterface = reflect.TypeOf((*context.Context)(nil)).Elem()
		isCtx        = gomock.AssignableToTypeOf(ctxInterface)
	)

	svc := mocks.NewMockAppsService(ctrl)

	// Page 1.
	svc.EXPECT().List(isCtx, &godo.ListOptions{Page: 1, PerPage: client.MaxItemsPerPage}).
		Return(
			[]*godo.App{
				fakeApp("11111111-0000-4000-0000-000000000000"),
				fakeApp("22222222-0000-4000-0000-000000000000"),
			},
			&godo.Response{
				Links: &godo.Links{
					Pages: &godo.Pages{
						First: "https://api.digitalocean.com/v2/apps",
						Next:  "https://api.digitalocean.com/v2/apps?page=2",
						Last:  "https://api.digitalocean.com/v2/apps?page=2",
					},
				},
			},
			nil,
		)
	svc.EXPECT().ListAlerts(isCtx, "11111111-0000-4000-0000-000000000000").
		Return(
			[]*godo.AppAlert{
				fakeAlert("11111111-1111-4000-0000-000000000000"),
				fakeAlert("11111111-2222-4000-0000-000000000000"),
			},
			&godo.Response{},
			nil,
		)
	svc.EXPECT().ListAlerts(isCtx, "22222222-0000-4000-0000-000000000000").
		Return(
			[]*godo.AppAlert{
				fakeAlert("22222222-1111-4000-0000-000000000000"),
				fakeAlert("22222222-2222-4000-0000-000000000000"),
			},
			&godo.Response{},
			nil,
		)

	// Page 2.
	svc.EXPECT().List(isCtx, &godo.ListOptions{Page: 2, PerPage: client.MaxItemsPerPage}).
		Return(
			[]*godo.App{
				fakeApp("33333333-0000-4000-0000-000000000000"),
				fakeApp("44444444-0000-4000-0000-000000000000"),
			},
			&godo.Response{
				Links: &godo.Links{
					Pages: &godo.Pages{
						First: "https://api.digitalocean.com/v2/apps",
						Prev:  "https://api.digitalocean.com/v2/apps",
						Last:  "https://api.digitalocean.com/v2/apps?page=2",
					},
				},
			},
			nil,
		)
	svc.EXPECT().ListAlerts(isCtx, "33333333-0000-4000-0000-000000000000").
		Return(
			[]*godo.AppAlert{
				fakeAlert("33333333-1111-4000-0000-000000000000"),
				fakeAlert("33333333-2222-4000-0000-000000000000"),
			},
			&godo.Response{},
			nil,
		)
	svc.EXPECT().ListAlerts(isCtx, "44444444-0000-4000-0000-000000000000").
		Return(
			[]*godo.AppAlert{
				fakeAlert("44444444-1111-4000-0000-000000000000"),
				fakeAlert("44444444-2222-4000-0000-000000000000"),
			},
			&godo.Response{},
			nil,
		)

	return client.Services{Apps: svc}
}

func fakeApp(id string) *godo.App {
	var app godo.App
	if err := faker.FakeObject(&app); err != nil {
		panic(err)
	}
	app.ID = id
	return &app
}

func fakeAlert(id string) *godo.AppAlert {
	var alert godo.AppAlert
	if err := faker.FakeObject(&alert); err != nil {
		panic(err)
	}
	alert.ID = id
	return &alert
}
