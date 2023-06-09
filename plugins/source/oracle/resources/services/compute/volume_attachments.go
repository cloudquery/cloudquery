package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func VolumeAttachments() *schema.Table {
	// We base this resource on core.IScsiVolumeAttachment.
	// The other supported type core.ParavirtualizedVolumeAttachment has just a subset of fields available.
	// We try to cast the result to *core.IScsiVolumeAttachment.
	// If unsuccessful, we just fill in the common fields using core.VolumeAttachment interface.
	return &schema.Table{
		Name:      "oracle_compute_volume_attachments",
		Resolver:  fetchVolumeAttachments,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(new(core.IScsiVolumeAttachment)),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchVolumeAttachments(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := core.ListVolumeAttachmentsRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].CoreComputeClient.ListVolumeAttachments(ctx, request)

		if err != nil {
			return err
		}

		res <- unwrapVolumeAttachments(response.Items)

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}

func unwrapVolumeAttachments(attachments []core.VolumeAttachment) []*core.IScsiVolumeAttachment {
	if attachments == nil {
		return nil
	}

	result := make([]*core.IScsiVolumeAttachment, len(attachments))
	for i, va := range attachments {
		switch va := va.(type) {
		case core.IScsiVolumeAttachment:
			result[i] = &va
		case *core.IScsiVolumeAttachment:
			result[i] = va
		default:
			result[i] = &core.IScsiVolumeAttachment{
				AvailabilityDomain:             va.GetAvailabilityDomain(),
				CompartmentId:                  va.GetCompartmentId(),
				Id:                             va.GetId(),
				InstanceId:                     va.GetInstanceId(),
				TimeCreated:                    va.GetTimeCreated(),
				VolumeId:                       va.GetVolumeId(),
				Device:                         va.GetDevice(),
				DisplayName:                    va.GetDisplayName(),
				IsReadOnly:                     va.GetIsReadOnly(),
				IsShareable:                    va.GetIsShareable(),
				IsPvEncryptionInTransitEnabled: va.GetIsPvEncryptionInTransitEnabled(),
				IsMultipath:                    va.GetIsMultipath(),
				LifecycleState:                 va.GetLifecycleState(),
				IscsiLoginState:                va.GetIscsiLoginState(),
			}
		}
	}
	return result
}
