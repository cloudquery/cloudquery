import { SOURCE_PLUGINS, DESTINATION_PLUGINS } from "./pluginData";

const LogoContainer: React.FC<{
  children?: React.ReactNode;
  title: string;
  href: string;
  external: boolean;
  logo: string;
  logoDark: string;
  name: string;
}> = ({ children, title, href = "/", external , logo, logoDark, name}) => {
  return (
      <a
          href={href}
          title={title}
          className="flex flex-col h-48 w-48 p-6 bg-gray-100 dark:bg-gray-900 items-center justify-center text-gray-600 dark:text-gray-300 dark:hover:text-gray-50 transition ease-in-out hover:scale-105"
          target={external ? "_blank" : undefined}
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

export function Integrations() {
  return (
    <div className="flex justify-center items-center flex-wrap gap-9 pt-8 sm:mt-4">
      {SOURCE_PLUGINS.map(({ name, logo, logoDark, id, href }) => (
        <LogoContainer
          title={name}
          href={href || `/integrations/${id}`}
          key={id}
          external={Boolean(href)}
          logo={logo}
          logoDark={logoDark}
          name={name}
        >
        </LogoContainer>
      ))}
    </div>
  );
}


export function Destinations() {
  return (
      <div className="flex justify-center items-center flex-wrap gap-9 pt-8 sm:mt-4">
        {DESTINATION_PLUGINS.map(({ name, logo, logoDark, id, href }) => (
            <LogoContainer
                title={name}
                href={href || `/docs/plugins/destinations/${id}/overview`}
                key={id}
                external={Boolean(href)}
                logo={logo}
                logoDark={logoDark}
                name={name}
            >
            </LogoContainer>
        ))}
      </div>
  );
}
