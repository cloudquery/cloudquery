{
  "root": true,
  "plugins": ["import", "custom-rules"],
  "extends": [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:jsx-a11y/recommended",
    "plugin:prettier/recommended",
    "plugin:unicorn/recommended",
    "plugin:import/typescript"
  ],
  "globals": {
    "JSX": true
  },
  "env": {
    "jest": true,
    "browser": true,
    "node": true
  },
  "rules": {
    "custom-rules/mui-material-imports": "error",
    "@typescript-eslint/explicit-module-boundary-types": 0,
    "@typescript-eslint/no-empty-interface": "error",
    "@typescript-eslint/no-explicit-any": 0,
    "@typescript-eslint/no-unused-vars": [
      "error",
      {
        "args": "all",
        "argsIgnorePattern": "^_",
        "vars": "all",
        "varsIgnorePattern": "^_"
      }
    ],
    "jsx-a11y/alt-text": "off",
    "jsx-a11y/anchor-is-valid": [
      "error",
      {
        "components": ["Link"],
        "specialLink": ["to"]
      }
    ],
    "jsx-a11y/label-has-associated-control": [
      "error",
      {
        "depth": 5
      }
    ],
    "import/newline-after-import": "error",
    "import/no-duplicates": ["error", { "prefer-inline": false }],
    "import/no-unresolved": "error",
    "import/order": [
      "error",
      {
        "alphabetize": {
          "caseInsensitive": true,
          "order": "asc"
        },
        "groups": [
          "builtin",
          "external",
          "internal",
          ["sibling", "parent"],
          "index",
          "object",
          "type",
          "unknown"
        ],
        "newlines-between": "always-and-inside-groups",
        "pathGroups": [
          {
            "group": "external",
            "pattern": "react",
            "position": "before"
          }
        ],
        "pathGroupsExcludedImportTypes": ["react"]
      }
    ],
    "jsx-a11y/media-has-caption": "off",
    "jsx-a11y/no-autofocus": "off",
    "jsx-a11y/no-onchange": "off",
    "newline-before-return": "error",
    "no-console": "warn",
    "no-const-assign": "error",
    "no-debugger": "error",
    "no-extra-semi": "off",
    "no-param-reassign": "error",
    "no-prototype-builtins": 0,
    "no-undef": "error",
    "no-unexpected-multiline": "error",
    "no-unused-vars": "error",
    "object-curly-newline": [
      "error",
      {
        "consistent": true
      }
    ],
    "object-shorthand": "error",
    "prefer-destructuring": [
      "error",
      {
        "AssignmentExpression": {
          "array": false,
          "object": false
        },
        "VariableDeclarator": {
          "array": false,
          "object": true
        }
      }
    ],
    "prefer-template": "error",
    "react/destructuring-assignment": [
      "error",
      "always",
      {
        "destructureInSignature": "always"
      }
    ],
    "react/display-name": "off",
    "react/jsx-boolean-value": ["error", "always"],
    "react/jsx-first-prop-new-line": ["error", "multiline"],
    "react/jsx-fragments": ["error", "syntax"],
    "react/jsx-uses-react": "error",
    "react/jsx-uses-vars": "error",
    "react/react-in-jsx-scope": "off",
    "unicorn/filename-case": "off",
    "unicorn/no-nested-ternary": "off",
    "unicorn/no-null": "off",
    "unicorn/no-useless-undefined": "off",
    "unicorn/prevent-abbreviations": "off"
  },
  "ignorePatterns": ["node_modules/", "dist/", "*.js", "*.cjs"],
  "settings": {
    "react": {
      "version": "detect"
    }
  }
}
