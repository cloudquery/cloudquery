const fs = require("fs");
const path = require("path");
const matter = require("gray-matter");
const title = require("title");

const getData = (pluginsDir, plugin) => {
  try {
    const overviewFile = fs.readFileSync(
      `${pluginsDir}/${plugin}/overview.mdx`,
      "utf-8"
    );
    return {
      id: plugin,
      name: title(plugin),
      stage: "Preview",
      ...matter(overviewFile).data,
    };
  } catch (e) {
    console.warn(`No overview file found for ${plugin}`);
    return { id: plugin, name: title(plugin), stage: "Preview" };
  }
};

const getPluginsData = (type) => {
  const pluginsDir = path.resolve(process.cwd(), `pages/docs/plugins/${type}`);
  const files = fs.readdirSync(pluginsDir, { withFileTypes: true });
  const plugins = files
    .filter((file) => file.isDirectory())
    .map((file) => file.name);

  const withData = plugins.map((plugin) => [
    plugin,
    getData(pluginsDir, plugin),
  ]);
  withData.sort((a, b) => a[1].name.localeCompare(b[1].name));
  return Object.fromEntries(withData);
};

module.exports = {
  getPluginsData,
};
