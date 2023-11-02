const { getPluginsData } = require("../../../../utils/plugins");

const getMeta = () => {
  const pluginsData = getPluginsData("destinations");
  const asMeta = Object.entries(pluginsData).map(([destination, { name, openInHub }]) => [
    destination,
    {
      title: name,
      ...(openInHub
        ? { href: `https://hub.cloudquery.io/plugins/destination/cloudquery/${destination}`, "newWindow": true }
        : undefined),
    },
  ]);
  const meta = {
    overview: "Overview",
    ...Object.fromEntries(asMeta),
  };
  return meta;
};

module.exports = getMeta();
