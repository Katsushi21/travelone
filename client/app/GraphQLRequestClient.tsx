import { GraphQLClient } from 'graphql-request';

const endpoint = process.env.NEXT_PUBLIC_URQL_REQUEST_DEST;

export const graphQLClient = new GraphQLClient(endpoint, {
  headers: {
    // authorization: 'Bearer MY_TOKEN',
  },
});
