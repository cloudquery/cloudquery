import React from "react";

export interface PricingTier {
    name: string;
    attributes: PricingTierAttributes;
}

export interface PricingTierAttributes {
    Price: string;
    PriceMonthly: boolean;
    PriceDetails?: string;
    Communication: string;
    SLA: string;
    TAM: string;
    Features: string;
    Contract: string;
}

export default function PricingCard({tier} : {tier: PricingTier}) {
    return <div
        className="flex flex-col p-6 mx-auto max-w-lg text-center text-gray-900 bg-white rounded-lg border border-gray-100 shadow dark:border-gray-600 xl:p-8 dark:bg-gray-900 dark:text-white">
        <h3 className="mb-4 text-2xl font-semibold">{tier.name}</h3>
        <div className="flex justify-center items-baseline my-8">
            <span className="mr-2 text-4xl font-extrabold">{tier.attributes.Price}</span>
            {tier.attributes.PriceMonthly ? <span className="text-gray-500 dark:text-gray-400">/month</span> : null}
        </div>

        <ul role="list" className="mb-8 space-y-4 text-left">
            <li className="flex items-center space-x-3">
                <span><strong>Communication</strong>: {tier.attributes.Communication}</span>
            </li>
            <li className="flex items-center space-x-3">
                <span><strong>SLA</strong>: {tier.attributes.SLA}</span>
            </li>
            <li className="flex items-center space-x-3">
                <span><strong>Technical Account Manager</strong>: {tier.attributes.TAM}</span>
            </li>
            <li className="flex items-center space-x-3">
                <span><strong>Priority feature development</strong>: {tier.attributes.Features}</span>
            </li>
            <li className="flex items-center space-x-3">
                <span><strong>Contract</strong>: {tier.attributes.Contract}</span>
            </li>
        </ul>
    </div>
}