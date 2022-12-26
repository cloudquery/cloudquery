import { Avatar } from "./Avatar";
import CLOUDQUERY_TEAM from "../content/team";
import type { Author } from "../content/team";
import Head from "next/head";
import { useConfig } from "nextra-theme-docs";

type PostMeta = {
  title: string;
  /** security */
  tag: string;
  description: string;
  /** yevgenypats */
  author?: string;

  /** /images/og-image.png */
  ogImage?: string;
};

function Authors({ data }: { data: PostMeta }) {
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

function GuideTitle({ data }: { data: PostMeta }) {
  const title = data.title;

  if (!title) {
    return null;
  }

  return <h1>{title}</h1>;
}

export function HowToGuideHeader() {
  const config = useConfig();
  const meta = config.frontMatter as PostMeta;

  if (!meta) {
    return null;
  }

  const image = `/og-image/${meta.title}`;

  return (
    <>
      <Head>
        <meta property="twitter:image" content={image} />
        <meta property="og:image" content={image} />
      </Head>
      <GuideTitle data={meta} />
      <Authors data={meta} />
    </>
  );
}
