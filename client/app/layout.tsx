import { Footer } from './components/Footer';
import { Header } from './components/Header';

import '../styles/globals.css';
import UrqlQueryWrapper from './UrqlQueryWrapper';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <UrqlQueryWrapper>
        <head>
          <title>travelone</title>
        </head>
        <body>
          <div data-theme="light">
            <header>
              <Header />
            </header>
            <main>{children}</main>
          </div>
          <Footer />
        </body>
      </UrqlQueryWrapper>
    </html>
  );
}
