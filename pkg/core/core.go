package core

import (
	"strings"

	"github.com/cloudquery/cloudquery/pkg/configv2"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
)

func ParseProviderSource(requestedProvider *configv2.RequiredProvider) (string, string, error) {
	var requestedSource string
	if requestedProvider.Source == nil || *requestedProvider.Source == "" {
		requestedSource = requestedProvider.Name
	} else {
		requestedSource = *requestedProvider.Source
		if !strings.Contains(requestedSource, "/") {
			requestedSource = strings.Join([]string{requestedSource, requestedProvider.Name}, "/")
		}
	}
	return registry.ParseProviderName(requestedSource)
}
