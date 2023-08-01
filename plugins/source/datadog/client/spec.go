package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	Accounts    []Account          `json:"accounts"`
	Concurrency int                `json:"concurrency"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

type Account struct {
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
	AppKey string `json:"app_key"`
	APIUrl string `json:"api_url,omitempty"`
}
