import { INTEGRATIONS } from "./integrationsData";

const LogoContainer: React.FC<{
  children?: React.ReactNode;
  title: string;
  href: string;
  external: boolean;
}> = ({ children, title, href = "/", external }) => {
  return (
    <a
      href={href}
      title={title}
      className="w-9 h-9 flex items-center justify-center text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-50 transition ease-in-out hover:scale-105"
      target={external ? "_blank" : undefined}
    >
      {children}
    </a>
  );
};

export function Integrations() {
  return (
    <div className="flex justify-center items-center flex-wrap gap-9 mt-8 sm:mt-4">
      {INTEGRATIONS.map(({ name, logo, id, href }) => (
        <LogoContainer
          title={name}
          href={href || `/docs/plugins/sources/${id}/overview`}
          key={id}
          external={Boolean(href)}
        >
          {logo}
        </LogoContainer>
      ))}
    </div>
  );
}
