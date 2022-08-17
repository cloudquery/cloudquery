package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
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
		ret, err := listFolders(context.Background(), svc, tc.BaseFolder, tc.MaxDepth)
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
