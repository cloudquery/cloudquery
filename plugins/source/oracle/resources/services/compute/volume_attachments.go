package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

type VolumeAttachment struct {
	// Both core.ParavirtualizedVolumeAttachment & core.EmulatedVolumeAttachment
	// have the same fields (listed here).
	// We fill them in based on the core.VolumeAttachment interface.
	AvailabilityDomain             *string
	CompartmentId                  *string
	Id                             *string
	InstanceId                     *string
	TimeCreated                    *common.SDKTime
	VolumeId                       *string
	Device                         *string
	DisplayName                    *string
	IsReadOnly                     *bool
	IsShareable                    *bool
	IsPvEncryptionInTransitEnabled *bool
	IsMultipath                    *bool
	LifecycleState                 core.VolumeAttachmentLifecycleStateEnum
	IscsiLoginState                core.VolumeAttachmentIscsiLoginStateEnum

	// This will be filled in and used only if the underlying attachment is core.IScsiVolumeAttachment.
	// It gives us extra fields to report, too.
	*core.IScsiVolumeAttachment
}

func VolumeAttachments() *schema.Table {
	return &schema.Table{
		Name:                "oracle_compute_volume_attachments",
		Resolver:            fetchVolumeAttachments,
		PreResourceResolver: unwrapVolumeAttachment,
		Multiplex:           client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(new(VolumeAttachment),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
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

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}

func unwrapVolumeAttachment(_ context.Context, _ schema.ClientMeta, r *schema.Resource) error {
	item := r.Item.(core.VolumeAttachment)
	res := &VolumeAttachment{
		AvailabilityDomain:             item.GetAvailabilityDomain(),
		CompartmentId:                  item.GetCompartmentId(),
		Id:                             item.GetId(),
		InstanceId:                     item.GetInstanceId(),
		LifecycleState:                 item.GetLifecycleState(),
		TimeCreated:                    item.GetTimeCreated(),
		VolumeId:                       item.GetVolumeId(),
		Device:                         item.GetDevice(),
		DisplayName:                    item.GetDisplayName(),
		IsReadOnly:                     item.GetIsReadOnly(),
		IsShareable:                    item.GetIsShareable(),
		IsPvEncryptionInTransitEnabled: item.GetIsPvEncryptionInTransitEnabled(),
		IsMultipath:                    item.GetIsMultipath(),
		IscsiLoginState:                item.GetIscsiLoginState(),
	}

	if iSCSI, ok := item.(*core.IScsiVolumeAttachment); ok {
		res.IScsiVolumeAttachment = iSCSI
	}
	r.SetItem(res)
	return nil
}
