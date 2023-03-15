import {DESTINATION_PLUGINS} from "./pluginData";
import {LogoContainer} from "./LogoContainer";

export function Destinations() {
    return (
        <div className="flex justify-center items-center flex-wrap gap-9 pt-8 sm:mt-4">
            {DESTINATION_PLUGINS.map(plugin => (
                <LogoContainer
                    title={plugin.name}
                    href={plugin.href || `/docs/plugins/destinations/${plugin.id}/overview`}
                    key={plugin.id}
                    external={Boolean(plugin.href)}
                    logo={plugin.logo}
                    logoDark={plugin.logoDark}
                    name={plugin.name}
                    category={plugin.category}
                >
                </LogoContainer>
            ))}
        </div>
    );
}