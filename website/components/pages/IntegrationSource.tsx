import {DESTINATION_PLUGINS, Plugin, SOURCE_PLUGINS} from "../pluginData";
import React, {ReactNode} from "react";

const LogoContainer: React.FC<{
    children?: React.ReactNode;
    title: string;
    href: string;
    logo: string;
    logoDark: string;
    name: string;
}> = ({ children, title, href = "/", logo, logoDark, name}) => {
    return (
        <a
            href={href}
            title={title}
            className="flex flex-col h-48 w-48 p-6 bg-gray-100 dark:bg-gray-900 items-center justify-center text-gray-600 dark:text-gray-300 dark:hover:text-gray-50 transition ease-in-out hover:scale-105"
        >
            <div className="flex items-center justify-center h-32 mt-6">
                <img className={"h-full max-h-16" + (logoDark ? " dark:hidden": "")} src={logo}/>
                {logoDark ? <img className="h-full max-h-16 hidden dark:block" src={logoDark}/> : null}
            </div>
            <div className="flex items-center text-center justify-center h-16 mt-6">
                <p className="item">{name}</p>
            </div>
        </a>
    );
};

export default function Integration({
    source
                                    }: {
    source: Plugin;
}) {
    return <>
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
                    CloudQuery is an open-source ELT platform that allows you to extract data from any source into any destination. CloudQuery is a free and open source alternative to Fivetran or Airbyte that requires no account, and it takes only minutes to get started.
                </p>
                <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                    The CloudQuery {source.name} plugin allows you to sync data from {source.name} to any destination. Select the destination you would like to sync {source.name} data from the list of supported destinations below:
                </p>
            </div>
        </div>

        <div className="mx-4 px-4 mx-auto w-auto">
            <div className="flex justify-left items-left flex-wrap gap-9 pt-8 sm:mt-4 max-w-2xl">
                {DESTINATION_PLUGINS.map(({ name, logo, logoDark, id, href }) => (
                    <LogoContainer
                        title={name}
                        href={href || `/integrations/${source.id}/${id}`}
                        key={id}
                        logo={logo}
                        logoDark={logoDark}
                        name={name}
                    >
                    </LogoContainer>
                ))}
            </div>
        </div>
    </>;
}
