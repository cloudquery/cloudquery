import react from "react";

const excludedPaths = ["pages/plugins/index.mdx"];
const externalPaths = {
  "pages/plugins/yandexcloud.md":
    "https://github.com/yandex-cloud/cq-provider-yandex/blob/main/docs/index.md",
};

const getGitHubEditLink = (filePath: string) => {
  if (externalPaths[filePath]) {
    return externalPaths[filePath];
  }
  const owner = process.env.VERCEL_GIT_REPO_OWNER || "cloudquery";
  const repo = process.env.VERCEL_GIT_REPO_SLUG || "cloudquery";
  const branch = process.env.VERCEL_GIT_COMMIT_REF || "main";
  const repoPath = `blob/${branch}/${filePath}`;
  const href = `https://github.com/${owner}/${repo}/${repoPath}`
    .replace(/pages\/plugins\/([^/]+?).md/, "plugins/source/$1/docs/index.md")
    .replace(
      /pages\/plugins\/(.*?)\/tables.mdx/,
      "plugins/source/$1/docs/tables"
    )
    .replace("pages/", "website/pages/");

  return href;
};

export function EditLink({
  className,
  filePath,
  children,
}: {
  className?: string;
  filePath?: string;
  children: react.ReactNode;
}) {
  if (!filePath) {
    return null;
  }
  // some files are autogenerated and don't have a GitHub link
  if (excludedPaths.includes(filePath)) {
    return null;
  }

  const href = getGitHubEditLink(filePath);
  return (
    <a href={href} className={className} target="_blank" rel="noreferrer">
      Edit this page on GitHub
    </a>
  );
}
