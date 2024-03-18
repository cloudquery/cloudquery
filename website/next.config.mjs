import nextra from "nextra";
import * as fs from "fs";
import path from "path";
import { h } from "hastscript";
import remarkDirective from "remark-directive";
import { visit } from "unist-util-visit";

const patterns = {
  cli: /VERSION_(CLI)/,
  sources: /VERSION_SOURCE_([a-zA-Z0-9_]+)/,
  destinations: /VERSION_DESTINATION_([a-zA-Z0-9_]+)/,
};

const pluginNamePatterns = {
  destinationName: /DESTINATION_NAME/,
};

const getKindAndName = (file) => {
  const match = file.history[0].match(/\/plugins\/(.+)\/(.+)\//);
  if(!match) return null;

  const [kind, name] = [match[1], match[2]];
  return {
    kind,
    name,
  };
};

function removeVersionPrefix(version) {
  return version.slice(1);
}

function parseVersion(version) {
  const parts = version.split("-");
  // plugins
  if (parts.length === 4) {
    return parts[3];
  }

  // cli, remove the `v` prefix
  return removeVersionPrefix(parts[1]);
}

function parseName(name) {
  const parts = name.split("-");
  if (parts.length === 2) {
    return parts[1];
  }
  return parts[0];
}

function getVersionsForPrefix(prefix, files) {
  return Object.fromEntries(
    files
      .filter((file) => file.name.split("-")[0] == prefix)
      .map((file) => [parseName(file.name), parseVersion(file.latest)]),
  );
}

function getStaticVersions() {
  const files = fs
    .readdirSync("./versions", { withFileTypes: true })
    .filter((dirent) => dirent.isFile())
    .map((file) => ({
      name: path.basename(file.name, path.extname(file.name)),
      latest: JSON.parse(fs.readFileSync(`./versions/${file.name}`, "utf8"))
        .latest,
    }));

  return {
    cli: getVersionsForPrefix("cli", files),
    sources: getVersionsForPrefix("source", files),
    destinations: getVersionsForPrefix("destination", files),
  };
}

async function getHubVersions() {
  const response = await fetch("https://api.cloudquery.io/plugins?per_page=1000");
  const { items: allPlugins } = await response.json();
  const cloudqueryPlugins = allPlugins.filter((plugin) => plugin.team_name === "cloudquery" && plugin.latest_version);
  const sources = cloudqueryPlugins.filter((plugin) => plugin.kind === "source");
  const destinations = cloudqueryPlugins.filter((plugin) => plugin.kind === "destination");

  return {
    sources: Object.fromEntries(
      sources.map((source) => [source.name, source.latest_version]),
    ),
    destinations: Object.fromEntries(
      destinations.map((destination) => [destination.name, destination.latest_version]),
    ),
  };
};

async function getVersions() {
  const staticVersions = getStaticVersions();
  const hubVersions = await getHubVersions();


  return {
    ...staticVersions,
    sources: {
      ...staticVersions.sources,
      ...hubVersions.sources,
    },
    destinations: {
      ...staticVersions.destinations,
      ...hubVersions.destinations,
    },
  };
}

const versions = await getVersions();

const getLatestVersion = (key, name) => {
  const version = versions[key][name] || "Unpublished";
  return version;
}

const customPlugin = () => {
  return (tree, file) => {
    visit(tree, function (node) {
      if (
        node.type === "containerDirective" ||
        node.type === "leafDirective" ||
        node.type === "textDirective"
      ) {
        const data = node.data || (node.data = {});
        const hast = h(node.name, node.attributes || {});
        data.hName = hast.tagName;
        data.hProperties = hast.properties;
        if (!['badge', 'configuration', 'authentication', 'callout', 'slack-app-link'].includes(data.hName)) {
          return;
        }
        const pluginData = getKindAndName(file);
        if (!pluginData) return;
        const { kind, name } = pluginData;
        if (data.hName === "badge") {
          data.hProperties = {
            text: "Latest: " + getLatestVersion(kind, name),
            ...data.hProperties,
          };
          return;
        }
        data.hProperties = {
          ...data.hProperties,
          kind,
          name,
        };
      }
    });
  };
};

const replaceMdxPluginNames = (node) => {
  if (node.type === "text") {
    Object.entries(pluginNamePatterns).forEach(([key, pattern]) => {
      const match = node.value.match(pattern);
      if (match) {
        node.value = node.value.replace(pattern, "postgresql"); // default to postgresql
      }
    });
  }
  if (node.children !== undefined) {
    node.children.forEach(replaceMdxPluginNames);
  }
  return;
};

const replaceMdxCodeVersions = (node) => {
  if (node.type === "text") {
    Object.entries(patterns).forEach(([key, pattern]) => {
      const match = node.value.match(pattern);
      if (match && match.length >= 1) {
        const name = match[1].toLowerCase();
        const version = getLatestVersion(key, name);
        node.value = node.value.replace(pattern, version);
      }
    });
  }
  if (node.children !== undefined) {
    node.children.forEach(replaceMdxCodeVersions);
  }
  return;
};

const withNextra = nextra({
  defaultShowCopyCode: true,
  theme: "nextra-theme-docs",
  themeConfig: "./theme.config.tsx",
  mdxOptions: {
    remarkPlugins: [remarkDirective, customPlugin],
    rehypePrettyCodeOptions: {
      theme: "nord",
      onVisitLine: (node) => {
        replaceMdxCodeVersions(node);
        replaceMdxPluginNames(node);
      },
    },
  },
});

export default withNextra({
  reactStrictMode: true,
  experimental: {
    legacyBrowsers: false,
    // default is 128KB
    largePageDataBytes: 512 * 1000,
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
      {
        source: "/register-for-cloud",
        permanent: true,
        destination: "https://cloud.cloudquery.io/auth/register",
      }
    ];
  },
  webpack: (config) => {
    config.cache.buildDependencies = {
      versions: fs
        .readdirSync("versions")
        .map((file) => path.resolve("versions", file)),
    };
    return config;
  },
});
