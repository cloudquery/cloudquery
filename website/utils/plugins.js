const fs = require("fs");
const path = require("path");
const matter = require("gray-matter");
const title = require("title");
const pluginTitle = (plugin) => title(plugin.replace(/-/g, " "));


const tryReadPaths = (paths) => {
  for (const p of paths) {
    try {
      return { content: fs.readFileSync(p, "utf8"), overviewFile: p };
    } catch (e) {
      if (p === paths[paths.length - 1]) {
        throw e;
      }
    }
  }
};

const getData = (plugin, type) => {
  try {
    const typeSingular = type === "sources" ? "source" : "destination";
    const oldsDocsPath = path.resolve(process.cwd(), `pages/docs/plugins/${type}/${plugin}/overview.md`);
    const newDocsPath = path.resolve(process.cwd(), `../plugins/${typeSingular}/${plugin}/docs/overview.md`);
    const { content, overviewFile } = tryReadPaths([oldsDocsPath, newDocsPath]);
    return {
      id: plugin,
      name: pluginTitle(plugin),
      stage: "Preview",
      openInHub: overviewFile === newDocsPath,
      ...matter(content).data,
    };
  } catch (e) {
    console.warn(`No overview file found for ${plugin}`);
    return { id: plugin, name: pluginTitle(plugin), stage: "Preview" };
  }
};

const getPlugins = (type) => {
  const oldFilesPath = fs.readdirSync(path.resolve(process.cwd(), `pages/docs/plugins/${type}`), { withFileTypes: true });
  const pluginsFromOldPath = oldFilesPath
    .filter((file) => file.isDirectory())
    .map((file) => file.name);
  
  const typeSingular = type === "sources" ? "source" : "destination";
  const newFilesPath = fs.readdirSync(path.resolve(process.cwd(), `../plugins/${typeSingular}`), { withFileTypes: true });
  const pluginsFromNewPath = newFilesPath
    .filter((file) => file.isDirectory() && file.name !== "test")
    .map((file) => file.name);


  return Array.from(new Set([...pluginsFromOldPath, ...pluginsFromNewPath]));
}

const getPluginsData = (type) => {
  const plugins = getPlugins(type)
  const withData = plugins.map((plugin) => [
    plugin,
    getData(plugin, type),
  ]);
  withData.sort((a, b) => a[1].name.localeCompare(b[1].name));
  return Object.fromEntries(withData);
};

module.exports = {
  getPluginsData,
};
