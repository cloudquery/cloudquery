export type Category =
    "cloud-infrastructure" |
    "cloud-finops" |
    "databases" |
    "data-warehouses-lakes" |
    "engineering-analytics" |
    "marketing-analytics" |
    "product-analytics" |
    "project-management" |
    "fleet-management" |
    "security" |
    "shipment-tracking" |
    "other";

export const CATEGORIES: Category[] = [
    "cloud-infrastructure",
    "cloud-finops",
    "databases",
    "data-warehouses-lakes",
    "engineering-analytics",
    "marketing-analytics",
    "product-analytics",
    "project-management",
    "fleet-management",
    "security",
    "shipment-tracking",
    "other"
];

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
        case "fleet-management":
            return "Fleet Management";
        case "marketing-analytics":
            return "Marketing Analytics";
        case "product-analytics":
            return "Product Analytics";
        case "project-management":
            return "Project Management";
        case "security":
            return "Security";
        case "shipment-tracking":
            return "Shipment Tracking";
        case "other":
            return "Other";
        default:
            return cat;
    }
}
