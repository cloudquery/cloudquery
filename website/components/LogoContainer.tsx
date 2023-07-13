import React from "react";
import {Category, CategoryTitle} from "./Category";
import {Availability} from "./pluginData";

const defaultLogo = "/images/logos/plugins/default.svg";
const defaultLogoDark = "/images/logos/plugins/default-dark.svg";

export const LogoContainer: React.FC<{
    children?: React.ReactNode;
    title: string;
    href: string;
    external: boolean;
    logo: string;
    logoDark: string;
    name: string;
    category?: Category;
    availability?: Availability;
}> = ({ children, title, href = "/", external , logo, logoDark, name, category, availability}) => {
    return (
        <a
            href={href}
            title={title}
            className="flex flex-col w-48 p-6 bg-gray-100 dark:bg-gray-900 items-center justify-center text-gray-600 dark:text-gray-300 dark:hover:text-gray-50 transition ease-in-out hover:scale-105"
            target={external ? "_blank" : undefined}
        >
            {availability === 'premium' ?
                <span className="text-xs font-semibold inline-block py-1 px-2 uppercase rounded text-green-600 bg-green-200 uppercase last:mr-0 mr-1">Premium</span> :
                availability === 'unpublished' ?
                    <span className="text-xs font-semibold inline-block py-1 px-2 uppercase rounded text-green-600 bg-green-200 uppercase last:mr-0 mr-1">Pre-Order</span> :
                    null
            }
            <div className="flex items-center justify-center h-32 mt-6">
                {logo ?
                    <>
                        <img className={"h-full max-h-16" + (logoDark ? " dark:hidden": "")} src={logo}/>
                        {logoDark ? <img className={"h-full max-h-16 hidden dark:block"} src={logoDark}/> : null}
                    </>
                    :
                    <>
                        <img className={"h-full max-h-16 dark:hidden"} src={defaultLogo}/>
                        <img className={"h-full max-h-16 hidden dark:block"} src={defaultLogoDark}/>
                    </>
                }
            </div>
            <div className="flex items-center text-center justify-center h-12 font-bold">
                <p className="item">{name}</p>
            </div>
            <div className="flex items-center text-center justify-center h-12 mt-2 text-slate-400">
                {category ? <p className="item">{CategoryTitle(category)}</p> : null}
            </div>
        </a>
    );
};
