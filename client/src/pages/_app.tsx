import '../styles/globals.css';
import { createClient, Provider } from 'urql';

import type { AppProps } from 'next/app';

const client = createClient({
  url: process.env.NEXT_PUBLIC_URQL_REQUEST_DEST,
});

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider value={client}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;
