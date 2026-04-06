import eslint from "@eslint/js";
import { FlatCompat } from "@eslint/eslintrc";
import path from "node:path";
import { fileURLToPath } from "node:url";
import tseslint from "typescript-eslint";
import prettierConfig from "eslint-config-prettier";
import unicornPlugin from "eslint-plugin-unicorn";
import unusedImportsPlugin from "eslint-plugin-unused-imports";
import importPlugin from "eslint-plugin-import-x";
import promisePlugin from "eslint-plugin-promise";
import avaPlugin from "eslint-plugin-ava";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({ baseDirectory: __dirname });

export default tseslint.config(
  {
    ignores: ["dist/**", "node_modules/**"],
  },
  eslint.configs.recommended,
  ...tseslint.configs.recommended,
  {
    languageOptions: {
      parserOptions: {
        projectService: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },
  prettierConfig,
  unicornPlugin.configs["flat/recommended"],
  promisePlugin.configs["flat/recommended"],
  ...avaPlugin.configs.recommended,
  ...compat.extends("plugin:you-dont-need-lodash-underscore/all"),
  {
    plugins: {
      "unused-imports": unusedImportsPlugin,
      "import-x": importPlugin,
    },
    settings: {
      "import-x/resolver": {
        typescript: true,
        node: true,
      },
    },
    rules: {
      "unicorn/no-null": "off",
      "unused-imports/no-unused-imports": "error",
      "no-console": "error",
      "require-await": "off",
      "@typescript-eslint/require-await": "error",
      "@typescript-eslint/naming-convention": "error",
      "import-x/no-cycle": "error",
      "import-x/no-self-import": "error",
      "@typescript-eslint/consistent-type-imports": "error",
      "import-x/order": [
        "error",
        {
          "newlines-between": "always",
          alphabetize: {
            order: "asc",
            caseInsensitive: true,
          },
        },
      ],
    },
  },
  {
    files: ["src/grpc/**/*.ts"],
    rules: {
      "@typescript-eslint/naming-convention": "off",
    },
  },
  {
    files: ["**/*test.ts"],
    rules: {
      "unicorn/no-array-for-each": "off",
    },
  },
);
