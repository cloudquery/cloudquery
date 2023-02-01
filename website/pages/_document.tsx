import Document, { Html, Head, Main, NextScript } from "next/document";
import type { DocumentInitialProps, DocumentContext } from "next/document";
import { HeaderOgImage } from "../components/HeaderOgImage";

class MyDocument extends Document {
  static async getInitialProps(
    ctx: DocumentContext
  ): Promise<DocumentInitialProps> {
    const initialProps = await Document.getInitialProps(ctx);

    return initialProps;
  }

  render() {
    return (
      <Html>
        <Head>
          <link rel="preconnect" href="https://fonts.googleapis.com" />
          <link
            rel="preconnect"
            href="https://fonts.gstatic.com"
            crossOrigin="anonymous"
          />
          <link
            href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;700;800&display=optional"
            rel="stylesheet"
          />
          <HeaderOgImage />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
