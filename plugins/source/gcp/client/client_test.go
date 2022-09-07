package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

func mockCRMFolders() (*cloudresourcemanager.FoldersService, error) {
	makeFolders := func(baseName string, num int) []*cloudresourcemanager.Folder {
		folders := make([]*cloudresourcemanager.Folder, num)
		for i := 0; i < num; i++ {
			var folder cloudresourcemanager.Folder
			_ = faker.FakeObject(&folder)
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
