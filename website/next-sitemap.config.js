const { getTablesData } = require("./utils/tables-data");
const { getIntegrationsPaths } = require("./utils/integrations");

/** @type {import('next-sitemap').IConfig} */
module.exports = {
  siteUrl: "https://www.cloudquery.io",
  generateRobotsTxt: true,
  exclude: [
    "/docs/plugins/sources/*/_*",
    "/docs/plugins/destinations/*/_*",
  ],
  robotsTxtOptions: {
    policies: [
      {
        userAgent: '*',
        allow: '/',
        disallow: [
          // disallow pages that start with underscores
          '/docs/plugins/sources/*/_*',
          '/docs/plugins/destinations/*/_*',
        ],
      },
    ],
  },
  additionalPaths: async () => {
    const tablesData = getTablesData();
    const tablesPaths = tablesData.map(({ plugin, table }) => {
      return { loc: `/docs/plugins/sources/${plugin}/tables/${table}`};
    });
    const integrationsData = await getIntegrationsPaths();
    const integrationsPaths = integrationsData.map(({ params: { slug } }) => {
      return { loc: `/integrations/${slug.join("/")}`};
    })
    return [...tablesPaths, ...integrationsPaths];
  },
};
