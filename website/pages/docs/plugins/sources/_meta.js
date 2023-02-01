const { getPluginsData } = require("../../../../utils/plugins");

const getMeta = () => {
  const pluginsData = getPluginsData("sources");
  const asMeta = Object.entries(pluginsData).map(([source, { name }]) => [source, name]);
  const meta = {
    overview: "Overview",
    ...Object.fromEntries(asMeta),
  }
  return meta;
};

module.exports = getMeta();
