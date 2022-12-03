import { GraphQLClient } from 'graphql-request';

import { getSdk } from './generated/graphql';

const endpoint = process.env.NEXT_PUBLIC_URQL_REQUEST_DEST;

const client = new GraphQLClient(endpoint, {
  headers: {
    // authorization: 'Bearer MY_TOKEN',
  },
});

export const sdk = getSdk(client);
