// Code generated by codegen; DO NOT EDIT.

package order

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Orders() *schema.Table {
	return &schema.Table{
		Name:     "shopify_orders",
		Resolver: fetchOrders,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "admin_graphql_api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdminGraphqlAPIID"),
			},
			{
				Name:     "app_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AppID"),
			},
			{
				Name:     "browser_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrowserIP"),
			},
			{
				Name:     "buyer_accepts_marketing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BuyerAcceptsMarketing"),
			},
			{
				Name:     "cancelled_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CancelledAt"),
			},
			{
				Name:     "cart_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CartToken"),
			},
			{
				Name:     "checkout_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CheckoutID"),
			},
			{
				Name:     "checkout_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CheckoutToken"),
			},
			{
				Name:     "closed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClosedAt"),
			},
			{
				Name:     "confirmed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Confirmed"),
			},
			{
				Name:     "contact_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactEmail"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "currency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Currency"),
			},
			{
				Name:     "current_subtotal_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentSubtotalPrice"),
			},
			{
				Name:     "current_total_discounts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentTotalDiscounts"),
			},
			{
				Name:     "current_total_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentTotalPrice"),
			},
			{
				Name:     "current_total_tax",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentTotalTax"),
			},
			{
				Name:     "customer_locale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerLocale"),
			},
			{
				Name:     "discount_codes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiscountCodes"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "estimated_taxes",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EstimatedTaxes"),
			},
			{
				Name:     "financial_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FinancialStatus"),
			},
			{
				Name:     "gateway",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Gateway"),
			},
			{
				Name:     "landing_site",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LandingSite"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "note_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NoteAttributes"),
			},
			{
				Name:     "number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Number"),
			},
			{
				Name:     "order_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OrderNumber"),
			},
			{
				Name:     "order_status_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OrderStatusURL"),
			},
			{
				Name:     "payment_gateway_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PaymentGatewayNames"),
			},
			{
				Name:     "phone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Phone"),
			},
			{
				Name:     "presentment_currency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PresentmentCurrency"),
			},
			{
				Name:     "processed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ProcessedAt"),
			},
			{
				Name:     "processing_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProcessingMethod"),
			},
			{
				Name:     "referring_site",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReferringSite"),
			},
			{
				Name:     "source_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceName"),
			},
			{
				Name:     "subtotal_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubtotalPrice"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "tax_lines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TaxLines"),
			},
			{
				Name:     "taxes_included",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TaxesIncluded"),
			},
			{
				Name:     "test",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Test"),
			},
			{
				Name:     "token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Token"),
			},
			{
				Name:     "total_discounts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalDiscounts"),
			},
			{
				Name:     "total_line_items_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalLineItemsPrice"),
			},
			{
				Name:     "total_outstanding",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalOutstanding"),
			},
			{
				Name:     "total_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalPrice"),
			},
			{
				Name:     "total_price_usd",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalPriceUsd"),
			},
			{
				Name:     "total_tax",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalTax"),
			},
			{
				Name:     "total_tip_received",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalTipReceived"),
			},
			{
				Name:     "total_weight",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalWeight"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UserID"),
			},
			{
				Name:     "customer",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Customer"),
			},
			{
				Name:     "discount_applications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiscountApplications"),
			},
			{
				Name:     "fulfillments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Fulfillments"),
			},
			{
				Name:     "line_items",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LineItems"),
			},
			{
				Name:     "refunds",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Refunds"),
			},
			{
				Name:     "shipping_lines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShippingLines"),
			},
		},
	}
}
