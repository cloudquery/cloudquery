package services

import (
	"context"

	"github.com/hermanschaaf/hackernews"
)

//go:generate mockgen -package=mocks -destination=../mocks/hackernews.go -source=hackernews.go HackernewsClient
type HackernewsClient interface {
	AskStories(context.Context) ([]int, error)
	BestStories(context.Context) ([]int, error)
	GetItem(context.Context, int) (hackernews.Item, error)
	GetUser(context.Context, string) (hackernews.User, error)
	JobStories(context.Context) ([]int, error)
	MaxItemID(context.Context) (int, error)
	NewStories(context.Context) ([]int, error)
	ShowStories(context.Context) ([]int, error)
	TopStories(context.Context) ([]int, error)
	Updates(context.Context) (*hackernews.Updates, error)
}
