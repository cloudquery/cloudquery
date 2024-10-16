module.exports = {
  meta: {
    type: 'suggestion',
    docs: {
      description: 'disallow named imports from @mui/material',
      category: 'Best Practices',
      recommended: false,
    },
    fixable: 'code',
    schema: [], // no options
  },
  create(context) {
    return {
      ImportDeclaration(node) {
        if (node.source.value === '@mui/material' && node.specifiers.length > 0) {
          const namedImports = node.specifiers.filter(
            (specifier) => specifier.type === 'ImportSpecifier',
          );
          if (namedImports.length > 0) {
            context.report({
              node,
              message: 'Use default import for each module from @mui/material',
              fix(fixer) {
                const fixes = namedImports.map((specifier) => {
                  const importName = specifier.local.name;
                  let newImportStatement = `import ${importName} from '@mui/material/${importName}';`;
                  switch (importName) {
                    case 'useTheme': {
                      newImportStatement = `import useTheme from '@mui/material/styles/useTheme';`;
                      break;
                    }
                    case 'useMediaQuery': {
                      newImportStatement = `import useMediaQuery from '@mui/material/useMediaQuery';`;
                      break;
                    }
                    case 'ThemeProvider': {
                      newImportStatement = `import ThemeProvider from '@mui/material/styles/ThemeProvider';`;
                      break;
                    }
                    case 'createTheme': {
                      newImportStatement = `import createTheme from '@mui/material/styles/createTheme';`;
                      break;
                    }
                    case 'createTypography': {
                      newImportStatement = `import createTypography from '@mui/material/styles/createTypography';`;
                      break;
                    }
                    case 'styled': {
                      newImportStatement = `import styled from '@mui/material/styles/styled';`;
                      break;
                    }
                    case 'alpha': {
                      newImportStatement = `import { alpha } from '@mui/material/styles';`;
                      break;
                    }
                    case 'Theme': {
                      newImportStatement = `import { Theme } from '@mui/material/styles';`;
                      break;
                    }
                    case 'Palette': {
                      newImportStatement = `import { Palette } from '@mui/material/styles/createPalette';`;
                      break;
                    }
                    case 'PaletteColor': {
                      newImportStatement = `import { PaletteColor } from '@mui/material/styles/createPalette';`;
                      break;
                    }
                    case 'PaletteOptions': {
                      newImportStatement = `import { PaletteOptions } from '@mui/material/styles/createPalette';`;
                      break;
                    }
                    case 'Components': {
                      newImportStatement = `import { Components } from '@mui/material/styles';`;
                      break;
                    }
                    case 'Shadows': {
                      newImportStatement = `import { Shadows } from '@mui/material/styles';`;
                      break;
                    }
                    case 'Breakpoint': {
                      newImportStatement = `import { Breakpoint } from '@mui/material/styles';`;
                      break;
                    }
                    case 'TypographyOptions': {
                      newImportStatement = `import { TypographyOptions } from '@mui/material/styles/createTypography';`;
                      break;
                    }
                  }
                  return fixer.insertTextBefore(node, newImportStatement + '\n');
                });

                const removeNamedImports = fixer.remove(node);

                return fixes.concat(removeNamedImports);
              },
            });
          }
        } else if (node.source.value === '@mui/icons-material' && node.specifiers.length > 0) {
          const namedImports = node.specifiers.filter(
            (specifier) => specifier.type === 'ImportSpecifier',
          );
          if (namedImports.length > 0) {
            context.report({
              node,
              message: 'Use default import for each module from @mui/icons-material',
              fix(fixer) {
                const fixes = namedImports.map((specifier) => {
                  const importName = specifier.local.name;
                  const newImportStatement = `import ${importName}Icon from '@mui/icons-material/${importName}';`;

                  return fixer.insertTextBefore(node, newImportStatement + '\n');
                });

                const removeNamedImports = fixer.remove(node);

                return fixes.concat(removeNamedImports);
              },
            });
          }
        } else if (node.source.value === '@mui/lab' && node.specifiers.length > 0) {
          const namedImports = node.specifiers.filter(
            (specifier) => specifier.type === 'ImportSpecifier',
          );
          if (namedImports.length > 0) {
            context.report({
              node,
              message: 'Use default import for each module from @mui/lab',
              fix(fixer) {
                const fixes = namedImports.map((specifier) => {
                  const importName = specifier.local.name;
                  const newImportStatement = `import ${importName} from '@mui/lab/${importName}';`;

                  return fixer.insertTextBefore(node, newImportStatement + '\n');
                });

                const removeNamedImports = fixer.remove(node);

                return fixes.concat(removeNamedImports);
              },
            });
          }
        }
      },
    };
  },
};
