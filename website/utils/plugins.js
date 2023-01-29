const fs = require("fs");
const path = require("path");
const matter = require("gray-matter");
const title = require("title");

const getData = (pluginsDir, destination) => {
  try {
    const overviewFile = fs.readFileSync(
      `${pluginsDir}/${destination}/overview.mdx`,
      "utf-8"
    );
    return {
      id: destination,
      name: title(destination),
      stage: "Preview",
      ...matter(overviewFile).data,
    };
  } catch (e) {
    console.warn(`No overview file found for ${destination}`);
    return { id: destination, name: title(destination), stage: "Preview" };
  }
};

const getPluginsData = (type) => {
  const pluginsDir = path.resolve(process.cwd(), `pages/docs/plugins/${type}`);
  const files = fs.readdirSync(pluginsDir, { withFileTypes: true });
  const destinations = files
    .filter((file) => file.isDirectory())
    .map((file) => file.name);

  const withData = destinations.map((destination) => [
    destination,
    getData(pluginsDir, destination),
  ]);
  withData.sort((a, b) => a[1].name.localeCompare(b[1].name));
  return Object.fromEntries(withData);
};

module.exports = {
  getPluginsData,
};
