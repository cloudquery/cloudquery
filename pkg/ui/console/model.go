package console

import (
	"fmt"
)

// ModuleCallRequest is the request used to call a module.
type ModuleCallRequest struct {
	// Name of the module
	Name string

	// Params are the invocation parameters specific to the module
	Params interface{}

	// Profile is the selected/overridden name of the profile
	Profile string

	// OutputPath is the filename to save output to
	OutputPath string
}

type ExitCodeError struct {
	OriginalError error
	ExitCode      int
}

func (e *ExitCodeError) Error() string {
	// BEWARE: uncommenting this to fix a panic means that we're reporting this error to sentry/telemetry when we probably shouldn't have
	//if e.OriginalError == nil {
	//	return fmt.Sprintf("exit code %d", e.ExitCode)
	//}

	return fmt.Sprintf("exit code %d. err: %s", e.ExitCode, e.OriginalError.Error())
}
