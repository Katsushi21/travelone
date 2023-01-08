import { executeExchange } from '@urql/exchange-execute';
import { buildASTSchema } from 'graphql';
import { GetServerSideProps } from 'next';
import { withUrqlClient, initUrqlClient } from 'next-urql';
import { ssrExchange, dedupExchange, cacheExchange, useQuery } from 'urql';

import { PostsDocument } from '../queries/query.generated';

const url = process.env.NEXT_PUBLIC_GRAPHQL_REQUEST_DEST;

function Posts({ props }) {
  const [res] = useQuery({ query: PostsDocument });
  return (
    <div>
      {res.data?.posts.map((post) => (
        <div key={post.id}>
          {post.id} - {post.body}
        </div>
      ))}
    </div>
  );
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const ssrCache = ssrExchange({ isClient: false });
  const schema = buildASTSchema(PostsDocument);
  const client = initUrqlClient(
    {
      url: url,
      exchanges: [
        dedupExchange,
        cacheExchange,
        ssrCache,
        executeExchange({ schema }),
      ],
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
};

export default withUrqlClient((ssrCache) => ({
  url: url,
}))(Posts);
