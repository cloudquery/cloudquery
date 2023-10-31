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
		Name:        "jira_boards",
		Description: "This table shows data for Jira Boards.",
		Transform:   transformers.TransformWithStruct(&jira.Board{}, transformers.WithPrimaryKeys("Self")),
		Resolver:    fetchBoards,
	}
}

func fetchBoards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	startAt := 0
	for {
		boardList, resp, err := c.Jira.Board.GetAllBoardsWithContext(ctx, &jira.BoardListOptions{
			SearchOptions: jira.SearchOptions{
				StartAt:    startAt,
				MaxResults: 1000,
			},
		})
		if err != nil {
			return err
		}
		res <- boardList.Values

		if resp.Total <= resp.StartAt+resp.MaxResults {
			break
		}
		startAt = resp.StartAt + resp.MaxResults
	}
	return nil
}
