const path = require("path");
const { promises: fs } = require("fs");

const getIntegrationsPaths = async () => {
  const sources = (await fs.readdir(path.join(process.cwd(), "integrations"), { withFileTypes: true })).filter((file) => !file.isDirectory() && file.name.endsWith(".mdx")).map((file) => path.basename(file.name, ".mdx"))
  const sourceDirs = (await fs.readdir(path.join(process.cwd(), "integrations"), { withFileTypes: true })).filter((dirent) => dirent.isDirectory()).map((dirent) => dirent.name);
  const destinations = await Promise.all(sourceDirs.map((source) => fs.readdir(path.join(process.cwd(), "integrations", source))));
  const destinationsWithSources = destinations.map((destinations, i) => destinations.map((destination) => ({ source: sourceDirs[i], destination: path.basename(destination, ".mdx") }))).flat();
  const paths = [
    ...sources.map((source) => ({ params: { slug: [source] } })),
    ...destinationsWithSources.map(({ source, destination }) => ({ params: { slug: [source, destination] } })),
  ];

  return paths;
};

module.exports = {
    getIntegrationsPaths,
};
  