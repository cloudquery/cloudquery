export type Category =
    "cloud-infrastructure" |
    "cloud-finops" |
    "databases" |
    "data-warehouses-lakes" |
    "engineering-analytics" |
    "marketing-analytics" |
    "product-analytics" |
    "other";

export function CategoryTitle(cat: Category) {
    switch (cat) {
        case "cloud-infrastructure":
            return "Cloud Infrastructure";
        case "databases":
            return "Databases";
        case "data-warehouses-lakes":
            return "Data Warehouses & Lakes";
        case "engineering-analytics":
            return "Engineering Analytics";
        case "cloud-finops":
            return "Cloud FinOps";
        case "marketing-analytics":
            return "Marketing Analytics";
        case "product-analytics":
            return "Product Analytics";
        case "other":
            return "Other";
    }
}
