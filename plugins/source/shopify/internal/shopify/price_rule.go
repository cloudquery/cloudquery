package shopify

import "time"

type PriceRule struct {
	ID                                     int64      `json:"id"`
	ValueType                              string     `json:"value_type"`
	Value                                  string     `json:"value"`
	CustomerSelection                      string     `json:"customer_selection"`
	TargetType                             string     `json:"target_type"`
	TargetSelection                        string     `json:"target_selection"`
	AllocationMethod                       string     `json:"allocation_method"`
	AllocationLimit                        any        `json:"allocation_limit"`
	OncePerCustomer                        bool       `json:"once_per_customer"`
	UsageLimit                             int        `json:"usage_limit"`
	StartsAt                               *time.Time `json:"starts_at"`
	EndsAt                                 *time.Time `json:"ends_at"`
	CreatedAt                              time.Time  `json:"created_at"`
	UpdatedAt                              time.Time  `json:"updated_at"`
	EntitledProductIds                     []any      `json:"entitled_product_ids"`
	EntitledVariantIds                     []any      `json:"entitled_variant_ids"`
	EntitledCollectionIds                  []any      `json:"entitled_collection_ids"`
	EntitledCountryIds                     []any      `json:"entitled_country_ids"`
	PrerequisiteProductIds                 []any      `json:"prerequisite_product_ids"`
	PrerequisiteVariantIds                 []any      `json:"prerequisite_variant_ids"`
	PrerequisiteCollectionIds              []any      `json:"prerequisite_collection_ids"`
	CustomerSegmentPrerequisiteIds         []any      `json:"customer_segment_prerequisite_ids"`
	PrerequisiteCustomerIds                []any      `json:"prerequisite_customer_ids"`
	PrerequisiteSubtotalRange              any        `json:"prerequisite_subtotal_range"`
	PrerequisiteQuantityRange              any        `json:"prerequisite_quantity_range"`
	PrerequisiteShippingPriceRange         any        `json:"prerequisite_shipping_price_range"`
	PrerequisiteToEntitlementQuantityRatio struct {
		PrerequisiteQuantity *int64 `json:"prerequisite_quantity"`
		EntitledQuantity     *int64 `json:"entitled_quantity"`
	} `json:"prerequisite_to_entitlement_quantity_ratio"`
	PrerequisiteToEntitlementPurchase struct {
		PrerequisiteAmount any `json:"prerequisite_amount"`
	} `json:"prerequisite_to_entitlement_purchase"`
	Title             string `json:"title"`
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
}

type GetPriceRulesResponse struct {
	PriceRules []PriceRule `json:"price_rules"`
	PageSize   int         `json:"page_size"`
}

type DiscountCode struct {
	ID          int64     `json:"id"`
	PriceRuleID int64     `json:"price_rule_id"`
	Code        string    `json:"code"`
	UsageCount  int       `json:"usage_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetDiscountCodesResponse struct {
	DiscountCodes []DiscountCode `json:"discount_codes"`
	PageSize      int            `json:"page_size"`
}
