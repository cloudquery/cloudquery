import { getPagesUnderRoute } from "nextra/context";
import Link from "next/link";

interface Page {
  name: string;
  route: string;
  children?: Page[];
  meta?: Record<string, any>;
  frontMatter?: any;
}

export function HowToGuideIndex({ more = "Read the guide" }) {
  return getPagesUnderRoute("/how-to-guides")
    .slice()
    .map((page) => {
      return (
        <div key={page.route} className="mb-10">
          <Link href={page.route}>
            <a
              style={{ color: "inherit", textDecoration: "none" }}
              className="block font-semibold mt-8 nx-text-2xl"
            >
              {page.meta?.title || page.meta?.title || page.name}
            </a>
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
