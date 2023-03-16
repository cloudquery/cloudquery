import {Plugin} from "../pluginData";
import DownloadSection from '../mdx/_download-all.mdx'
import React, {ReactNode} from "react";

function unpublishedPlugin({
                               source,
                               destination,
                           }: {
    source: Plugin;
    destination: Plugin;
}) {
    return <>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto">
                <div className="flex flex-col justify-between md:mr-4">
                    <div>
                        <h1 className="max-w-5xl mx-auto nx-text-6xl font-extrabold tracking-tighter leading-[1.1] text-7xl sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
                            Export from&nbsp;
                            <span className="hidden lg:block"></span>
                            <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
                             {source.name}
                            </span>
                            &nbsp;to&nbsp;
                            <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-green-500 ">
                            {destination.name}
                            </span>
                        </h1>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            CloudQuery is an open-source data integration platform that allows you to export data from any source to any destination. It's free, open source, requires no account, and takes only minutes to get started.
                        </p>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            A CloudQuery {source.name} plugin will allow you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any destination, including {destination.name} and <a href={"/integrations/" + source.id} className="text-blue-500 hover:text-blue-600">many others</a>.
                        </p>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            Upvote and subscribe to <a href={source.href} target="_blank" className="text-blue-500 hover:text-blue-600">the issue on GitHub ↗</a> to show your interest and get notified when the {source.name} plugin is released. New plugins are created all the time, often within days of a request. Showing your interest in a plugin helps us prioritize it.
                        </p>
                        <hr className="h-px my-8 bg-gray-200 border-0 dark:bg-gray-700" />
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            Eager to get started immediately? CloudQuery is open source and pluggable, so you can build your own {source.name} plugin and start loading data into any supported destination today. Check out our <a href="/docs/developers/creating-new-plugin" className="text-blue-500 hover:text-blue-600">plugin development guide</a> to get started.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </>;
}

export default function Integration({
    source,
    sourceConfiguration,
    sourceAuthentication,
    destination,
    destinationConfiguration,
    destinationAuthentication,
    syncCommand,
                                    }: {
    source: Plugin;
    sourceConfiguration?: ReactNode;
    sourceAuthentication?: ReactNode;
    destination: Plugin;
    destinationConfiguration?: ReactNode;
    destinationAuthentication?: ReactNode;
    syncCommand: ReactNode;
}) {
    const sourceFilename = source.id === destination.id ? `source-${source.id}.yaml` : `${source.id}.yaml`;
    const destinationFilename = source.id === destination.id ? `destination-${destination.id}.yaml` : `${destination.id}.yaml`;
    if (source.kind === "unpublished" || destination.kind === "unpublished") {
        return unpublishedPlugin({source, destination});
    }
    return <>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto">
                <div className="flex flex-col justify-between md:mr-4">
                    <div>
                        <h1 className="max-w-5xl mx-auto nx-text-6xl font-extrabold tracking-tighter leading-[1.1] text-7xl sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
                            Export from&nbsp;
                            <span className="hidden lg:block"></span>
                            <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
                             {source.name}
                            </span>
                            &nbsp;to&nbsp;
                            <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-green-500 ">
                            {destination.name}
                            </span>
                        </h1>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            CloudQuery is an open-source data integration platform that allows you to export data from any source to any destination.
                        </p>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                            The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any destination, including {destination.name}. It's free, open source, requires no account, and takes only minutes to get started.
                        </p>
                        <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        Ready? Let's dive right in!
                        </p>
                    </div>
                </div>
            </div>

            <div className="relative from-gray-50 to-gray-100">
                <div className="py-8">
                    <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                        Step 1. Install the CloudQuery CLI
                    </h2>
                    <p className="mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                        The CloudQuery CLI is a command-line tool that runs the sync. It supports MacOS, Linux and Windows.
                    </p>
                    <DownloadSection />
                </div>
            </div>

            <div className="relative from-gray-50 to-gray-100">
                <div className="py-8">
                    <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                        Step 2. Configure the {source.name} source plugin
                    </h2>
                    <p className="mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                        Create a configuration file for the {source.name} plugin and set up authentication.
                    </p>
                    {source.kind !== "official" ?
                        <>
                            <p className="mt-4">
                                {source.name} is a {source.kind} plugin, which means that it is maintained by the {source.kind === "community" ? "CloudQuery community" : source.name + " team"}. Create a file called <code className="text-lg nx-font-bold">{sourceFilename}</code>, then copy the example and follow the instructions in the <a target="_blank" href={source.href} className="text-blue-500 hover:text-blue-600">{source.name} Plugin Documentation ↗</a> to fit your needs.
                            </p>
                        </>
                        :
                        <>
                            <h3 className="mt-4 nx-text-2xl font-extrabold tracking-tight lg:nx-text-3xl xl:nx-text-4xl dark:text-white">Configuration</h3>
                            <p className="mt-4">
                                Create a file called <code className="text-lg nx-font-bold">{sourceFilename}</code> and add the following contents:
                            </p>
                            {sourceConfiguration}
                            <p className="mt-4">
                                Fine-tune this configuration to match your needs. For more information, see the <a target="_blank" href={"/docs/plugins/sources/" + source.id + "/overview"} className="text-blue-500 hover:text-blue-600">{source.name} Plugin ↗</a> page in the docs.
                            </p>
                            {sourceAuthentication ? <div>
                                <h3 className="mt-4 nx-text-2xl font-extrabold tracking-tight lg:nx-text-3xl xl:nx-text-4xl dark:text-white">Authentication</h3>
                                {sourceAuthentication}
                            </div> : <></>}
                        </>
                    }
                </div>
            </div>

            <div className="relative from-gray-50 to-gray-100">
                <div className="py-8">
                    <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                        Step 3. Configure the {destination.name} destination plugin
                    </h2>
                    {destination.kind !== "official" ?
                        <>
                            <p className="mt-4">
                                {destination.name} is a {destination.kind} plugin, which means that it is maintained by the {destination.kind === "community" ? "CloudQuery community" : destination.name + " team"}. Create a file called <code className="text-lg nx-font-bold">{destinationFilename}</code>, then copy the example and follow the instructions in the <a target="_blank" href={destination.href} className="text-blue-500 hover:text-blue-600">{destination.name} Plugin Documentation ↗</a> to fit your needs.
                            </p>
                        </>
                        :
                        <>
                            <p className="mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                                Create a configuration file for the {destination.name} plugin and set up authentication.
                            </p>
                            <h3 className="mt-4 nx-text-2xl font-extrabold tracking-tight lg:nx-text-3xl xl:nx-text-4xl dark:text-white">Configuration</h3>
                            <p className="mt-4">
                                Create a file called <code className="text-lg nx-font-bold">{destinationFilename}</code> and add the following contents:
                            </p>
                            {destinationConfiguration}
                            <p className="mt-4">
                                Fine-tune this configuration to match your needs. For more information, see the <a target="_blank" href={"/docs/plugins/destinations/" + destination.id + "/overview"} className="text-blue-500 hover:text-blue-600">{destination.name} Plugin ↗</a> page in the docs.
                            </p>
                            {destinationAuthentication ? <div>
                                <h3 className="mt-4 nx-text-2xl font-extrabold tracking-tight lg:nx-text-3xl xl:nx-text-4xl dark:text-white">Authentication</h3>
                                {destinationAuthentication}
                            </div> : <></>}
                        </>
                    }
                </div>
            </div>

            <div className="relative from-gray-50 to-gray-100">
                <div className="py-8">
                    <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                        Step 4. Start the Sync
                    </h2>
                    <p className="mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                        Run the following command in your terminal to start the sync
                    </p>
                    {syncCommand}
                    <p>
                        And away we go! 🚀 The sync will run until completion, fetching all selected tables from {source.name}. Any errors will be logged to a file called <code className="text-lg nx-font-bold">cloudquery.log</code>.
                    </p>
                </div>
            </div>

            <div className="relative from-gray-50 to-gray-100">
                <div className="py-8">
                    <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                        Further Reading
                    </h2>
                    <p className="mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                        Now that you've seen the basics of syncing {source.name} to {destination.name}, you should know that there's a lot more you can do. Check out the CloudQuery <a href={"/docs"} className="text-blue-500 hover:text-blue-600">Documentation</a>, <a href={"https://github.com/cloudquery/cloudquery"} className="text-blue-500 hover:text-blue-600">Source Code</a> and <a href={"/how-to-guides"} className="text-blue-500 hover:text-blue-600">How-to Guides</a> for more details.
                    </p>
                </div>
            </div>
        </div>
    </>;
}
