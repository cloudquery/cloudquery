package shopify

import "time"

type Checkout struct {
	ID                       int64      `json:"id"`
	Token                    string     `json:"token"`
	CartToken                string     `json:"cart_token"`
	Email                    string     `json:"email"`
	Gateway                  string     `json:"gateway"`
	BuyerAcceptsMarketing    bool       `json:"buyer_accepts_marketing"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
	LandingSite              string     `json:"landing_site"`
	Note                     any        `json:"note"`
	NoteAttributes           []any      `json:"note_attributes"`
	ReferringSite            string     `json:"referring_site"`
	ShippingLines            []any      `json:"shipping_lines"`
	TaxesIncluded            bool       `json:"taxes_included"`
	TotalWeight              int        `json:"total_weight"`
	Currency                 string     `json:"currency"`
	CompletedAt              time.Time  `json:"completed_at"`
	ClosedAt                 *time.Time `json:"closed_at"`
	UserID                   *int64     `json:"user_id"`
	LocationID               any        `json:"location_id"`
	SourceIdentifier         any        `json:"source_identifier"`
	SourceURL                any        `json:"source_url"`
	DeviceID                 any        `json:"device_id"`
	Phone                    any        `json:"phone"`
	CustomerLocale           string     `json:"customer_locale"`
	LineItems                []any      `json:"line_items"`
	Name                     string     `json:"name"`
	Source                   any        `json:"source"`
	AbandonedCheckoutURL     string     `json:"abandoned_checkout_url"`
	DiscountCodes            []any      `json:"discount_codes"`
	TaxLines                 []any      `json:"tax_lines"`
	SourceName               string     `json:"source_name"`
	PresentmentCurrency      string     `json:"presentment_currency"`
	BuyerAcceptsSmsMarketing bool       `json:"buyer_accepts_sms_marketing"`
	SmsMarketingPhone        any        `json:"sms_marketing_phone"`
	TotalDiscounts           string     `json:"total_discounts"`
	TotalLineItemsPrice      string     `json:"total_line_items_price"`
	TotalPrice               string     `json:"total_price"`
	TotalTax                 string     `json:"total_tax"`
	SubtotalPrice            string     `json:"subtotal_price"`
	TotalDuties              any        `json:"total_duties"`
	BillingAddress           any        `json:"billing_address"`
	Customer                 struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	} `json:"customer"`
}

type GetCheckoutsResponse struct {
	Checkouts []Checkout `json:"checkouts"`
	PageSize  int        `json:"page_size"`
}
