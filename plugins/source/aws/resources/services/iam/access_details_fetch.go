package iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
)

func fetchIamAccessDetails(ctx context.Context, res chan<- interface{}, svc services.IamClient, arn string) error {
	config := iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         &arn,
		Granularity: types.AccessAdvisorUsageGranularityTypeActionLevel,
	}
	output, err := svc.GenerateServiceLastAccessedDetails(ctx, &config)
	if err != nil {
		return err
	}

	getDetails := iam.GetServiceLastAccessedDetailsInput{
		JobId: output.JobId,
	}
	for {
		details, err := svc.GetServiceLastAccessedDetails(ctx, &getDetails)
		if err != nil {
			return err
		}

		switch details.JobStatus {
		case types.JobStatusTypeInProgress:
			time.Sleep(time.Second)
			continue
		case types.JobStatusTypeFailed:
			return fmt.Errorf("failed to get last accessed details with error: %s - %s", *details.Error.Code, *details.Error.Message)
		case types.JobStatusTypeCompleted:
			for _, s := range details.ServicesLastAccessed {
				if err := fetchDetailEntities(ctx, res, svc, s, *output.JobId, arn); err != nil {
					return err
				}
			}
			if details.Marker == nil {
				return nil
			}
			if details.Marker != nil {
				getDetails.Marker = details.Marker
			}
		}
	}
}

func fetchDetailEntities(ctx context.Context, res chan<- interface{}, svc services.IamClient, sla types.ServiceLastAccessed, jobId string, arn string) error {
	config := iam.GetServiceLastAccessedDetailsWithEntitiesInput{
		JobId:            &jobId,
		ServiceNamespace: sla.ServiceNamespace,
		MaxItems:         aws.Int32(1000),
	}
	details := models.ServiceLastAccessedEntitiesWrapper{
		ResourceARN:         arn,
		JobId:               &jobId,
		ServiceLastAccessed: &sla,
	}
	for {
		output, err := svc.GetServiceLastAccessedDetailsWithEntities(ctx, &config)
		if err != nil {
			return err
		}
		details.Entities = append(details.Entities, output.EntityDetailsList...)
		if output.Marker == nil {
			break
		}
		if output.Marker != nil {
			config.Marker = output.Marker
		}
	}
	res <- details
	return nil
}
