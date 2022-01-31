package drift

import (
	"fmt"

	"github.com/cloudquery/cloudquery/internal/s3"
	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
)

func loadIACStatesFromS3(iacID, bucket string, keys []string, region, roleARN string) (interface{}, error) {
	svc, err := s3.Session(bucket, region, roleARN)
	if err != nil {
		return nil, err
	}

	ret := make([]*terraform.Data, 0, len(keys))
	for _, globKey := range keys {
		matches, err := s3.Glob(svc, bucket, globKey)
		if err != nil {
			return nil, err
		}
		for _, key := range matches {
			body, err := s3.GetObject(svc, bucket, key)
			if err != nil {
				return nil, err
			}
			data, err := terraform.LoadState(body)
			_ = body.Close()
			if err != nil {
				return nil, fmt.Errorf("parse s3://%s/%s: %w", bucket, key, err)
			}

			ret = append(ret, data)
		}
	}

	if len(ret) == 0 {
		return nil, fmt.Errorf("no matches for specified %s state patterns", iacID)
	}

	return ret, nil
}
