import nextra from "nextra";
import * as fs from "fs";
import path from "path";

const patterns = {
  cli: /VERSION_(CLI)/,
  sources: /VERSION_SOURCE_([a-zA-Z0-9_]+)/,
  destinations: /VERSION_DESTINATION_([a-zA-Z0-9_]+)/,
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

function isUnreleasedPlugin(key, name) {
  if (key === "cli") {
    return false;
  }
  const pluginDir = path.resolve("..", "plugins", key.slice(0, -1), name);
  try {
    // check that this is a plugin
    fs.accessSync(path.resolve(pluginDir, "go.mod"), fs.constants.F_OK);
  } catch (e) {
    return false;
  }

  const changelogPath = path.resolve(pluginDir, "CHANGELOG.md");
  try {
    const changelogContent = fs.readFileSync(changelogPath, "utf8");
    const emptyChangelog = !changelogContent.includes("1.0.0");
    return emptyChangelog;
  } catch (err) {
    // no Changelog, this is new plugin
    return true;
  }
}

function getVersionsForPrefix(prefix, files) {
  return Object.fromEntries(
    files
      .filter((file) => file.name.split("-")[0] == prefix)
      .map((file) => [parseName(file.name), parseVersion(file.latest)])
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

const replaceMdxCodeVersions = (node) => {
  if (node.type === "text") {
    Object.entries(patterns).forEach(([key, pattern]) => {
      const match = node.value.match(pattern);
      if (match && match.length >= 1) {
        const name = match[1].toLowerCase();
        const version = versions[key][name];
        if (version !== undefined || isUnreleasedPlugin(key, name)) {
          node.value = node.value.replace(pattern, version || "v1.0.0");
        } else {
          throw new Error(`Could not find version for ${key} ${name}`);
        }
      }
    });
  }
  if (node.children !== undefined) {
    node.children.forEach(replaceMdxCodeVersions);
  }
  return;
};

const withNextra = nextra({
  theme: "nextra-theme-docs",
  themeConfig: "./theme.config.tsx",
  unstable_staticImage: true,
  mdxOptions: {
    rehypePrettyCodeOptions: {
      theme: "nord",
      onVisitLine: (node) => {
        replaceMdxCodeVersions(node);
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
  webpack: (config) => {
    config.cache.buildDependencies = {
      versions: fs
        .readdirSync("versions")
        .map((file) => path.resolve("versions", file)),
    };
    return config;
  },
});
