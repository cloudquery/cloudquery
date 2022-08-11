declare module "nextra-theme-docs/tabs";

declare module "nextra/context" {
  export interface Page {
    children?: Array<Page>;
    meta?: { title: string; theme: Record<string, string> };
    frontMatter?: Record<string, any>;
    name: string;
    route: string;
  }

  declare function getAllPages(): Array<Page>;

  declare function getPagesUnderRoute(route: string): Array<Page>;
}
