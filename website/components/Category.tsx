export type Category =
    "cloud-infrastructure" |
    "cloud-finops" |
    "crm" |
    "databases" |
    "data-warehouses-lakes" |
    "engineering-analytics" |
    "hr-software" |
    "marketing-analytics" |
    "product-analytics" |
    "project-management" |
    "fleet-management" |
    "shipment-tracking" |
    "other";

export function CategoryTitle(cat: Category) {
    switch (cat) {
        case "crm":
            return "CRM Software";
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
        case "fleet-management":
            return "Fleet Management";
        case "hr-software":
            return "HR Software";
        case "marketing-analytics":
            return "Marketing Analytics";
        case "product-analytics":
            return "Product Analytics";
        case "shipment-tracking":
            return "Shipment Tracking";
        case "other":
            return "Other";
        default:
            return cat;
    }
}
