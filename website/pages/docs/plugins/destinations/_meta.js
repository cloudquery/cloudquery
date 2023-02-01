const { getPluginsData } = require("../../../../utils/plugins");

const getMeta = () => {
  const pluginsData = getPluginsData("destinations");
  const asMeta = Object.entries(pluginsData).map(([destination, { name }]) => [destination, name]);
  const meta = {
    overview: "Overview",
    ...Object.fromEntries(asMeta),
  }
  return meta;
};

module.exports = getMeta();
