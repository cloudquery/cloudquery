package cdn

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func marshalConditions(conditions []cdn.BasicDeliveryRuleCondition) ([]byte, error) {
	data := make([]interface{}, 0, len(conditions))
	for _, c := range conditions {
		switch c.(type) {
		case cdn.DeliveryRuleRemoteAddressCondition:
			w, _ := c.AsDeliveryRuleRemoteAddressCondition()
			data = append(data, w)
		case cdn.DeliveryRuleRequestMethodCondition:
			w, _ := c.AsDeliveryRuleRequestMethodCondition()
			data = append(data, w)
		case cdn.DeliveryRuleQueryStringCondition:
			w, _ := c.AsDeliveryRuleQueryStringCondition()
			data = append(data, w)
		case cdn.DeliveryRulePostArgsCondition:
			w, _ := c.AsDeliveryRulePostArgsCondition()
			data = append(data, w)
		case cdn.DeliveryRuleRequestURICondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleRequestHeaderCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleRequestBodyCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleRequestSchemeCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleURLPathCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleURLFileExtensionCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleURLFileNameCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleHTTPVersionCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleCookiesCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleIsDeviceCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		case cdn.DeliveryRuleCondition:
			w, _ := c.AsDeliveryRuleCondition()
			data = append(data, w)
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func marshalActions(actions []cdn.BasicDeliveryRuleAction) ([]byte, error) {
	data := make([]interface{}, 0, len(actions))
	for _, c := range actions {
		switch c.(type) {
		case cdn.URLRedirectAction:
			w, _ := c.AsURLRedirectAction()
			data = append(data, w)
		case cdn.URLSigningAction:
			w, _ := c.AsURLSigningAction()
			data = append(data, w)
		case cdn.URLRewriteAction:
			w, _ := c.AsURLRewriteAction()
			data = append(data, w)
		case cdn.OriginGroupOverrideAction:
			w, _ := c.AsOriginGroupOverrideAction()
			data = append(data, w)
		case cdn.DeliveryRuleResponseHeaderAction:
			w, _ := c.AsDeliveryRuleResponseHeaderAction()
			data = append(data, w)
		case cdn.DeliveryRuleCacheExpirationAction:
			w, _ := c.AsDeliveryRuleCacheExpirationAction()
			data = append(data, w)
		case cdn.DeliveryRuleCacheKeyQueryStringAction:
			w, _ := c.AsDeliveryRuleCacheKeyQueryStringAction()
			data = append(data, w)
		case cdn.DeliveryRuleRequestHeaderAction:
			w, _ := c.AsDeliveryRuleRequestHeaderAction()
			data = append(data, w)
		case cdn.DeliveryRuleAction:
			w, _ := c.AsDeliveryRuleAction()
			data = append(data, w)
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}
