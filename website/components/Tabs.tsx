import * as NextraTabs from "nextra-theme-docs/tabs";
import React, { useState } from "react";

export const Tab = NextraTabs.Tab;

interface Props {
  options: Array<string>;
  children: React.ReactNode;
}

export function Tabs({ options, children }: Props) {
  const [index, changeIndex] = useState(0);

  const items = options.map((value) => ({ label: value }));
  return (
    <NextraTabs.Tabs
      onChange={(index) => changeIndex(index)}
      selectedIndex={index}
      items={items}
    >
      <div className="mb-4" />
      {children}
    </NextraTabs.Tabs>
  );
}
