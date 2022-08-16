import { ComponentProps } from "react";
import {
  ArrowsExpandIcon,
  BeakerIcon,
  ChipIcon,
  DatabaseIcon,
  LightningBoltIcon,
  CodeIcon,
} from "@heroicons/react/outline";

export type Feature = {
  name: string;
  description: string;
  Icon: (props: ComponentProps<"svg">) => JSX.Element;
  page: "all" | "home" | "docs";
};

export type Features = Array<Feature>;

const FEATURES: Features = [
  {
    name: "Open source",
    description: `Easily extensible plugin architecture. Contribute to our official plugins or develop your own with CloudQuery SDK.`,
    Icon: CodeIcon,
    page: "all",
  },
  {
    name: "Blazing fast",
    description: `CloudQuery concurrency system utilizes the excelent Go concurrency model with light-weight goroutines.`,
    Icon: LightningBoltIcon,
    page: "all",
  },
  {
    name: "Database agnostic",
    description: `CloudQuery can store your configuration in any supported destination such as database, datalake, streaming for further analysis.`,
    Icon: DatabaseIcon,
    page: "all",
  },
  {
    name: "Raw access to data",
    description: `Decouple data ingestion and having raw access to your data you can built your own security stack and re-use best-of-breed tools for querying (SQL, ...), transformation (dbt,...) and visualization (Grafana, Preset, Metabase, PowerBI, ...).`,
    Icon: ArrowsExpandIcon,
    page: "all",
  },
  {
    name: "Bre-built queries",
    description: `CloudQuery maintains rich set of SQL queries and Grafana dashboards for asset inventory, CSPMs, Security & Compliance, Cost, use-cases.`,
    Icon: BeakerIcon,
    page: "all",
  },
];

export const DOCS_FEATURES = FEATURES.filter(
  (f) => f.page === "docs" || f.page === "all"
);

export const HOME_FEATURES = FEATURES.filter(
  (f) => f.page === "home" || f.page === "all"
);

export default FEATURES;
