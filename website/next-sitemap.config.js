const { getTablesData } = require("./utils/tables-data");

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
  additionalPaths: () => {
    const tablesData = getTablesData();
    const paths = tablesData.map(({ plugin, table }) => {
      return { loc: `/docs/plugins/sources/${plugin}/tables/${table}`};
    });
    return paths;
  },
};
