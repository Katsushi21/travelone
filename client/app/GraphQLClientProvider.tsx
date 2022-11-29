import { GraphQLClient } from 'graphql-request';

export const useGraphQLClient = () => {
  const endpoint = process.env.NEXT_PUBLIC_URQL_REQUEST_DEST;
  const graphQLClient = new GraphQLClient(endpoint, {
    headers: {
      // authorization: 'Bearer MY_TOKEN',
    },
  });
  return graphQLClient;
};
