import {Sources} from "../Sources";
import {Destinations} from "../Destinations";
import React from "react";

export default function Integrations() {
    return <>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="sm:py-20 lg:py-24">
                <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">
                    <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
                        Source Integrations
                    </h2>
                    <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Export data from a growing list of <a className="dark:text-white" href="/docs/plugins/sources/overview">30+ cloud providers, databases and SaaS apps</a> with more than 1,000 unique tables. Sync to your <a href="/docs/plugins/destinations/overview" className="dark:text-white">favorite database, data warehouse or data lake</a>.
                    </p>
                    <Sources />
                    <p className="mx-auto mt-24 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Couldn't find the source you were looking for? <a href="/new-source-plugin" className="text-blue-500 hover:text-blue-600">Request a New Source Plugin</a>
                    </p>
                </div>
                <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">
                    <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
                        Destinations
                    </h2>
                    <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        CloudQuery currently supports the following destination databases, data warehouses and data lakes:
                    </p>
                    <Destinations />
                    <p className="mx-auto mt-24 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Couldn't find the destination you were looking for? Please <a href="https://github.com/cloudquery/cloudquery/issues/new/choose" target="blank" className="text-blue-500 hover:text-blue-600">request it on GitHub ↗</a> and we'll add it to our roadmap. New destinations are created all the time, often within days of a request. Showing your interest helps us prioritize.
                    </p>
                </div>
            </div>
        </div>
    </>;
}
