package client

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/oracle/oci-go-sdk/v65/common"
)

const (
	tfVarEnvironmentVariable = "TF_VAR"
	ocCLIEnvironmentVariable = "OCI_CLI"
)

type rawPrivateKeyConfigProvider struct {
	common.ConfigurationProvider
}

func (rawPrivateKeyConfigProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	const (
		tfVarEnvironmentVariable = tfVarEnvironmentVariable + "_private_key"
		ocCLIEnvironmentVariable = ocCLIEnvironmentVariable + "_private_key"
	)
	var envName string
	var ok bool
	var value string
	var err error
	envsToTry := []string{tfVarEnvironmentVariable, ocCLIEnvironmentVariable}
	for _, env := range envsToTry {
		if value, ok = os.LookupEnv(env); ok {
			envName = env
			break
		}
	}
	if !ok {
		err = errors.Join(err, fmt.Errorf("can not read PrivateKey from env variable: %s", tfVarEnvironmentVariable))
		err = errors.Join(err, fmt.Errorf("can not read PrivateKey from env variable: %s", ocCLIEnvironmentVariable))
		return nil, err
	}
	withLineBreaksNormalized := strings.ReplaceAll(value, "\\n", "\n")
	privateKey, err := common.PrivateKeyFromBytesWithPassword([]byte(withLineBreaksNormalized), nil)
	if err != nil {
		return nil, fmt.Errorf("can not parse PrivateKey from env variable: %s", envName)
	}

	return privateKey, nil
}

func newRawPrivateKeyConfigProvider() rawPrivateKeyConfigProvider {
	tfVarEnvProvider := common.ConfigurationProviderEnvironmentVariables(tfVarEnvironmentVariable, "")
	ociCLIEnvProvider := common.ConfigurationProviderEnvironmentVariables(ocCLIEnvironmentVariable, "")
	// ComposingConfigurationProvider fails on empty provider list or nil provider so we can safely ignore the error
	provider, _ := common.ComposingConfigurationProvider([]common.ConfigurationProvider{tfVarEnvProvider, ociCLIEnvProvider})
	return rawPrivateKeyConfigProvider{
		ConfigurationProvider: provider,
	}
}
