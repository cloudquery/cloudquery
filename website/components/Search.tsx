import React from "react";
import { useRouter } from "next/router";
import ReactSearchBox from "react-search-box";

export default (props) => {
  const router = useRouter();

  const handleClick = (key: string) => {
    router.push(key);
  };

  const data = props.data.map(({ plugin, table }) => ({
    key: `/docs/plugins/sources/${plugin}/tables/${table}`,
    value: table,
  }));
  return (
    <div
      className="font-medium dark:text-gray-900 tables-search"
      style={{ width: "75%", marginLeft: "auto", marginRight: "auto" }}
    >
      <ReactSearchBox
        placeholder="Discover supported tables... (e.g. 'aws_s3_buckets')"
        data={data}
        onSelect={({ item: { key } }) => handleClick(key)}
        onChange={() => undefined}
        fuseConfigs={{ ignoreLocation: true, useExtendedSearch: true }}
        clearOnSelect={true}
        autoFocus={true}
      />
    </div>
  );
};
