import {
  createClient,
  dedupExchange,
  cacheExchange,
  fetchExchange,
  ssrExchange,
  Provider,
} from 'urql';

type Props = {
  children: React.ReactNode;
};

const isServerSide = typeof window === 'undefined';
const ssr = ssrExchange({
  isClient: !isServerSide,
});
if (!isServerSide) {
  ssr.restoreData(window.__URQL_DATA__);
}
const client = createClient({
  url: process.env.NEXT_PUBLIC_URQL_REQUEST_DEST,
  exchanges: [dedupExchange, cacheExchange, ssr, fetchExchange],
  requestPolicy: 'cache-and-network',
  suspense: true,
  fetchOptions: () => {
    return { headers: {} };
  },
});

export const UrqlQueryWrapper = ({ children }: Props) => {
  return <Provider value={client}>{children}</Provider>;
};
