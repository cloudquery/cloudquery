import { getPagesUnderRoute } from "nextra/context";
import Link from "next/link";

interface Page {
  name: string;
  route: string;
  children?: Page[];
  meta?: Record<string, any>;
  frontMatter?: any;
}

export function CaseStudiesIndex({ more = "Read the case study" }) {
  return getPagesUnderRoute("/case-studies")
    .slice()
    .map((page) => {
      // small hack to change it in the first case study
      let moreDescription = "Read the case study";
      if (page.route === "/case-studies/add-your-case-study") {
        moreDescription = "Add your case study";
      }
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
              <Link href={page.route}>{moreDescription + " →"}</Link>
            </span>
          </p>
        </div>
      );
    });
}
