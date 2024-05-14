/** @type {import('next-sitemap').IConfig} */
module.exports = {
  siteUrl: "https://docs.cloudquery.io",
  generateRobotsTxt: true,
  exclude: [
    "/buy/*",
    "/landing/*",
    "/docs/plugins/sources/_*",
    "/docs/plugins/destinations/_*",
    "/docs/plugins/sources/*/_*",
    "/docs/plugins/destinations/*/_*",
  ],
  robotsTxtOptions: {
    policies: [
      {
        userAgent: '*',
        allow: '/',
        disallow: [
          // disallow buy redirection pages
          '/buy/*',
          // disallow pages that start with underscores
          "/docs/plugins/sources/_*",
          "/docs/plugins/destinations/_*",
          '/docs/plugins/sources/*/_*',
          '/docs/plugins/destinations/*/_*',
        ],
      },
    ],
  },
};
