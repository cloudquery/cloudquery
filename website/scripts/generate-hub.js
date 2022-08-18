const path = require("path");
const { promises: fs } = require("fs");

const PLUGINS_DATA = {
  aws: { name: "AWS" },
  azure: { name: "Azure" },
  cloudflare: { name: "CloudFlare" },
  digitalocean: { name: "DigitalOcean" },
  fuzz: { name: "Fuzz", ignore: true },
  gcp: { name: "GCP" },
  github: { name: "GitHub" },
  k8s: { name: "Kubernetes" },
  okta: { name: "Okta" },
  terraform: { name: "Terraform" },
};

const PLUGINS_SOURCE = `${__dirname}/../../plugins/source`;
const HUB_PATH = `${__dirname}/../pages/hub`;
const PLUGINS_PATH = `${HUB_PATH}/plugins`;

const NAME_PLACEHOLDER = "NAME_PLACEHOLDER";
const CONTENT_PLACEHOLDER = "CONTENT_PLACEHOLDER";

const PLUGIN_TEMPLATE = `---
title: ${NAME_PLACEHOLDER} Plugin
---

${CONTENT_PLACEHOLDER}
`;

const PLUGINS_PLACEHOLDER = "PLUGINS_PLACEHOLDER";

const PLUGINS_TEMPLATE = `---
title: Plugins
---

import { Plugins } from "../../components/hub/Plugins";

# Plugins

<Plugins plugins={${PLUGINS_PLACEHOLDER}} />
`;

const TABLES_TEMPLATE = `---
title: ${NAME_PLACEHOLDER} Plugin Tables
---

|${NAME_PLACEHOLDER} Plugin Tables|
|---|
${CONTENT_PLACEHOLDER}
`;

const getPlugins = async () => {
  const pluginsDirent = await fs.readdir(PLUGINS_SOURCE, {
    withFileTypes: true,
  });
  const plugins = pluginsDirent
    .filter(
      (dirent) => dirent.isDirectory() && !PLUGINS_DATA[dirent.name].ignore
    )
    .map((dirent) => ({
      id: dirent.name,
      name: PLUGINS_DATA[dirent.name].name,
    }));
  return plugins.sort((plugin1, plugin2) =>
    plugin1.id.localeCompare(plugin2.id)
  );
};

const getPluginOverview = async (plugin) => {
  const overview = await fs.readFile(
    `${PLUGINS_SOURCE}/${plugin.id}/docs/index.md`,
    { encoding: "utf8" }
  );
  return overview;
};

const generatePluginsPagesMeta = async (plugins) => {
  const pluginsMeta = Object.fromEntries(
    plugins.map((plugin) => [plugin.id, plugin.name])
  );
  await fs.writeFile(
    `${PLUGINS_PATH}/meta.json`,
    JSON.stringify(pluginsMeta, null, 2)
  );
};

const generatePluginsPages = async (plugins) => {
  const overviews = await Promise.all(
    plugins.map((plugin) => getPluginOverview(plugin))
  );
  const pages = plugins.map((plugin, index) => ({
    path: `${PLUGINS_PATH}/${plugin.id}.md`,
    content: PLUGIN_TEMPLATE.replace(NAME_PLACEHOLDER, plugin.name).replace(
      CONTENT_PLACEHOLDER,
      overviews[index]
    ),
  }));
  await Promise.all(
    pages.map(async (page) => {
      await fs.mkdir(path.dirname(page.path), { recursive: true });
      await fs.writeFile(page.path, page.content, { encoding: "utf8" });
    })
  );
  await generatePluginsPagesMeta(plugins);
};

const getPluginAsString = (plugin) => {
  const asString = Object.entries(plugin)
    .map(([key, value]) => `${key}: "${value}"`)
    .join(", ");
  return `{${asString}}`;
};

const generatePluginsIndexPage = async (plugins) => {
  const pluginsArray = plugins
    .map((plugin) => getPluginAsString(plugin))
    .join(", ");

  const page = {
    content: PLUGINS_TEMPLATE.replace(PLUGINS_PLACEHOLDER, `[${pluginsArray}]`),
    path: `${HUB_PATH}/plugins.mdx`,
  };
  await fs.mkdir(path.dirname(page.path), { recursive: true });
  await fs.writeFile(page.path, page.content, { encoding: "utf8" });
};

const generatePluginTablePages = async (plugin) => {
  const tablesSource = `${PLUGINS_SOURCE}/${plugin.id}/docs/tables`;
  const tablesList = (await fs.readdir(tablesSource, { withFileTypes: true }))
    .filter((dirent) => dirent.isFile() && dirent.name.endsWith(".md"))
    .map((dirent) => path.basename(dirent.name, path.extname(dirent.name)));

  const tablesPath = `${PLUGINS_PATH}/${plugin.id}/tables.mdx`;
  const tablesLinks = tablesList
    .map(
      (table) =>
        `|[${table}](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/${plugin.id}/docs/tables/${table}.md)|`
    )
    .join("\n");
  const content = TABLES_TEMPLATE.replaceAll(
    NAME_PLACEHOLDER,
    plugin.name
  ).replace(CONTENT_PLACEHOLDER, tablesLinks);
  await fs.mkdir(path.dirname(tablesPath), { recursive: true });
  await fs.writeFile(tablesPath, content, { encoding: "utf8" });
};

const generatePluginsTablesPages = async (plugins) => {
  await Promise.all(
    plugins.map(async (plugin) => generatePluginTablePages(plugin))
  );
};

const main = async () => {
  await fs.rm(PLUGINS_PATH, { recursive: true, force: true });

  const plugins = await getPlugins();
  await generatePluginsIndexPage(plugins);
  await generatePluginsPages(plugins);
  await generatePluginsTablesPages(plugins);
};

main();
