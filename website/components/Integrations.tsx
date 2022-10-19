import { INTEGRATIONS } from "./integrationsData";

const LogoContainer: React.FC<{ title: string; href: string }> = ({
  children,
  title,
  href = "/",
}) => {
  return (
    <a
      href={href}
      title={title}
      className="w-9 h-9 flex items-center justify-center text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-50 transition ease-in-out hover:scale-105"
    >
      {children}
    </a>
  );
};

export function Integrations() {
  return (
    <div className="flex justify-center items-center flex-wrap gap-9 mt-8 sm:mt-4">
      {INTEGRATIONS.map(({ name, logo, id }) => (
        <LogoContainer title={name} href={`/docs/plugins/sources`} key={id}>
          {logo}
        </LogoContainer>
      ))}
    </div>
  );
}
