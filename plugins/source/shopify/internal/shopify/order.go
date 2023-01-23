package shopify

import "time"

type Order struct {
	ID                       int64      `json:"id"`
	AdminGraphqlAPIID        string     `json:"admin_graphql_api_id"`
	AppID                    int64      `json:"app_id"`
	BrowserIP                string     `json:"browser_ip"`
	BuyerAcceptsMarketing    bool       `json:"buyer_accepts_marketing"`
	CancelReason             any        `json:"cancel_reason"`
	CancelledAt              *time.Time `json:"cancelled_at"`
	CartToken                string     `json:"cart_token"`
	CheckoutID               int64      `json:"checkout_id"`
	CheckoutToken            string     `json:"checkout_token"`
	ClientDetails            any        `json:"client_details"`
	ClosedAt                 *time.Time `json:"closed_at"`
	Confirmed                bool       `json:"confirmed"`
	ContactEmail             string     `json:"contact_email"`
	CreatedAt                time.Time  `json:"created_at"`
	Currency                 string     `json:"currency"`
	CurrentSubtotalPrice     string     `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet  any        `json:"current_subtotal_price_set"`
	CurrentTotalDiscounts    string     `json:"current_total_discounts"`
	CurrentTotalDiscountsSet any        `json:"current_total_discounts_set"`
	CurrentTotalDutiesSet    any        `json:"current_total_duties_set"`
	CurrentTotalPrice        string     `json:"current_total_price"`
	CurrentTotalPriceSet     any        `json:"current_total_price_set"`
	CurrentTotalTax          string     `json:"current_total_tax"`
	CurrentTotalTaxSet       any        `json:"current_total_tax_set"`
	CustomerLocale           string     `json:"customer_locale"`
	DeviceID                 any        `json:"device_id"`
	DiscountCodes            []any      `json:"discount_codes"`
	Email                    string     `json:"email"`
	EstimatedTaxes           bool       `json:"estimated_taxes"`
	FinancialStatus          string     `json:"financial_status"`
	FulfillmentStatus        any        `json:"fulfillment_status"`
	Gateway                  string     `json:"gateway"`
	LandingSite              string     `json:"landing_site"`
	LandingSiteRef           any        `json:"landing_site_ref"`
	LocationID               any        `json:"location_id"`
	Name                     string     `json:"name"`
	Note                     any        `json:"note"`
	NoteAttributes           []any      `json:"note_attributes"`
	Number                   int64      `json:"number"`
	OrderNumber              int64      `json:"order_number"`
	OrderStatusURL           string     `json:"order_status_url"`
	OriginalTotalDutiesSet   any        `json:"original_total_duties_set"`
	PaymentGatewayNames      []string   `json:"payment_gateway_names"`
	Phone                    *string    `json:"phone"`
	PresentmentCurrency      string     `json:"presentment_currency"`
	ProcessedAt              *time.Time `json:"processed_at"`
	ProcessingMethod         string     `json:"processing_method"`
	Reference                any        `json:"reference"`
	ReferringSite            string     `json:"referring_site"`
	SourceIdentifier         any        `json:"source_identifier"`
	SourceName               string     `json:"source_name"`
	SourceURL                any        `json:"source_url"`
	SubtotalPrice            string     `json:"subtotal_price"`
	SubtotalPriceSet         any        `json:"subtotal_price_set"`
	Tags                     string     `json:"tags"`
	TaxLines                 []any      `json:"tax_lines"`
	TaxesIncluded            bool       `json:"taxes_included"`
	Test                     bool       `json:"test"`
	Token                    string     `json:"token"`
	TotalDiscounts           string     `json:"total_discounts"`
	TotalDiscountsSet        any        `json:"total_discounts_set"`
	TotalLineItemsPrice      string     `json:"total_line_items_price"`
	TotalLineItemsPriceSet   any        `json:"total_line_items_price_set"`
	TotalOutstanding         string     `json:"total_outstanding"`
	TotalPrice               string     `json:"total_price"`
	TotalPriceSet            any        `json:"total_price_set"`
	TotalPriceUsd            string     `json:"total_price_usd"`
	TotalShippingPriceSet    any        `json:"total_shipping_price_set"`
	TotalTax                 string     `json:"total_tax"`
	TotalTaxSet              any        `json:"total_tax_set"`
	TotalTipReceived         string     `json:"total_tip_received"`
	TotalWeight              int        `json:"total_weight"`
	UpdatedAt                *time.Time `json:"updated_at"`
	UserID                   *int64     `json:"user_id"`
	BillingAddress           any        `json:"billing_address"`
	Customer                 struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	} `json:"customer"`
	DiscountApplications []any `json:"discount_applications"`
	Fulfillments         []any `json:"fulfillments"`
	LineItems            []any `json:"line_items"`
	PaymentTerms         any   `json:"payment_terms"`
	Refunds              []any `json:"refunds"`
	ShippingLines        []any `json:"shipping_lines"`
}

type GetOrdersResponse struct {
	Orders   []Order `json:"orders"`
	PageSize int     `json:"page_size"`
}
