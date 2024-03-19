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
	provider common.ConfigurationProvider
}

func (r rawPrivateKeyConfigProvider) AuthType() (common.AuthConfig, error) {
	return r.provider.AuthType()
}

func (r rawPrivateKeyConfigProvider) KeyFingerprint() (string, error) {
	return r.provider.KeyFingerprint()
}

func (r rawPrivateKeyConfigProvider) KeyID() (string, error) {
	return r.provider.KeyID()
}

func (r rawPrivateKeyConfigProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	tfVarEnvironmentVariable := fmt.Sprintf("%s_%s", tfVarEnvironmentVariable, "private_key")
	ocCLIEnvironmentVariable := fmt.Sprintf("%s_%s", ocCLIEnvironmentVariable, "private_key")
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

func (r rawPrivateKeyConfigProvider) Region() (string, error) {
	return r.provider.Region()
}

func (r rawPrivateKeyConfigProvider) TenancyOCID() (string, error) {
	return r.provider.TenancyOCID()
}

func (r rawPrivateKeyConfigProvider) UserOCID() (string, error) {
	return r.provider.UserOCID()
}

func newRawPrivateKeyConfigProvider() rawPrivateKeyConfigProvider {
	tfVarEnvProvider := common.ConfigurationProviderEnvironmentVariables(tfVarEnvironmentVariable, "")
	ociCLIEnvProvider := common.ConfigurationProviderEnvironmentVariables(ocCLIEnvironmentVariable, "")
	// ComposingConfigurationProvider fails on empty provider list or nil provider so we can safely ignore the error
	provider, _ := common.ComposingConfigurationProvider([]common.ConfigurationProvider{tfVarEnvProvider, ociCLIEnvProvider})
	return rawPrivateKeyConfigProvider{
		provider: provider,
	}
}
