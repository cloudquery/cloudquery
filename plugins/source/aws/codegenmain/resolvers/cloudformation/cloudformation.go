package cloudformation

import "regexp"

var (
	ValidStackNotFoundRegex = regexp.MustCompile("Stack with id (.*) does not exist")
)
