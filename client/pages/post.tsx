import { withUrqlClient, initUrqlClient } from 'next-urql';
import {
  ssrExchange,
  dedupExchange,
  cacheExchange,
  fetchExchange,
  useQuery,
} from 'urql';
import { executeExchange } from '@urql/exchange-execute';

import { schema } from '@/server/graphql'; // our GraphQL server's executable schema

const TODOS_QUERY = `
  query { todos { id text } }
`;

function Todos() {
  const [res] = useQuery({ query: TODOS_QUERY });
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
      url: '', // not needed without `fetchExchange`
      exchanges: [
        dedupExchange,
        cacheExchange,
        ssrCache,
        executeExchange({ schema }), // replaces `fetchExchange`
      ],
    },
    false,
  );

  await client.query(TODOS_QUERY).toPromise();

  return {
    props: {
      urqlState: ssrCache.extractData(),
    },
  };
}

export default withUrqlClient((ssr) => ({
  url: 'your-url',
}))(Todos);
