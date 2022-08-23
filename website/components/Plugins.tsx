import Feature from "./Feature";
import { INTEGRATIONS } from "./integrationsData";

export const Plugins = () => {
  return (
    <div className="grid grid-cols-1 mt-12 gap-x-6 gap-y-12 sm:grid-cols-2 lg:mt-16 lg:grid-cols-3 lg:gap-x-8 lg:gap-y-12">
      {INTEGRATIONS.map(({ name, id, logo }) => (
        <a
          href={`/plugins/${id}`}
          key={id}
          className="group no-underline text-gray-600 group-hover:text-gray-900 dark:text-gray-300 dark:group-hover:text-gray-50"
        >
          <Feature
            feature={{
              name: name,
              page: "docs",
              Icon: () => (
                <div className="w-9 h-9 flex items-center justify-center transition ease-in-out group-hover:scale-105">
                  {logo}
                </div>
              ),
              description: "",
            }}
            detailed={false}
          />
        </a>
      ))}
    </div>
  );
};
