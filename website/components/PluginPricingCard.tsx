import React from "react";

export interface PluginLicense {
    name: string;
    description?: string;
    attributes: PluginLicenseAttributes;
}

export interface PluginLicenseAttributes {
    Price: string;
    PriceMonthly: boolean;
    PriceDetails?: string;
    Limits?: string;
    Includes?: string;
    NormalPrice?: string;
}

export default function PluginPricingCard({license, preOrder, buyLink} : {license: PluginLicense, preOrder: boolean, buyLink: string}) {
    const check = <svg className="flex-shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                       viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
        <path fillRule="evenodd"
              d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
              clipRule="evenodd"></path>
    </svg>;

    return <div
        className="flex flex-col px-6 mx-auto max-w-lg text-center text-gray-900 bg-white rounded-lg border border-gray-100 shadow dark:border-gray-600 p-8 dark:bg-gray-900 dark:text-white">
        <h3 className="mb-4 text-2xl font-semibold">{license.name}</h3>
        <p className="font-light text-gray-500 sm:text-lg dark:text-gray-400">{license.description}</p>
        <div className={"flex justify-center items-baseline mt-8" + (license.attributes.PriceDetails ? " mb-2" : " sm:mb-8 lg:mb-16") }>
            <span className="mr-2 text-4xl font-extrabold">{license.attributes.Price}</span>
            {license.attributes.PriceMonthly ? <span className="text-gray-500 dark:text-gray-400">/month</span> : null}
        </div>
        {license.attributes.PriceDetails ?
            <div className="mb-4">
                <p className="font-light text-gray-500 text-sm dark:text-gray-400">{license.attributes.PriceDetails}</p>
            </div>
            : null}

        <ul role="list" className="space-y-4 text-left mb-20">
            {license.attributes.Limits ?
                <li className="flex items-center space-x-3">
                    {check}
                    <span><strong>Unlimited</strong>: {license.attributes.Limits}</span>
                </li> : null }
            {license.attributes.Includes ?
                <li className="flex items-center space-x-3">
                    {check}
                    <span><strong>Updates</strong>: {license.attributes.Includes}</span>
                </li> : null }
            {license.attributes.NormalPrice ?
                <li className="flex items-center space-x-3">
                    {check}
                    <span><strong>Pre-order Discount</strong>: Normal price {license.attributes.NormalPrice}</span>
                </li> : null }
            <li className="flex items-center space-x-3">
                {check}
                <span><strong>Peace of Mind Guarantee</strong>: 14-day money-back guarantee</span>
            </li>
        </ul>
        <div className="relative h-full">
            <div className="absolute inset-x-0 bottom-0">
                <a href={buyLink}
                   className="btn btn-green w-full block"
                >
                    {preOrder ? "Pre-order Now" : "Buy Now"}
                </a>
            </div>
        </div>

    </div>
}