const path = require("path");
const { promises: fs } = require("fs");

const PLUGINS_DATA = {
  aws: { name: "AWS" },
  azure: { name: "Azure" },
  cloudflare: { name: "CloudFlare" },
  digitalocean: { name: "DigitalOcean" },
  gcp: { name: "GCP" },
  github: { name: "GitHub" },
  heroku: { name: "Heroku" },
  k8s: { name: "Kubernetes" },
  okta: { name: "Okta" },
  terraform: { name: "Terraform" },
  yandexcloud: {
    name: "Yandex Cloud",
    external: true,
    url: "https://github.com/yandex-cloud/cq-provider-yandex",
  },
};

const PLUGINS_SOURCE = `${__dirname}/../../plugins/source`;
const PLUGINS_PATH = `${__dirname}/../pages/plugins`;

const NAME_PLACEHOLDER = "NAME_PLACEHOLDER";
const CONTENT_PLACEHOLDER = "CONTENT_PLACEHOLDER";
const EXTERNAL_PLUGINS_LINK_PLACEHOLDER = "EXTERNAL_PLUGINS_LINK_PLACEHOLDER";

const PLUGIN_TEMPLATE = `---
title: ${NAME_PLACEHOLDER} Plugin
---

${CONTENT_PLACEHOLDER}
`;

const PLUGINS_TEMPLATE = `---
title: Overview
---

import { Plugins } from '../../components/Plugins'

# CloudQuery Plugins

Discover plugins that power CloudQuery, the open source high performance data integration platform designed for security and infrastructure teams.

<Plugins />
`;

const TABLES_TEMPLATE = `---
title: ${NAME_PLACEHOLDER} Plugin Tables
---

|${NAME_PLACEHOLDER} Plugin Tables|
|---|
${CONTENT_PLACEHOLDER}
`;

const EXTERNAL_PLUGIN_TEMPLATE = `
# ${NAME_PLACEHOLDER} Plugin

The CloudQuery ${NAME_PLACEHOLDER} plugin pulls configuration out of ${NAME_PLACEHOLDER} resources, normalizes them and stores them in PostgreSQL.

For more details see [${NAME_PLACEHOLDER} Plugin repository](${EXTERNAL_PLUGINS_LINK_PLACEHOLDER})
`;

const getPlugins = async () => {
  return Object.entries(PLUGINS_DATA)
    .map(([id, plugin]) => ({
      id,
      ...plugin,
    }))
    .sort((a, b) => a.name.localeCompare(b.name));
};

const getPluginOverview = async (plugin) => {
  if (plugin.external) {
    return EXTERNAL_PLUGIN_TEMPLATE.replace(
      EXTERNAL_PLUGINS_LINK_PLACEHOLDER,
      plugin.url
    ).replaceAll(NAME_PLACEHOLDER, plugin.name);
  }

  const overview = await fs.readFile(
    path.resolve(`${PLUGINS_SOURCE}/${plugin.id}/docs/index.md`),
    { encoding: "utf8" }
  );
  return overview;
};

const generatePluginsPagesMeta = async (plugins) => {
  const pluginsMeta = Object.fromEntries(
    plugins.map((plugin) => [plugin.id, plugin.name])
  );
  await fs.writeFile(
    `${PLUGINS_PATH}/_meta.json`,
    JSON.stringify({ index: "Overview", ...pluginsMeta }, null, 2)
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

const generatePluginsIndexPage = async () => {
  const page = {
    content: PLUGINS_TEMPLATE,
    path: `${PLUGINS_PATH}/index.mdx`,
  };
  await fs.mkdir(path.dirname(page.path), { recursive: true });
  await fs.writeFile(page.path, page.content, { encoding: "utf8" });
};

const generatePluginTablePages = async (plugin) => {
  const sourcePlugin = `${PLUGINS_SOURCE}/${plugin.id}`;
  const tablesDir = `${PLUGINS_PATH}/${plugin.id}/tables`;
  await fs.mkdir(tablesDir, { recursive: true });
  await fs.cp(`${sourcePlugin}/docs/tables`, tablesDir, { recursive: true });
  const tablesIndexPath = `${PLUGINS_PATH}/${plugin.id}/tables.mdx`;
  const tablesList = (await fs.readdir(tablesDir, { withFileTypes: true }))
    .filter((dirent) => dirent.isFile() && dirent.name.endsWith(".md"))
    .map((dirent) => path.basename(dirent.name, path.extname(dirent.name)));

  await generatePluginTablePagesMeta(plugin.id, tablesList);

  const tablesLinks = tablesList
    .map((table) => `|[${table}](tables/${table})|`)
    .join("\n");
  const content = TABLES_TEMPLATE.replaceAll(
    NAME_PLACEHOLDER,
    plugin.name
  ).replace(CONTENT_PLACEHOLDER, tablesLinks);
  await fs.writeFile(tablesIndexPath, content, { encoding: "utf8" });
};

const generatePluginTablePagesMeta = async (plugin, tablesList) => {
  const tablesMeta = Object.fromEntries(
    tablesList.map((table) => [table, table.toLowerCase()])
  );
  await fs.writeFile(
    `${PLUGINS_PATH}/${plugin}/tables/_meta.json`,
    JSON.stringify({ ...tablesMeta }, null, 2)
  );
};

const generatePluginsTablesPages = async (plugins) => {
  await Promise.all(
    plugins
      .filter((plugin) => !plugin.external)
      .map(async (plugin) => generatePluginTablePages(plugin))
  );
};

const main = async () => {
  await fs.rm(PLUGINS_PATH, { recursive: true, force: true });

  const plugins = await getPlugins();
  await generatePluginsIndexPage();
  await generatePluginsPages(plugins);
  await generatePluginsTablesPages(plugins);
};

main();
