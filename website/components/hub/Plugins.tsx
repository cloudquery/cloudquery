import type { Plugins as PluginsDetails } from "../../content/hub";

export const Plugins = ({ plugins }: { plugins: PluginsDetails }) => {
  return (
    <div>
      {plugins.map(({ id, name }) => (
        <div key={id}>{name}</div>
      ))}
    </div>
  );
};
