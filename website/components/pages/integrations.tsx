import {Sources} from "../Sources";
import {Destinations} from "../Destinations";

export default function Integrations() {
    return <>
        <div className="max-w-5xl px-4 pb-12 mx-auto lg:px-8">
            <div className="sm:py-20 lg:py-24">
                <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">
                    <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
                        Source Integrations
                    </h2>
                    <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        Integrate with a growing list of <a className="dark:text-white" href="/docs/plugins/sources/overview">30+ cloud providers and SaaS apps</a> with more than 1,000 unique tables. Sync to your <a href="/docs/plugins/destinations/overview" className="dark:text-white">favorite database, data warehouse or data lake</a>.
                    </p>
                    <Sources />
                </div>
                <div className="max-w-8xl px-4 pb-12 mx-auto lg:px-8 ">
                    <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
                        Destinations
                    </h2>
                    <p className="mx-auto mt-4 font-medium text-lg text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
                        CloudQuery currently supports the following destination databases, data warehouses and data lakes:
                    </p>
                    <Destinations />
                </div>
            </div>
        </div>
    </>;
}
