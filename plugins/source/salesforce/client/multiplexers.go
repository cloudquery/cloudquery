package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var unsupportedObjects = map[string]bool{
	"AIPredictionEvent":         true,
	"AsyncOperationStatus":      true,
	"UserRecordAccess":          true,
	"AttachedContentDocument":   true,
	"AttachedContentNote":       true,
	"AggregateResult":           true,
	"AssetTokenEvent":           true,
	"ActivityHistory":           true,
	"BatchApexErrorEvent":       true,
	"AsyncOperationEvent":       true,
	"CombinedAttachment":        true,
	"ContentBody":               true,
	"DataObjectDataChgEvent":    true,
	"DatacloudAddress":          true,
	"EmailStatus":               true,
	"FeedLike":                  true,
	"FeedTrackedChange":         true,
	"FlowOrchestrationEvent":    true,
	"FolderedContentDocument":   true,
	"LogoutEventStream":         true,
	"LookedUpFromActivity":      true,
	"Name":                      true,
	"NoteAndAttachment":         true,
	"OpenActivity":              true,
	"OrgLifecycleNotification":  true,
	"OutgoingEmail":             true,
	"OutgoingEmailRelation":     true,
	"OwnedContentDocument":      true,
	"PlatformStatusAlertEvent":  true,
	"PlatformAction":            true,
	"ProcessExceptionEvent":     true,
	"ProcessInstanceHistory":    true,
	"QuoteTemplateRichTextData": true,
	"ListViewChartInstance":     true,

	// objects that needs where clause
	"AppTabMember":                true,
	"ApexTypeImplementor":         true,
	"ColorDefinition":             true,
	"ContentFolderMember":         true,
	"ContentDocumentLink":         true,
	"DataStatistics":              true,
	"DataType":                    true,
	"EntityParticle":              true,
	"FeedSignal":                  true,
	"FlexQueueItem":               true,
	"FieldDefinition":             true,
	"FlowVersionView":             true,
	"FlowTestView":                true,
	"IdeaComment":                 true,
	"IconDefinition":              true,
	"FlowExecutionErrorEvent":     true,
	"OwnerChangeOptionInfo":       true,
	"PicklistValueInfo":           true,
	"RelatedListDefinition":       true,
	"RelatedListColumnDefinition": true,
	"RelationshipDomain":          true,
	"RelationshipInfo":            true,
	"Vote":                        true,
	"UserFieldAccess":             true,
	"SiteDetail":                  true,
	"SearchLayout":                true,
	"ContentFolderItem":           true,
	"FlowVariableView":            true,
	"UserEntityAccess":            true,
}

func (c *Client) withObject(object string) *Client {
	newC := *c
	newC.Object = object
	return &newC
}

func MultiplexStandardObjects(meta schema.ClientMeta) []schema.ClientMeta {
	c := meta.(*Client)
	var clients []schema.ClientMeta
	// clients = append(clients, c.withObject("Account"))
	for _, object := range c.ListObjectsResponse.Sobject {
		if strings.HasSuffix(object.Name, "ChangeEvent") || unsupportedObjects[object.Name] {
			continue
		}
		for _, exclude := range c.spec.ExcludeObjects {
			if exclude == object.Name {
				continue
			}
		}
		if c.spec.IncludeObjects[0] == "*" {
			clients = append(clients, c.withObject(object.Name))
		} else {
			for _, include := range c.spec.IncludeObjects {
				if include == object.Name {
					clients = append(clients, c.withObject(object.Name))
				}
			}
		}
	}
	return clients
}
