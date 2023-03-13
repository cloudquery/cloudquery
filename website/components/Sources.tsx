import {SOURCE_PLUGINS, UNPUBLISHED_SOURCE_PLUGINS} from "./pluginData";

import {LogoContainer} from "./LogoContainer";

export function Sources() {
  return (
    <div className="flex justify-center items-center flex-wrap gap-9 pt-8 sm:mt-4">
       {SOURCE_PLUGINS.map(plugin => (
            <LogoContainer
                title={plugin.name}
                href={`/integrations/${plugin.id}`}
                key={plugin.id}
                external={false}
                logo={plugin.logo}
                logoDark={plugin.logoDark}
                name={plugin.name}
                published={true}
                category={plugin.category}
            >
            </LogoContainer>
       ))}
        <div className="hidden">
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
      <LogoContainer
          title={"Request a Source Plugin"}
          href={`/docs/plugins/overview`}
          key={"more"}
          external={false}
          logo={"/images/logos/plugins/more.svg"}
          logoDark={"/images/logos/plugins/more-dark.svg"}
          name={"Request a Source Plugin"}
      >
      </LogoContainer>
    </div>
  );
}
