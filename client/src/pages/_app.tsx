import '../styles/globals.css';
import {
  createClient,
  dedupExchange,
  cacheExchange,
  ssrExchange,
  fetchExchange,
  Provider,
} from 'urql';

import type { AppProps } from 'next/app';

const isServerSide = typeof window === 'undefined';
const ssrCache = ssrExchange({ isClient: !isServerSide });
const client = createClient({
  url: process.env.NEXT_PUBLIC_URQL_REQUEST_DEST,
  exchanges: [dedupExchange, cacheExchange, ssrCache, fetchExchange],
  fetchOptions: () => {
    return { headers: {} };
  },
});

function MyApp({ Component, pageProps }: AppProps) {
  if (pageProps.urqlState) {
    ssrCache.restoreData(pageProps.urqlState);
  }
  return (
    <Provider value={client}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;
