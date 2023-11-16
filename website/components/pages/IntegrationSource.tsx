import {ALL_DESTINATION_PLUGINS, Plugin} from "../pluginData";
import {LogoContainer} from "../LogoContainer";
import React from "react";
import Head from "next/head";
import PluginPricingCard from "../PluginPricingCard";

const licenses = {
    "unpublished": [
        {
            id: "standard",
            name: "Standard License",
            description: "For internal use in one organization",
            attributes: {
                Price: "$250",
                PriceDetails: "USD. Perpetual fallback license.",
                Limits: "No usage-based billing or row restrictions.",
                Includes: "6 months of support and updates included.",
                NormalPrice: "$500 (50% pre-order discount)",
            }
        },
        {
            id: "extended",
            name: "Extended License",
            description: "For use in customer-facing products",
            attributes: {
                Price: "$1250",
                PriceDetails: "USD. Perpetual fallback license.",
                Limits: "No usage-based billing or row restrictions.",
                Includes: "6 months of support and updates included.",
                NormalPrice: "$2500 (50% pre-order discount)",
            }
        }
    ],
    "premium": [
        {
            id: "standard",
            name: "Standard License",
            description: "For internal use in one organization",
            attributes: {
                Price: "$500",
                PriceDetails: "USD. Perpetual fallback license.",
                Limits: "No usage-based billing or row restrictions.",
                Includes: "6 months of support and updates included.",
            }
        },
        {
            id: "extended",
            name: "Extended License",
            description: "For use in customer-facing products",
            attributes: {
                Price: "$2500",
                PriceDetails: "USD. Perpetual fallback license",
                Limits: "No usage-based billing or row restrictions.",
                Includes: "6 months of support and updates included.",
            }
        }
    ]
};

export default function Integration({
    source
                                    }: {
    source: Plugin;
}) {
    return <>
        <Head>
            <meta property="og:description" content={"Sync from " + source.name + " to any destination: PostgreSQL, MySQL, BigQuery, S3, DuckDB and many more"} />
        </Head>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto lg:max-w-7xl">
                <div>
                    <h1 className="nx-text-6xl font-extrabold tracking-tighter leading-[1.1] text-7xl sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
                        Export from&nbsp;
                        <span className="hidden lg:block"></span>
                        <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
                        {source.name}
                        </span>
                        &nbsp;to any destination
                    </h1>
                    <p className="mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        CloudQuery is an open-source, self-hosted ELT platform that allows you to extract data from any source into any destination. It requires no account{source.availability === 'free' ? ', imposes no limits on rows and takes only minutes to get started.' : ' and imposes no limits on rows.'}
                    </p>
                    <p className="mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        {
                            {
                                'unpublished': <>The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination, including <a href="https://hub.cloudquery.io/plugins/destination"  className="text-blue-500 hover:text-blue-600">PostgreSQL, MySQL, S3, Snowflake, BigQuery, and many more</a>. It is currently available for <strong>pre-order</strong>.</>,
                                'free': <>The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination. Select the destination you would like to sync {source.name} data from the list of supported destinations below:</>,
                                'premium': <>The CloudQuery {source.name} plugin is a premium plugin that allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination, including <a href="https://hub.cloudquery.io/plugins/destination"  className="text-blue-500 hover:text-blue-600">PostgreSQL, MySQL, S3, Snowflake, BigQuery, and many more</a>. Purchase it using the links below.</>,
                                'partner': <>The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination. Select the destination you would like to sync {source.name} data from the list of supported destinations below:</>,
                                'community': <>The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination. Select the destination you would like to sync {source.name} data from the list of supported destinations below:</>,
                            }[source.availability]
                        }
                    </p>
                </div>
            </div>

            {
                (source.availability === 'unpublished' || source.availability === 'premium') ?
                    <>
                        <div className="space-y-8 lg:grid lg:grid-cols-2 sm:gap-6 xl:gap-10 lg:space-y-0 max-w-3xl mx-auto">
                            {licenses[source.availability].map((license) => (
                                <PluginPricingCard
                                    key={license.name}
                                    license={license}
                                    preOrder={source.availability === "unpublished"}
                                    buyLink={"/buy/" + source.id + "-" + license.id}
                                />
                            ))}
                        </div>
                        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
                            <h2 className="mx-auto mt-24 nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">FAQ</h2>
                            <h3 className="font-bold text-lg text-gray-700 lg:max-w-3xl lg:nx-text-xl mt-4 dark:text-white">
                                What is the difference between the standard and extended license?
                            </h3>
                            <p className="font-light text-gray-700 lg:nx-text-xl dark:text-white">
                                The standard license is for use in internal products, such as dashboards, internal cloud inventory tools, or other products that are not customer-facing. The extended license is for use of the data in customer-facing products, such as SaaS applications or security products.
                            </p>
                            <h3 className="font-bold text-lg text-gray-700 lg:nx-text-xl mt-4 dark:text-white">
                                What is a perpetual fallback license?
                            </h3>
                            <p className="font-light text-gray-700 lg:nx-text-xl dark:text-white">
                                A perpetual fallback license is a license that allows you to use the plugin in perpetuity, even if you stop paying for updates. The license will never expire, but you will not receive updates after your license expires. You will have access to all new versions of the plugin published up to the date your license expires.
                            </p>
                            <h3 className="font-bold text-lg text-gray-700 lg:nx-text-xl mt-4 dark:text-white">
                                What tables and columns will be included?
                            </h3>
                            <p className="font-light text-gray-700 lg:nx-text-xl dark:text-white">
                                {source.availability === "premium" ?
                                    <>You can find a list of the tables and their columns in the <a className="text-blue-500 hover:text-blue-600" href={"/docs/plugins/sources/" + source.id + "/tables"}>{source.name} plugin documentation</a>.</>
                                    :
                                    <>The tables in the first version will be closely modeled around the endpoints available in the official {source.name} API. After initial release you will also be able to request more tables and columns as part of the included support.</>
                                }
                            </p>
                        </div>
                    </>
                    :
                    <>
                        <div className="mx-4 px-4 mx-auto w-auto">
                            <div className="flex justify-left items-left flex-wrap gap-9 pt-8 sm:mt-4">
                                {ALL_DESTINATION_PLUGINS.map((plugin) => (
                                    <LogoContainer
                                        title={plugin.name}
                                        href={plugin.href || `/integrations/${source.id}/${plugin.id}`}
                                        key={plugin.id}
                                        logo={plugin.logo}
                                        logoDark={plugin.logoDark}
                                        name={plugin.name}
                                        external={false}
                                        category={plugin.category}
                                        availability={plugin.availability}
                                    >
                                    </LogoContainer>
                                ))}
                            </div>
                        </div>
                        <p className="mx-auto mt-24 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                            Couldn't find the destination you were looking for? Please <a href="https://github.com/cloudquery/cloudquery/issues/new/choose" target="_blank" className="text-blue-500 hover:text-blue-600">request it on GitHub ↗</a> and we'll add it to our roadmap. New plugins are created all the time, often within days of a request. Showing your interest helps us prioritize.
                        </p>
                    </>
            }
        </div>
    </>;
}
