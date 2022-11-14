import { Tab as NextraTab, Tabs as NextraTabs } from "nextra-theme-docs";
import React from "react";

export const Tab = NextraTab;

interface Props {
  options: Array<string>;
  children: React.ReactNode;
  defaultIndex: number;
}

export function Tabs({ options, children, defaultIndex }: Props) {
  const items = options.map((value) => ({ label: value }));
  return (
    <NextraTabs items={items} defaultIndex={defaultIndex}>
      <div className="mb-4" />
      {children}
    </NextraTabs>
  );
}
