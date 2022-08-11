import { getPagesUnderRoute, Page } from "nextra/context";
import Link from "next/link";

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
          <Link href={page.route}>
            <a
              style={{ color: "inherit", textDecoration: "none" }}
              className="block font-semibold mt-8 text-2xl"
            >
              {page.meta?.title || page.frontMatter?.title || page.name}
            </a>
          </Link>
          <p className="opacity-80" style={{ marginTop: ".5rem" }}>
            {page.frontMatter?.description}{" "}
            <span className="inline-block">
              <Link href={page.route}>{more + " →"}</Link>
            </span>
          </p>
          {page.frontMatter?.date ? (
            <p className="opacity-50 text-sm">{page.frontMatter.date}</p>
          ) : null}
        </div>
      );
    });
}
