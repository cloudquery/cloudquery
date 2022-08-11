declare module "nextra-theme-docs/tabs" {
  type TabItem = {
    label: ReactElement;
    disabled?: boolean;
  };

  export const Tabs: React.FC<{
    items: ReactNode[] | ReadonlyArray<ReactNode> | TabItem[];
    selectedIndex?: number;
    defaultIndex?: number;
    onChange?: (index: number) => void;
    children: ReactNode;
  }>;

  export const Tab: React.FC<{
    children: ReactNode;
  }>;
}

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
