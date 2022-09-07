package organizations

import (
	"github.com/google/go-github/v45/github"
)

type Member struct {
	*github.User
	Membership *github.Membership
}
