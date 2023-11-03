const { getPluginsData } = require("../../../../utils/plugins");

const getMeta = () => {
  const pluginsData = getPluginsData("sources");
  const asMeta = Object.entries(pluginsData).map(
    ([source, { name, openInHub }]) => [
      source,
      {
        title: name,
        ...(openInHub
          ? { href: `https://hub.cloudquery.io/plugins/source/cloudquery/${source}`, "newWindow": true }
          : undefined),
      },
    ],
  );
  const meta = {
    overview: "Overview",
    ...Object.fromEntries(asMeta),
  };
  return meta;
};

module.exports = getMeta();
