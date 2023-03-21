import React from "react";

export interface PricingTier {
    name: string;
    description?: string;
    attributes: PricingTierAttributes;
}

export interface PricingTierAttributes {
    Price: string;
    PriceMonthly: boolean;
    PriceDetails?: string;
    Communication?: string;
    SLA?: string;
    TAM?: string;
    Features?: string;
    Contract?: string;
}

export default function PricingCard({tier} : {tier: PricingTier}) {
    const check = <svg className="flex-shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                       viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
        <path fill-rule="evenodd"
              d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
              clip-rule="evenodd"></path>
    </svg>;

    return <div
        className="flex flex-col px-6 mx-auto max-w-lg text-center text-gray-900 bg-white rounded-lg border border-gray-100 shadow dark:border-gray-600 p-8 dark:bg-gray-900 dark:text-white">
        <h3 className="mb-4 text-2xl font-semibold">{tier.name}</h3>
        <p className="font-light text-gray-500 sm:text-lg dark:text-gray-400">{tier.description}</p>
        <div className={"flex justify-center items-baseline mt-8" + (tier.attributes.PriceDetails ? " mb-2" : " sm:mb-8 lg:mb-16") }>
            <span className="mr-2 text-4xl font-extrabold">{tier.attributes.Price}</span>
            {tier.attributes.PriceMonthly ? <span className="text-gray-500 dark:text-gray-400">/month</span> : null}
        </div>
        {tier.attributes.PriceDetails ?
        <div className="mb-4">
                <p className="font-light text-gray-500 text-sm dark:text-gray-400">{tier.attributes.PriceDetails}</p>
        </div>
        : null}

        <ul role="list" className="space-y-4 text-left mb-20">
            {tier.attributes.Communication ?
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>Communication</strong>: {tier.attributes.Communication}</span>
            </li> : null }
            {tier.attributes.SLA ?
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>SLA</strong>: {tier.attributes.SLA}</span>
            </li> : null }
            {tier.attributes.TAM ?
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>Technical Account Manager</strong>: {tier.attributes.TAM}</span>
            </li> : null }
            {tier.attributes.Features ?
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>Priority feature development</strong>: {tier.attributes.Features}</span>
            </li> : null }
            {tier.attributes.Contract ?
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>Contract</strong>: {tier.attributes.Contract}</span>
            </li> : null }
        </ul>
        <div className="relative h-full">
            <div className="absolute inset-x-0 bottom-0">
                {tier.name == 'Community' ?
                    <a href="/docs/quickstart"
                       className="btn btn-purple w-full block"
                    >
                        Get Started
                    </a>
                    :
                    <a href="/contact"
                        className="btn btn-purple w-full block"
                    >
                        Contact Us
                    </a>
                }
            </div>
        </div>

    </div>
}