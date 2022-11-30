package users

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.GitlabServices {
	userMock := mocks.NewMockUsersClient(ctrl)

	var user *gitlab.User
	if err := faker.FakeObject(&user, faker.WithMaxDepth(25)); err != nil {
		t.Fatal(err)
	}

	userMock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(user, &gitlab.Response{}, nil)
	// GetUser(user int, opt gitlab.GetUsersOptions, options ...gitlab.RequestOptionFunc) (*gitlab.User, *gitlab.Response, error)

	groupMock := mocks.NewMockGroupsClient(ctrl)

	var group *gitlab.Group
	if err := faker.FakeObject(&group, faker.WithMaxDepth(10)); err != nil {
		t.Fatal(err)
	}
	groupMock.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Return([]*gitlab.Group{group}, &gitlab.Response{}, nil)

	var groupMember *gitlab.GroupMember
	if err := faker.FakeObject(&groupMember, faker.WithMaxDepth(12)); err != nil {
		t.Fatal(err)
	}
	groupMock.EXPECT().ListGroupMembers(gomock.Any(), gomock.Any()).Return([]*gitlab.GroupMember{groupMember}, &gitlab.Response{}, nil)
	return client.GitlabServices{
		Users:  userMock,
		Groups: groupMock,
	}
}

func TestStorageBillings(t *testing.T) {
	client.GitlabMockTestHelper(t, Groups(), buildUsers, client.TestOptions{})
}

// Nov 29 20:38:52.000000 ERR column resolver finished with panic error={} client= column=projects duration=1.926292 module=gitlab_users_groups-src stack="runtime error: index out of range [123] with length 6\ngoroutine 9 [running]:\nruntime/debug.Stack()\n\t/Users/benbernays/sdk/go1.19.3/src/runtime/debug/stack.go:24 +0x64\ngithub.com/cloudquery/plugin-sdk/plugins.(*SourcePlugin).resolveColumn.func1()\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/plugins/source_scheduler_dfs.go:206 +0x70\npanic({0x100be1760, 0x14000025878})\n\t/Users/benbernays/sdk/go1.19.3/src/runtime/panic.go:890 +0x26c\nencoding/json.(*encodeState).marshal.func1()\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:327 +0x108\npanic({0x100be1760, 0x14000025878})\n\t/Users/benbernays/sdk/go1.19.3/src/runtime/panic.go:890 +0x26c\ngithub.com/xanzy/go-gitlab.NotificationLevelValue.String(0x7b)\n\t/Users/benbernays/go/pkg/mod/github.com/xanzy/go-gitlab@v0.74.0/types.go:499 +0x68\ngithub.com/xanzy/go-gitlab.NotificationLevelValue.MarshalJSON(0x7b)\n\t/Users/benbernays/go/pkg/mod/github.com/xanzy/go-gitlab@v0.74.0/types.go:504 +0x34\nencoding/json.addrMarshalerEncoder(0x140000bf400, {0x100b9fd60, 0x14000027438, 0x182}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:495 +0xf8\nencoding/json.condAddrEncoder.encode({0x100c0a608, 0x100c0a658}, 0x140000bf400, {0x100b9fd60, 0x14000027438, 0x182}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:959 +0x78\nencoding/json.structEncoder.encode({{{0x14000001560, 0x2, 0x2}, 0x14000192960}}, 0x140000bf400, {0x100bbb680, 0x14000027430, 0x199}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:760 +0x308\nencoding/json.ptrEncoder.encode({0x14000192990}, 0x140000bf400, {0x100b71040, 0x1400006dd10, 0x196}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:944 +0x3b0\nencoding/json.structEncoder.encode({{{0x14000001440, 0x2, 0x2}, 0x14000192a80}}, 0x140000bf400, {0x100bbb5e0, 0x1400006dd10, 0x199}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:760 +0x308\nencoding/json.ptrEncoder.encode({0x14000192ab0}, 0x140000bf400, {0x100b71000, 0x140000f6610, 0x196}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:944 +0x3b0\nencoding/json.structEncoder.encode({{{0x140001e0000, 0x6a, 0x8e}, 0x14000192f00}}, 0x140000bf400, {0x100c099e0, 0x140000f6480, 0x199}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:760 +0x308\nencoding/json.ptrEncoder.encode({0x14000192f30}, 0x140000bf400, {0x100b969c0, 0x14000076a20, 0x196}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:944 +0x3b0\nencoding/json.arrayEncoder.encode({0x140001bebf0}, 0x140000bf400, {0x100b766a0, 0x140001abb60, 0x97}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:915 +0xf8\nencoding/json.sliceEncoder.encode({0x140001bec10}, 0x140000bf400, {0x100b766a0, 0x140001abb60, 0x97}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:888 +0x434\nencoding/json.(*encodeState).reflectValue(0x140000bf400, {0x100b766a0, 0x140001abb60, 0x97}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:359 +0x78\nencoding/json.(*encodeState).marshal(0x140000bf400, {0x100b766a0, 0x140001abb60}, {0x0, 0x1})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:331 +0xb8\nencoding/json.Marshal({0x100b766a0, 0x140001abb60})\n\t/Users/benbernays/sdk/go1.19.3/src/encoding/json/encode.go:160 +0x5c\ngithub.com/cloudquery/plugin-sdk/schema.(*JSON).Set(0x140001b4ca0, {0x100b766a0, 0x140001abb60})\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/schema/json.go:75 +0x618\ngithub.com/cloudquery/plugin-sdk/schema.(*Resource).Set(0x14000035f00, {0x100ae170c, 0x8}, {0x100b766a0, 0x140001abb60})\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/schema/resource.go:73 +0xbc\ngithub.com/cloudquery/plugin-sdk/schema.PathResolver.func1({0x4a, 0x1f4}, {0x100c0cd60, 0x140000c3540}, 0x14000035f00, {{0x100ae170c, 0x8}, 0xa, {0x0, 0x0}, ...})\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/schema/resolvers.go:17 +0xec\ngithub.com/cloudquery/plugin-sdk/plugins.(*SourcePlugin).resolveColumn(0x1400018c000, {0x100c10388, 0x140000957a0}, {{0x100c0eed0, 0x1400006d500}, 0x0, {0x0, 0x0}, {0x1400018aa00, 0x4d, ...}, ...}, ...)\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/plugins/source_scheduler_dfs.go:213 +0x1a4\ngithub.com/cloudquery/plugin-sdk/plugins.(*SourcePlugin).resolveResource(0x1400018c000, {0x100c10388, 0x140000957a0}, 0x140000bec80, {0x100c0cd60, 0x140000c3540}, 0x0, {0x100b70d80, 0x140000faa00})\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/plugins/source_scheduler_dfs.go:189 +0x784\ngithub.com/cloudquery/plugin-sdk/plugins.(*SourcePlugin).resolveResourcesDfs.func1.1()\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/plugins/source_scheduler_dfs.go:128 +0x170\ncreated by github.com/cloudquery/plugin-sdk/plugins.(*SourcePlugin).resolveResourcesDfs.func1\n\t/Users/benbernays/Documents/GitHub/plugin-sdk/plugins/source_scheduler_dfs.go:124 +0x364\n" table=gitlab_users_groups
