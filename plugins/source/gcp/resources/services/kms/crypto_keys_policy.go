package kms

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudkms/v1"
)

func resolveKmsKeyringCryptoKeyPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client_ := meta.(*client.Client)
	p := resource.Item.(*cloudkms.CryptoKey)
	resp, err := client_.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(p.Name).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	var policy map[string]interface{}
	data, err := json.Marshal(resp)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, policy))
}
