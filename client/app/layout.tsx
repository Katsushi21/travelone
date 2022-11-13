import { Footer } from './components/Footer';
import { Header } from './components/Header';
import '../styles/globals.css';
import UrqlQueryWrapper from './UrqlQueryWrapper';

const RootLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <html lang="ja">
      <head>
        <title>travelone</title>
      </head>
      <body>
        <UrqlQueryWrapper>
          <div data-theme="light">
            <header>
              <Header />
            </header>
            <main>{children}</main>
          </div>
          <Footer />
        </UrqlQueryWrapper>
      </body>
    </html>
  );
};

export default RootLayout;
