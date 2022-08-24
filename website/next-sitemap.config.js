/** @type {import('next-sitemap').IConfig} */
module.exports = {
  siteUrl: "https://www.cloudquery.io",
  generateRobotsTxt: true,
  robotsTxtOptions: {
    additionalSitemaps: [
      "https://hub.cloudquery.io/sitemap.xml",
      "https://hub.cloudquery.io/hub-sitemap.xml",
    ],
  },
};
