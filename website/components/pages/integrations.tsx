import {Sources} from "../Sources";
import {Destinations} from "../Destinations";
import React from "react";
import PricingCard from "../PricingCard";

export default function Integrations() {
    return <>
        <section>
            <div className="py-8 max-w-8xl px-4 mx-auto lg:py-14 lg:px-8">
                <div className="mx-auto max-w-screen-md text-center ">
                    <h2 className="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">Integrations</h2>
                    <p className="mb-5 font-light text-gray-500 sm:text-xl dark:text-gray-400">
                        Export data from an ever-growing list of <a className="dark:text-white" href="/docs/plugins/sources/overview">cloud providers, databases and SaaS apps</a> with more than 1,000 unique tables. Sync to your <a href="/docs/plugins/destinations/overview" className="dark:text-white">favorite database, data warehouse or data lake</a>.
                    </p>
                </div>
                <div>
                    <Sources />
                </div>
            </div>
        </section>
        {/*<div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8">*/}
        {/*    <div className="sm:py-20 lg:py-24">*/}
        {/*        <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">*/}
        {/*            <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">*/}
        {/*                */}
        {/*            </h2>*/}
        {/*            <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">*/}
        {/*                */}
        {/*            </p>*/}
        {/*            */}
        {/*        </div>*/}
        {/*    </div>*/}
        {/*</div>*/}
    </>;
}
