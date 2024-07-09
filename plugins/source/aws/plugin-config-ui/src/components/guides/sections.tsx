import { Stack, Typography } from '@mui/material';
import { CodeSnippet } from './codeSnippet';

type Section = {
  header?: string;
  bodies: {
    code?: string;
    image?: any;
    text?: string;
    Node?: any;
  }[];
};

type Props = {
  sections: Section[];
};

export function Sections({ sections }: Props) {
  return (
    <Stack gap={1.5}>
      {sections.map((section, index) => (
        <Stack key={index} gap={1}>
          <Typography variant="h6">{section.header}</Typography>
          {section.bodies.map((body, index) => {
            if (body.Node) {
              return <body.Node />;
            } else if (body.code) {
              return <CodeSnippet key={index} text={body.code} />;
            } else if (body.image) {
              return <img key={body.image} src={body.image} alt={body.text} />;
            } else {
              return (
                <Typography key={index} variant="body1">
                  {body.text}
                </Typography>
              );
            }
          })}
        </Stack>
      ))}
    </Stack>
  );
}
