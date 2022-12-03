import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { Footer } from './components/Footer';
import { Header } from './components/Header';
import '../styles/globals.css';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 1000,
    },
  },
});

const RootLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <html lang="ja">
      <head>
        <title>travelone</title>
      </head>
      <body>
        {/* <QueryClientProvider client={queryClient}> */}
        <div data-theme="light">
          <header>
            <Header />
          </header>
          <main>{children}</main>
        </div>
        <footer>
          <Footer />
        </footer>
        {/* </QueryClientProvider> */}
      </body>
    </html>
  );
};

export default RootLayout;
