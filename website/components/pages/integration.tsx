import copy from "copy-to-clipboard";
import toast, { Toaster } from "react-hot-toast";
import Link from "next/link";
import {Plugin} from "../pluginData";
import {QueriesExamples} from "../QueriesExamples";

export default function Integration({
                                        source,
                                        destination,
                                    }: {
    source?: Plugin;
    destination?: Plugin;
}) {
    const onClick = (code: string) => {
        copy(code);
        toast.success("Copied to clipboard");
    };

    return <>
        <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto lg:max-w-7xl">
            <div className="flex flex-col justify-between md:mr-4">
                <div>
                    <h1 className="max-w-5xl mx-auto nx-text-6xl font-extrabold tracking-tighter leading-[1.1] text-7xl sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
                        Export from&nbsp;
                        <br className="hidden lg:block" />
                        <span className="pr-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
             {source.name}
             </span>
                        &nbsp;to&nbsp;
                        {destination ?
                            <span className="pr-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-green-500 ">
             {destination.name}
             </span> : "any destination"}
                    </h1>
                    <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        CloudQuery is an open-source tool that allows you to export data from any source to any destination.
                    </p>
                    <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
                        {destination ?
                            <div>Let's see how to start syncing data
                                from {source.name} to {destination ? destination.name : "any destination"}.</div>
                            : <div>Let's see how to start syncing data from {source.name} to any
                                destination.</div>
                        }
                    </p>
                </div>
            </div>
        </div>

        <div className="relative from-gray-50 to-gray-100">
            <div className="px-4 py-16 mx-auto sm:pt-20 sm:pb-24 lg:max-w-7xl lg:pt-8">
                <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                    Step 1. Install the CloudQuery CLI
                </h2>
                <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                    The CloudQuery CLI is a command-line tool that allows you to export data from any source to any destination.
                </p>
            </div>
        </div>

        <div className="relative from-gray-50 to-gray-100">
            <div className="px-4 py-16 mx-auto sm:pt-20 sm:pb-24 lg:max-w-7xl lg:pt-8">
                <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                    Step 2. Create the {source.name} {destination ? " to " + destination.name : ""} configuration file
                </h2>
                <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                    The CloudQuery CLI is a command-line tool that allows you to export data from any source to any destination.
                </p>
            </div>
        </div>


        <div className="relative from-gray-50 to-gray-100">
            <div className="px-4 py-16 mx-auto sm:pt-20 sm:pb-24 lg:max-w-7xl lg:pt-8">
                <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl dark:text-white">
                    Step 3. Start the Sync
                </h2>
                <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl">
                    The CloudQuery CLI is a command-line tool that allows you to export data from any source to any destination.
                </p>
            </div>
        </div>

        <Toaster position="bottom-right" />
    </>;
}
