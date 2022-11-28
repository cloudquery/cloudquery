import nextra from 'nextra'
import * as fs from 'fs';

const reSourcePluginVersion = /VERSION_SOURCE_([a-zA-Z0-9_]+)/;
const reDestPluginVersion = /VERSION_DESTINATION_([a-zA-Z0-9_]+)/;
const reCLI = "VERSION_CLI";

function getVersions() {
  let versions = {
    sources: {},
    destinations: {},
    cli: "",
  }
  const dir = fs.opendirSync('./versions')
  let dirent
  while ((dirent = dir.readSync()) !== null) {
    if (dirent.isFile() && dirent.name.startsWith('source-')) {
      let name = dirent.name.split('-')[1].split('.')[0]
      let {latest} = JSON.parse(fs.readFileSync('./versions/' + dirent.name, 'utf8'))
      versions.sources[name] = latest.split('-')[3]
    } else if (dirent.isFile() && dirent.name.startsWith('destination-')) { 
      let name = dirent.name.split('-')[1].split('.')[0]
      let {latest} = JSON.parse(fs.readFileSync('./versions/' + dirent.name, 'utf8'))
      versions.destinations[name] = latest.split('-')[3]
    } else if (dirent.isFile() && dirent.name == "cli.json") {
      versions.cli = JSON.parse(fs.readFileSync('./versions/' + dirent.name, 'utf8')).latest.split('-')[1]
    }
  }
  dir.closeSync()
  return versions
}

const versions = getVersions()

const replaceMdxCodeVersions = (node) => {
  if (node.type === 'text') {
    let match = node.value.match(reSourcePluginVersion)
    if (match && match.length >= 1) {
      let version = versions.sources[match[1].toLowerCase()]
      if (version === undefined) {
        throw new Error(`Could not find version for source plugin ${match[1]}`)
      }
      node.value = node.value.replace(reSourcePluginVersion, version)
    }
    match = node.value.match(reDestPluginVersion)
    if (match && match.length >= 1) {
      let version = versions.destinations[match[1].toLowerCase()]
      if (version === undefined) {
        throw new Error(`Could not find version for destination plugin ${match[1]}`)
      }
      node.value = node.value.replace(reDestPluginVersion, version)
    }
    if (node.value.includes(reCLI)) {
      let version = versions.cli
      if (version === undefined) {
        throw new Error(`Could not find version for cli ${match}`)
      }
      node.value = node.value.replace(reCLI, version)
    }
  }
  if (node.children !== undefined) {
    node.children.map(replaceMdxCodeVersions)
  }
  return
}

const withNextra = nextra({
  theme: "nextra-theme-docs",
  themeConfig: "./theme.config.tsx",
  unstable_flexsearch: true,
  unstable_staticImage: true,
  mdxOptions: {
    rehypePrettyCodeOptions: {
      theme: 'nord',
      onVisitLine: (node) => {
        replaceMdxCodeVersions(node)
      },
    },
  },
});

export default withNextra({
  reactStrictMode: true,
  experimental: {
    legacyBrowsers: false,
    images: { allowFutureImage: true },
  },
  env: {
    VERCEL_GIT_REPO_OWNER: process.env.VERCEL_GIT_REPO_OWNER,
    VERCEL_GIT_REPO_SLUG: process.env.VERCEL_GIT_REPO_SLUG,
    VERCEL_GIT_COMMIT_REF: process.env.VERCEL_GIT_COMMIT_REF,
  },
  async redirects() {
    return [
      {
        source: "/docs/changelog",
        permanent: true,
        destination: "https://github.com/cloudquery/cloudquery/releases",
      },
    ];
  },
});
