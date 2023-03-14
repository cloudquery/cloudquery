import {UNPUBLISHED_SOURCE_PLUGINS} from "./pluginData";
import {LogoContainer} from "./LogoContainer";

export default function RequestedSources() {
    return (
        <div className="flex justify-center items-center flex-wrap gap-9 pt-8 sm:mt-4">
            {UNPUBLISHED_SOURCE_PLUGINS.map(plugin => (
                <LogoContainer
                    title={plugin.name}
                    href={`/integrations/${plugin.id}`}
                    key={plugin.id}
                    external={false}
                    logo={plugin.logo}
                    logoDark={plugin.logoDark}
                    name={plugin.name}
                    published={false}
                    category={plugin.category}
                >
                </LogoContainer>
            ))}
        </div>
    );
}
