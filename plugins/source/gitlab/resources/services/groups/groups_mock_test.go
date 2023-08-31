package groups

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

func buildGroups(mux *httprouter.Router) error {
	var group *gitlab.Group
	if err := faker.FakeObject(&group, faker.WithMaxDepth(10)); err != nil {
		return err
	}

	isoTime := gitlab.ISOTime(time.Now())
	group.MarkedForDeletionOn = &isoTime

	groupResp, err := json.Marshal([]*gitlab.Group{group})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/groups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(groupResp))
	})

	var groupMember *gitlab.GroupMember
	if err := faker.FakeObject(&groupMember, faker.WithMaxDepth(12)); err != nil {
		return err
	}

	groupMember.ExpiresAt = &isoTime
	groupMembers, err := json.Marshal([]*gitlab.GroupMember{groupMember})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/groups/:group/members", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(groupMembers))
	})

	var billableGroupMember *gitlab.BillableGroupMember
	if err := faker.FakeObject(&billableGroupMember, faker.WithMaxDepth(12)); err != nil {
		return err
	}

	billableGroupMember.LastActivityOn = &isoTime
	billableGroupMembers, err := json.Marshal([]*gitlab.BillableGroupMember{billableGroupMember})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/groups/:group/billable_members", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(billableGroupMembers))
	})

	return nil
}

func TestGroups(t *testing.T) {
	client.GitlabMockTestHelper(t, Groups(), buildGroups, client.TestOptions{})
}
