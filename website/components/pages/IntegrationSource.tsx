import {DESTINATION_PLUGINS, Plugin} from "../pluginData";
import {LogoContainer} from "../LogoContainer";
import React from "react";

export default function Integration({
    source
                                    }: {
    source: Plugin;
}) {
    return <>

        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto lg:max-w-7xl">
                <div>
                    <h1 className="max-w-5xl mx-auto nx-text-6xl font-extrabold tracking-tighter leading-[1.1] text-7xl sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
                        Export from&nbsp;
                        <span className="hidden lg:block"></span>
                        <span className="pr-1 pb-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
                        {source.name}
                        </span>
                        &nbsp;to any destination
                    </h1>
                    <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        CloudQuery is an open-source ELT platform that allows you to extract data from any source into any destination. It is a free and open source alternative to Fivetran or Airbyte that requires no account, and it takes only minutes to get started.
                    </p>
                    <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        {source.kind === "unpublished" ?
                            <>A CloudQuery {source.name} plugin is on our roadmap. Please upvote and subscribe to the <a href={source.href} className="text-blue-500 hover:text-blue-600">issue on GitHub ↗</a> to register your interest and get updates on its progress. Once it is released, it will allow you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any supported CloudQuery destination from the list below:</>
                            :
                            <>The CloudQuery {source.name} plugin allows you to sync data from {source.website ? <a href={source.website} className="text-blue-500 hover:text-blue-600">{source.name}</a> : source.name } to any destination. Select the destination you would like to sync {source.name} data from the list of supported destinations below:</>
                        }
                    </p>
                </div>
            </div>

            <div className="mx-4 px-4 mx-auto w-auto">
                <div className="flex justify-left items-left flex-wrap gap-9 pt-8 sm:mt-4">
                    {DESTINATION_PLUGINS.map((plugin) => (
                        <LogoContainer
                            title={plugin.name}
                            href={plugin.href || `/integrations/${source.id}/${plugin.id}`}
                            key={plugin.id}
                            logo={plugin.logo}
                            logoDark={plugin.logoDark}
                            name={plugin.name}
                            external={false}
                            category={plugin.category}
                        >
                        </LogoContainer>
                    ))}
                </div>
            </div>
            <p className="mx-auto mt-24 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                Couldn't find the destination you were looking for? Please <a href="https://github.com/cloudquery/cloudquery/issues/new/choose" target="_blank" className="text-blue-500 hover:text-blue-600">request it on GitHub ↗</a> and we'll add it to our roadmap. New plugins are created all the time, often within days of a request. Showing your interest helps us prioritize.
            </p>
        </div>
    </>;
}
