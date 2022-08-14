package client

import (
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	return classifyError(err, diag.RESOLVING, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, opts ...diag.BaseErrorOption) diag.Diagnostics {
	if strings.Contains(err.Error(), `your Okta domain should not contain -admin.`) {
		return diag.Diagnostics{
			diag.NewBaseError(err, diag.USER, opts...),
		}
	}

	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			diag.NewBaseError(d, d.Type(), opts...),
		}
	}

	return diag.Diagnostics{
		diag.NewBaseError(err, fallbackType, opts...),
	}
}
