import RequestedSources from "../RequestedSources";

export default function NewSourcePlugin() {
    return <>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="sm:py-20 lg:py-24">
                <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">
                    <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
                        Request a new Source Plugin
                    </h2>
                    <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Find your favorite cloud provider, database or SaaS app in the list of in-progress or requested plugins below.
                    </p>
                    <RequestedSources />
                    <p className="mx-auto mt-24 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Couldn't find the integration you were looking for? Please <a href="https://github.com/cloudquery/cloudquery/issues/new/choose" className="text-blue-500 hover:text-blue-600">request it on GitHub</a> and we'll add it to our roadmap. New plugins are created all the time, often within days of a request. Showing your interest in a plugin helps us prioritize it.
                    </p>
                </div>
            </div>
        </div>
    </>;
}
