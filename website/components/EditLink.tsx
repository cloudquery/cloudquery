import react from "react";
import Link from "next/link";

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
    return;
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
    );

  return (
    <a href={href} className={className} target="_blank" rel="noreferrer">
      Edit this page on GitHub
    </a>
  );
}
