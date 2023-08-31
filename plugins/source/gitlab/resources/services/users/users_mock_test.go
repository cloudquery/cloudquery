package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/xanzy/go-gitlab"
)

func buildUsers(mux *httprouter.Router) error {
	var user *gitlab.User
	if err := faker.FakeObject(&user, faker.WithMaxDepth(25)); err != nil {
		return err
	}

	isoTime := gitlab.ISOTime(time.Now())
	user.LastActivityOn = &isoTime
	userResp, err := json.Marshal([]*gitlab.User{user})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(userResp))
	},
	)

	return nil
}

func TestGroups(t *testing.T) {
	client.GitlabMockTestHelper(t, Users(), buildUsers, client.TestOptions{})
}
