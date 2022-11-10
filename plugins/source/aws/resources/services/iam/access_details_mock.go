package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func iamAccessDetailsMock(t *testing.T, m *mocks.MockIamClient) {
	serviceAccessReport := iam.GenerateServiceLastAccessedDetailsOutput{}
	err := faker.FakeObject(&serviceAccessReport)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&serviceAccessReport, nil)

	serviceAccessDetails := iam.GetServiceLastAccessedDetailsOutput{}
	err = faker.FakeObject(&serviceAccessDetails)
	if err != nil {
		t.Fatal(err)
	}
	serviceAccessDetails.Marker = nil
	serviceAccessDetails.JobStatus = iamTypes.JobStatusTypeCompleted
	m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&serviceAccessDetails, nil)

	serviceAccessDetailsEntities := iam.GetServiceLastAccessedDetailsWithEntitiesOutput{}
	err = faker.FakeObject(&serviceAccessDetailsEntities)
	if err != nil {
		t.Fatal(err)
	}
	serviceAccessDetailsEntities.Marker = nil
	m.EXPECT().GetServiceLastAccessedDetailsWithEntities(gomock.Any(), gomock.Any(), gomock.Any()).Return(&serviceAccessDetailsEntities, nil)
}
