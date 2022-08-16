import { Avatar } from "./Avatar";
import CLOUDQUERY_TEAM from "../content/team";
import { getAllPages } from "nextra/context";
import type { Author } from "../content/team";
import Head from "next/head";

type BlogPosFrontMatter = {
  title: string;
  /** security */
  tag: string;
  /** 2022/06/06 */
  date: string;
  description: string;
  /** yevgenypats */
  author?: string;

  /** /images/og-image.png */
  ogImage?: string;
};

function Authors({ data }: { data: BlogPosFrontMatter }) {
  const authorName = data?.author as Author;

  if (!authorName) {
    return null;
  }

  const author = CLOUDQUERY_TEAM[authorName];

  return (
    <div className="w-full border-b border-gray-400 authors border-opacity-20">
      <div className="flex flex-wrap justify-center py-8 mx-auto gap-7">
        <Avatar {...author} />
      </div>
    </div>
  );
}

function BlogTitle({ data }: { data: BlogPosFrontMatter }) {
  const title = data.title;

  if (!title) {
    return null;
  }

  return <h1>{title}</h1>;
}

export function BlogHeader() {
  const routes = getAllPages();
  const currentRoute = globalThis.__nextra_internal__.route;

  const blogPage = routes.find(
    ({ name, children }) => name === "blog" && Boolean(children)
  );

  const currentPage = blogPage.children?.find(
    ({ route }) => route === currentRoute
  );

  const frontMatter = currentPage?.frontMatter as BlogPosFrontMatter;

  if (!frontMatter) {
    return null;
  }

  const image = `https://docs.cloudquery.io/og-image/${frontMatter.title}`;

  return (
    <>
      <Head>
        <meta property="twitter:image" content={image} />
        <meta property="og:image" content={image} />
      </Head>
      <BlogTitle data={frontMatter} />
      <Authors data={frontMatter} />
    </>
  );
}
