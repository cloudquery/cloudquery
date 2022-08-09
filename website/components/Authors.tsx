import { Avatar } from "./Avatar";
import cn from "classnames";
import CLOUDQUERY_TEAM from "../content/team";
import { getAllPages } from "nextra/context";
import type { Author } from "../content/team";

export function Authors() {
  const routes = getAllPages();
  const currentRoute = globalThis.__nextra_internal__.route;

  const blogPage = routes.find(
    ({ name, children }) => name === "blog" && Boolean(children)
  );

  const currentPage = blogPage.children?.find(
    ({ route }) => route === currentRoute
  );

  const authorName: Author = currentPage?.frontMatter?.author;

  if (!authorName) {
    return null;
  }

  const author = CLOUDQUERY_TEAM[authorName];

  return (
    <div className="w-full border-b border-gray-400 authors border-opacity-20">
      <div className={cn("flex flex-wrap justify-center py-8 mx-auto gap-7")}>
        <Avatar {...author} />
      </div>
    </div>
  );
}
