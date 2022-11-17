// Code generated by codegen; DO NOT EDIT.
package policy

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

type PolicyClient struct {
	AssignmentsClient AssignmentsClient
}

func NewPolicyClient(subscriptionID string, credentials azcore.TokenCredential, options *arm.ClientOptions) (*PolicyClient, error) {
	var client PolicyClient
	var err error

	client.AssignmentsClient, err = armpolicy.NewAssignmentsClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
