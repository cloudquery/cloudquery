package shopify

import "time"

type Product struct {
	ID             int64            `json:"id"`
	Title          string           `json:"title"`
	BodyHTML       string           `json:"body_html"`
	Vendor         string           `json:"vendor"`
	ProductType    string           `json:"product_type"`
	CreatedAt      time.Time        `json:"created_at"`
	Handle         string           `json:"handle"`
	UpdatedAt      time.Time        `json:"updated_at"`
	PublishedAt    *time.Time       `json:"published_at"`
	TemplateSuffix string           `json:"template_suffix"`
	Status         string           `json:"status"`
	PublishedScope string           `json:"published_scope"`
	Tags           Tags             `json:"tags"`
	Variants       []ProductVariant `json:"variants"`

	Options           any            `json:"options"`
	Images            []ProductImage `json:"images"`
	Image             ProductImage   `json:"image"`
	AdminGraphqlAPIID string         `json:"admin_graphql_api_id"`
}

type ProductVariant struct {
	ProductID            int64     `json:"product_id"`
	ID                   int64     `json:"id"`
	Title                string    `json:"title"`
	Price                string    `json:"price"`
	Sku                  string    `json:"sku"`
	Position             int64     `json:"position"`
	InventoryPolicy      string    `json:"inventory_policy"`
	CompareAtPrice       string    `json:"compare_at_price"`
	FulfillmentService   string    `json:"fulfillment_service"`
	InventoryManagement  string    `json:"inventory_management"`
	Option1              any       `json:"option1"`
	Option2              any       `json:"option2"`
	Option3              any       `json:"option3"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Taxable              bool      `json:"taxable"`
	Barcode              string    `json:"barcode"`
	Grams                int64     `json:"grams"`
	ImageID              any       `json:"image_id"`
	Weight               float64   `json:"weight"`
	WeightUnit           string    `json:"weight_unit"`
	InventoryItemID      int64     `json:"inventory_item_id"`
	InventoryQuantity    int64     `json:"inventory_quantity"`
	OldInventoryQuantity int64     `json:"old_inventory_quantity"`
	RequiresShipping     bool      `json:"requires_shipping"`
	AdminGraphqlAPIID    string    `json:"admin_graphql_api_id"`
}

type ProductImage struct {
	ProductID         int64     `json:"product_id"`
	ID                int64     `json:"id"`
	Position          int64     `json:"position"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Alt               string    `json:"alt"`
	Width             int64     `json:"width"`
	Height            int64     `json:"height"`
	Src               string    `json:"src"`
	VariantIDs        []any     `json:"variant_ids"`
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
	PageSize int       `json:"page_size"`
}
