import {Plugins} from "../Plugins";
import React from "react";

export default function Integrations() {
    return <>
        <section>
            <div className="py-8 max-w-8xl px-4 mx-auto lg:py-14 lg:px-8">
                <div className="mx-auto max-w-screen-md text-center ">
                    <h2 className="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">Integrations</h2>
                    <p className="mb-5 font-light text-gray-500 sm:text-xl dark:text-gray-400">
                        Export data from an ever-growing list of <a className="dark:text-white" href="https://hub.cloudquery.io/plugins/source">cloud providers, databases and SaaS apps</a> with more than 1,000 unique tables. Sync to your <a href="https://hub.cloudquery.io/plugins/destination" className="dark:text-white">favorite database, data warehouse or data lake</a>.
                    </p>
                </div>
                <div>
                    <Plugins />
                </div>
            </div>
        </section>
    </>;
}
