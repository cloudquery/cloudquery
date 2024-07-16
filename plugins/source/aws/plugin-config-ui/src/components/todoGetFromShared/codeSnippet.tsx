import Box from '@mui/material/Box';
import { useTheme } from '@mui/material';
import { useMemo } from 'react';

// NOTE: idea is to keep this lightweight and not need to import a full library. Maybe worth putting something in cloud-ui..
// https://dev.to/gauravadhikari1997/show-json-as-pretty-print-with-syntax-highlighting-3jpm
function syntaxHighlight(json: any) {
  if (!json) return ''; //no JSON from response

  json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  return json.replace(
    /("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g,
    function (match: any) {
      var cls = 'number';
      if (/^"/.test(match)) {
        if (/:$/.test(match)) {
          cls = 'key';
        } else {
          cls = 'value';
        }
      } else {
        cls = 'value';
      }
      return '<span class="' + cls + '">' + match + '</span>';
    },
  );
}

interface Props {
  text: string;
}

export function CodeSnippet({ text }: Props) {
  const { palette } = useTheme();

  const html = useMemo(() => {
    return syntaxHighlight(JSON.stringify(JSON.parse(text), undefined, 4));
  }, [text]);

  return (
    <Box
      sx={{
        '& pre': { whiteSpace: 'break-spaces', outline: 'none', margin: 1.5, fontSize: '12px' },
        '& .value': { color: palette.text.secondary },
        '& .key': { color: palette.info.main },
      }}
    >
      <pre
        dangerouslySetInnerHTML={{
          __html: html,
        }}
      />
    </Box>
  );
}
