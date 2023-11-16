export type Policy = {
    id: string;
    name: string;
    buyLinks?: any;
    availableForPurchase?: boolean;
};

export const ALL_PREMIUM_POLICIES: Policy[] = [
    {
        id: "aws-cost-pack",
        name: "AWS Cost Optimization Pack",
        buyLinks: {
            standard: "https://buy.stripe.com/14kbIMdapgW35ig4kj",
            extended: "https://buy.stripe.com/fZe9AE0nD7lt1207wu",
        },
        availableForPurchase: true
    },

    {
        id: "aws_cis_v1_5_0_bigquery",
        name: "AWS CIS V1.5.0 for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/4gw8wA2vLbBJ7qoeYK",
            extended: "https://buy.stripe.com/fZe7sw7Q5axFeSQ17V",
        },
    },

    {
        id: "aws_cis_v1_5_0_snowflake",
        name: "AWS CIS V1.5.0 for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/fZe8wAc6laxF5ig6sg",
            extended: "https://buy.stripe.com/14k9AE3zPfRZ8usaIx",
        },
    },

    {
        id: "aws_foundational_security_bigquery",
        name: "AWS Foundational Security for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/cN2bIM0nD7lt2647we",
            extended: "https://buy.stripe.com/cN29AEfix49hfWU2bV",
        },
    },

    {
        id: "aws_foundational_security_snowflake",
        name: "AWS Foundational Security for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/fZe6os8U9dJR8us9Eo",
            extended: "https://buy.stripe.com/6oEdQUdap8pxaCAeYJ",
        },
        availableForPurchase: true,
    },

    {
        id: "azure_cis_v1_3_0_bigquery",
        name: "Azure CIS V1.3.0 for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/14kbIM6M10X57qog2W",
            extended: "https://buy.stripe.com/aEUaEIgmB5dlh0Y03Z",
        },
    },

    {
        id: "azure_cis_v1_3_0_snowflake",
        name: "Azure CIS V1.3.0 for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/3cseUY0nDeNVh0Y7ws",
            extended: "https://buy.stripe.com/eVa28c3zP35ddOMbMJ",
        },
    },

    {
        id: "azure-cost-pack",
        name: "Azure Cost Optimization Pack",
        buyLinks: {
            standard: "https://buy.stripe.com/14k8wA8U97lt264eYZ",
            extended: "https://buy.stripe.com/7sI8wAfix5dldOMaII",
        },
    },

    {
        id: "azure_hippa_hitrust_v9_2_bigquery",
        name: "Azure HIPPA HITRUST v9.2 for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/dR69AE8U99tB4ec8Aq",
            extended: "https://buy.stripe.com/bIYbIM0nD5dl3a83g7",
        },
    },

    {
        id: "azure_hippa_hitrust_v9_2_snowflake",
        name: "Azure HIPPA HITRUST v9.2 for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/9AQ28cb2hgW34eccQI",
            extended: "https://buy.stripe.com/dR6eUYeetgW3bGEbMF",
        },
    },

    {
        id: "gcp_cis_v1_2_0_bigquery",
        name: "GCP CIS V1.2.0 for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/7sI3cg5HX8px120dUy",
            extended: "https://buy.stripe.com/9AQ0045HX7ltcKI3fV",
        },
    },

    {
        id: "gcp_cis_v1_2_0_snowflake",
        name: "GCP CIS V1.2.0 for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/fZeeUY2vL8pxeSQ5o4",
            extended: "https://buy.stripe.com/8wM8wAdap35d8us5o5",
        },
    },

    {
        id: "gcp-cost-pack",
        name: "GCP Cost Optimization Pack",
        buyLinks: {
            standard: "https://buy.stripe.com/28ocMQ2vL49hh0Y2ce",
            extended: "https://buy.stripe.com/3cs3cgfixbBJ4ecaIL",
        },
    },

    {
        id: "k8s_nsa_cisa_v1_bigquery",
        name: "K8S NSA CISA V1 for BigQuery",
        buyLinks: {
            standard: "https://buy.stripe.com/9AQ4gk4DT35deSQ5nY",
            extended: "https://buy.stripe.com/8wM148fix2196mk03F",
        },
    },

    {
        id: "k8s_nsa_cisa_v1_snowflake",
        name: "K8S NSA CISA V1 for Snowflake",
        buyLinks: {
            standard: "https://buy.stripe.com/00g9AE5HXdJRbGEeYA",
            extended: "https://buy.stripe.com/cN27sw1rH9tBfWU7w9",
        },
    }
];