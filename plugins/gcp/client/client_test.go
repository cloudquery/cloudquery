package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/serviceusage/v1"
)

func TestAppendWithoutDupes(t *testing.T) {
	cases := []struct {
		Base     []string
		Add      []string
		Expected []string
	}{
		{
			Base:     nil,
			Add:      []string{"a", "b", "a"},
			Expected: []string{"a", "b"},
		},
		{
			Base:     []string{"a", "b"},
			Add:      []string{"c", "b", "a"},
			Expected: []string{"a", "b", "c"},
		},
		{
			Base:     []string{"a", "b"},
			Add:      []string{"b"},
			Expected: []string{"a", "b"},
		},
	}
	for _, tc := range cases {
		appendWithoutDupes(&tc.Base, tc.Add)
		assert.Equal(t, tc.Expected, tc.Base)
	}
}

func TestListFolders(t *testing.T) {
	svc, err := mockCRMFolders()
	if err != nil {
		assert.NoError(t, err)
	}

	cases := []struct {
		BaseFolder string
		MaxDepth   int
		Expected   []string
	}{
		{
			BaseFolder: "",
			MaxDepth:   20,
			Expected:   []string{"", "root_0", "root_1", "root1.sub1_0", "root1.sub1.sub2_0", "root1.sub1.sub2_1", "root1.sub1.sub21.sub3_0", "root1.sub1.sub21.sub3_1", "root1.sub1.sub21.sub3_2", "root1.sub1.sub21.sub3_3", "root1.sub1.sub21.sub3_4", "root1.sub1.sub2_2", "root1.sub1_1", "root_2"},
		},
		{
			BaseFolder: "",
			MaxDepth:   3,
			Expected:   []string{"", "root_0", "root_1", "root1.sub1_0", "root1.sub1.sub2_0", "root1.sub1.sub2_1", "root1.sub1.sub2_2", "root1.sub1_1", "root_2"},
		},
		{
			BaseFolder: "root_1",
			MaxDepth:   1,
			Expected:   []string{"root_1", "root1.sub1_0", "root1.sub1_1"},
		},
		{
			BaseFolder: "root_1",
			MaxDepth:   2,
			Expected:   []string{"root_1", "root1.sub1_0", "root1.sub1.sub2_0", "root1.sub1.sub2_1", "root1.sub1.sub2_2", "root1.sub1_1"},
		},
	}

	for i, tc := range cases {
		ret, err := listFolders(context.Background(), hclog.NewNullLogger(), svc, tc.BaseFolder, tc.MaxDepth)
		assert.NoError(t, err)
		assert.Equalf(t, tc.Expected, ret, "Test #%d", i+1)
	}
}

func mockCRMFolders() (*cloudresourcemanager.FoldersService, error) {
	makeFolders := func(baseName string, num int) []*cloudresourcemanager.Folder {
		folders := make([]*cloudresourcemanager.Folder, num)
		for i := 0; i < num; i++ {
			var folder cloudresourcemanager.Folder
			_ = faker.FakeData(&folder)
			folder.State = "ACTIVE"
			folder.Name = baseName + "_" + strconv.Itoa(i) // we're omitting the 'folders/' prefix here to simplify the testdata
			folders[i] = &folder
		}
		return folders
	}

	mux := httprouter.New()
	mux.GET("/v3/folders", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_ = r.ParseForm()

		resp := &cloudresourcemanager.ListFoldersResponse{}

		switch strings.TrimPrefix(r.Form.Get("parent"), "folders/") {
		case "":
			resp.Folders = makeFolders("root", 3)
		case "root_1":
			resp.Folders = makeFolders("root1.sub1", 2)
		case "root1.sub1_0":
			resp.Folders = makeFolders("root1.sub1.sub2", 3)
		case "root1.sub1.sub2_1":
			resp.Folders = makeFolders("root1.sub1.sub21.sub3", 5)
		}

		b, _ := json.Marshal(resp)
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := cloudresourcemanager.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc.Folders, nil
}

func mockServiceusageService(t *testing.T, name, state string) *serviceusage.GoogleApiServiceusageV1Service {
	var service serviceusage.GoogleApiServiceusageV1Service
	if err := faker.FakeDataSkipFields(&service, []string{"Config"}); err != nil {
		t.Fatal(err)
	}
	service.State = state
	service.Config = &serviceusage.GoogleApiServiceusageV1ServiceConfig{}
	if err := faker.FakeDataSkipFields(service.Config, []string{"Documentation"}); err != nil {
		t.Fatal(err)
	}
	service.Config.Name = name
	service.Config.Documentation = &serviceusage.Documentation{}
	service.Config.Apis[0].Methods[0].Options[0].Value = []byte("{}")
	service.Config.Apis[0].Options[0].Value = []byte("{}")
	if err := faker.FakeDataSkipFields(service.Config.Documentation, []string{"Pages"}); err != nil {
		t.Fatal(err)
	}
	return &service
}

func createServices(t *testing.T) *Services {
	ctx := context.Background()
	service := mockServiceusageService(t, "service1", "ENABLED")
	mux := httprouter.New()
	mux.GET("/v1/projects/project1/services", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &serviceusage.ListServicesResponse{
			Services: []*serviceusage.GoogleApiServiceusageV1Service{service},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	ts := httptest.NewServer(mux)
	svc, err := serviceusage.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(ts.Close)
	return &Services{
		ServiceUsage: svc,
	}
}

func TestClient_configureEnabledServices(t *testing.T) {
	cl := NewGcpClient(
		logging.New(&hclog.LoggerOptions{
			Level: hclog.Warn,
		}),
		BackoffSettings{},
		[]string{"project1"},
		createServices(t),
	)
	assert.True(t, !cl.configureEnabledServices().HasErrors())
	assert.Equal(t, map[string]map[GcpService]bool{
		"project1": {"service1": true},
	}, cl.EnabledServices)
}
