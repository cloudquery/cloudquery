const { getTablesData } = require("./utils/tables-data");

/** @type {import('next-sitemap').IConfig} */
module.exports = {
  siteUrl: "https://www.cloudquery.io",
  generateRobotsTxt: true,
  additionalPaths: () => {
    const tablesData = getTablesData();
    const paths = tablesData.map(({ plugin, table }) => {
      return { loc: `/docs/tables/${plugin}/${table}`};
    });
    return paths;
  },
};
