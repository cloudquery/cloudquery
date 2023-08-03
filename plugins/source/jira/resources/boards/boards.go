package boards

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Boards() *schema.Table {
	return &schema.Table{
		Name:      "jira_boards",
		Transform: transformers.TransformWithStruct(&jira.Board{}, transformers.WithPrimaryKeys("Self")),
		Resolver:  fetchBoards,
	}
}

func fetchBoards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	boardList, _, err := c.Jira.Board.GetAllBoardsWithContext(ctx, nil)
	if err != nil {
		return err
	}
	res <- boardList.Values
	return nil
}
