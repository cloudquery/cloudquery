import { getPagesUnderRoute } from "nextra/context";
import Link from "next/link";

interface Page {
  name: string;
  route: string;
  children?: Page[];
  meta?: Record<string, any>;
  frontMatter?: any;
}

function sortByDate(a: Page, b: Page) {
  return (
    new Date(b.frontMatter?.date).getTime() -
    new Date(a.frontMatter?.date).getTime()
  );
}

export function BlogIndex({ more = "Read more" }) {
  return getPagesUnderRoute("/blog")
    .slice()
    .sort(sortByDate)
    .map((page) => {
      return (
        <div key={page.route} className="mb-10">
          <Link
            href={page.route}
            style={{ color: "inherit", textDecoration: "none" }}
            className="block font-semibold mt-8 nx-text-2xl"
            >
            {page.meta?.title || page.meta?.title || page.name}
          </Link>
          <p className="opacity-80" style={{ marginTop: ".5rem" }}>
            {page.meta?.description}{" "}
            <span className="inline-block">
              <Link href={page.route}>{more + " →"}</Link>
            </span>
          </p>
          {page.meta?.date ? (
            <p className="opacity-50 text-sm">{page.meta.date}</p>
          ) : null}
        </div>
      );
    });
}
