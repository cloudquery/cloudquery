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
  const match = file.history[0].match(/pages\/docs\/plugins\/(.+)\/(.+)\//);
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

function getVersions() {
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

const versions = getVersions();

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
        if (!['badge', 'configuration', 'authentication', 'callout'].includes(data.hName)) {
          return;
        }

        const { kind, name } = getKindAndName(file);
        if (data.hName === "badge") {
          data.hProperties = {
            ...data.hProperties,
            text: "Latest: " + getLatestVersion(kind, name),
          };
          return;
        }
        if (data.hName === "configuration" || data.hName === "authentication" || data.hName === "callout") {
          data.hProperties = {
            ...data.hProperties,
            kind,
            name,
          };
        }
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
