package console

import (
	"fmt"
	"strconv"
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
	return fmt.Sprintf("exit code %s. err: %s", strconv.Itoa(e.ExitCode), e.OriginalError.Error())
}
