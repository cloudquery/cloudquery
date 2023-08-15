import { Avatar } from "./Avatar";
import {AuthorByName} from "../content/authors";
import Head from "next/head";
import { useConfig } from "nextra-theme-docs";
import { Title as CaseStudyTitle } from "./Title";

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
  const authorName = data?.author;

  if (!authorName) {
    return null;
  }

  const authorNames = authorName.split(",").map((name) => name.trim());
  const authors = authorNames.map((name) => AuthorByName(name));

  return (
    <div className="w-full border-b border-gray-400 authors border-opacity-20">
      <div className="flex flex-wrap justify-center py-8 mx-auto gap-7">
        {authors.map((author) => (<Avatar key={author.name} {...author} />))}
      </div>
    </div>
  );
}

export function CaseStudyHeader() {
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
      <CaseStudyTitle data={meta} />
      <Authors data={meta} />
    </>
  );
}
