import {getPagesUnderRoute} from "nextra/context";
import Link from "next/link";

interface Page {
  name: string;
  route: string;
  children?: Page[];
  meta?: Record<string, any>;
  frontMatter?: any;
}

export function HowToGuideIndex({ more = "Read the guide" }) {
  const howTo = getPagesUnderRoute("/how-to-guides");
  const deploy = getPagesUnderRoute("/docs/deployment");
  // combine howTo and deploy
  const allGuides = [...howTo, ...deploy];
  return allGuides
    .slice()
    .sort((a, b) => {
        const aTitle = a.frontMatter?.title || a.meta?.title || a.name;
        const bTitle = b.frontMatter?.title || b.meta?.title || b.name;
        return aTitle.localeCompare(bTitle);
    })
    .filter((page) => {
        return page.name !== "overview";
    })
    .map((page) => {
      return (
        <div key={page.route} className="mb-10">
          <Link
            href={page.route}
            style={{ color: "inherit", textDecoration: "none" }}
            className="block font-semibold mt-8 nx-text-2xl"
            >
            {page.frontMatter?.title || page.meta?.title || page.name}
          </Link>
          <p className="opacity-80" style={{ marginTop: ".5rem" }}>
            {page.meta?.description}{" "}
            <span className="inline-block">
              <Link href={page.route}>{more + " →"}</Link>
            </span>
          </p>
        </div>
      );
    });
}
