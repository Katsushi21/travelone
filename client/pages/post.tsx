import { withUrqlClient, initUrqlClient } from 'next-urql';
import { ssrExchange, dedupExchange, cacheExchange, fetchExchange } from 'urql';

import { PostList } from 'components/layouts/PostList';

import { PostsDocument } from '../queries/query.generated';

const url = process.env.NEXT_PUBLIC_GRAPHQL_REQUEST_DEST;

export async function getServerSideProps() {
  const ssrCache = ssrExchange({ isClient: false });
  const client = initUrqlClient(
    {
      url: url,
      exchanges: [dedupExchange, cacheExchange, ssrCache, fetchExchange],
    },
    false,
  );

  if (!client) {
    return {
      props: {},
    };
  }

  await client.query(PostsDocument, {}).toPromise();

  return {
    props: {
      urqlState: ssrCache.extractData(),
    },
  };
}

function Posts(props) {
  return <PostList props={props} />;
}

export default withUrqlClient((ssr) => ({
  url: url,
}))(Posts);
