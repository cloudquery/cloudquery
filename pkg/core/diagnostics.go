package core

import (
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/viper"
)

type SentryDiagnostic struct {
	diag.Diagnostic

	Tags   map[string]string
	Ignore bool
}

func (d *SentryDiagnostic) Redacted() diag.Diagnostic {
	v, ok := d.Diagnostic.(diag.Redactable)
	if !ok {
		return d
	}
	if r := v.Redacted(); r != nil {
		return &SentryDiagnostic{
			Diagnostic: r,
			Tags:       d.Tags,
			Ignore:     d.Ignore,
		}
	}
	return d
}

func (d *SentryDiagnostic) IsSentryDiagnostic() (bool, map[string]string, bool) {
	return true, d.Tags, d.Ignore
}

func convertToConfigureDiags(diags diag.Diagnostics) diag.Diagnostics {
	return convertToSentryDiags(diags, func(d diag.Diagnostic) *SentryDiagnostic {
		return &SentryDiagnostic{
			Diagnostic: d,
			Tags:       map[string]string{"source": "configure"},
			Ignore:     d.Type() == diag.ACCESS,
		}
	})
}

func convertToFetchDiags(diags diag.Diagnostics, providerName, providerVersion string) diag.Diagnostics {
	allowUnmanaged := viper.GetBool("debug-sentry")

	return convertToSentryDiags(diags, func(d diag.Diagnostic) *SentryDiagnostic {
		return &SentryDiagnostic{
			Diagnostic: d,
			Tags: map[string]string{
				"provider":         providerName,
				"provider_version": providerVersion,
				"resource":         d.Description().Resource,
			},
			Ignore: !allowUnmanaged && providerVersion == plugin.Unmanaged,
		}
	})
}

// convertToSentryDiags gets the diags and applies the given handleFunc to each diag which converts them to a Sentry Diagnostic.
func convertToSentryDiags(diags diag.Diagnostics, handleFunc func(diag.Diagnostic) *SentryDiagnostic) diag.Diagnostics {
	fd := make(diag.Diagnostics, 0, len(diags))
	for i := range diags {
		sd := handleFunc(diags[i])
		if sd == nil {
			continue
		}
		fd = append(fd, sd)
	}
	return fd
}
