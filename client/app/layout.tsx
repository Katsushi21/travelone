import { Footer } from './components/Footer';
import { Header } from './components/Header';
import '../styles/globals.css';

const RootLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <html lang="ja">
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
        <footer>
          <Footer />
        </footer>
      </body>
    </html>
  );
};

export default RootLayout;
