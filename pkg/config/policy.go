package config

type Policy struct {
	Views   []PolicyView
	Queries []PolicyQueries
}

type PolicyView struct {
	Name  string
	Query string
}

type PolicyQueries struct {
	Name   string
	Invert bool
	Query  string
}
