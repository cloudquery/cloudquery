import { Avatar } from "./Avatar";
import Head from "next/head";
import { useConfig } from "nextra-theme-docs";
import {AuthorByName} from "../content/authors";
import { Title as BlogTitle } from "./Title";

type BlogPostMeta = {
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

function Authors({ data }: { data: BlogPostMeta }) {
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

function BlogDate({ data }: { data: BlogPostMeta }) {
  const date = data.date;
  if (!date) {
    return null;
  }
  const options: Intl.DateTimeFormatOptions = {year: 'numeric', month: 'long', day: 'numeric' };
  let dateObject = new Date(date);
  let formattedDate = dateObject.toLocaleDateString("en-US", options);
  return <div className="text-center">
    <p className="opacity-50 text-sm">{formattedDate}</p>
  </div>
}

export function BlogHeader() {
  const config = useConfig();
  const meta = config.frontMatter as BlogPostMeta;

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
      <BlogTitle data={meta} />
      <BlogDate data={meta} />
      <Authors data={meta} />
    </>
  );
}
