import { executeExchange } from '@urql/exchange-execute';
import { withUrqlClient, initUrqlClient } from 'next-urql';
import { ssrExchange, dedupExchange, cacheExchange, useQuery } from 'urql';

import { PostsDocument } from '../queries/query.generated';

import { schema } from '@/server/graphql'; // our GraphQL server's executable schema

const url = process.env.NEXT_PUBLIC_GRAPHQL_REQUEST_DEST;

const POSTS_QUERY = `
  query { todos { id text } }
`;

function Posts() {
  const [res] = useQuery({ query: POSTS_QUERY });
  return (
    <div>
      {res.data.todos.map((todo) => (
        <div key={todo.id}>
          {todo.id} - {todo.text}
        </div>
      ))}
    </div>
  );
}

export async function getServerSideProps(ctx) {
  const ssrCache = ssrExchange({ isClient: false });
  const client = initUrqlClient(
    {
      url: url, // not needed without `fetchExchange`
      exchanges: [
        dedupExchange,
        cacheExchange,
        ssrCache,
        executeExchange({ schema }), // replaces `fetchExchange`
      ],
    },
    false,
  );

  await client.query(PostsDocument).toPromise();

  return {
    props: {
      urqlState: ssrCache.extractData(),
    },
  };
}

export default withUrqlClient((ssr) => ({
  url: url,
}))(Posts);
