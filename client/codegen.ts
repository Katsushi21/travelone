import type { CodegenConfig } from '@graphql-codegen/cli';

const schema = process.env.NEXT_PUBLIC_GRAPHQL_REQUEST_DEST;

const config: CodegenConfig = {
  overwrite: true,
  schema: schema,
  documents: '**/*.graphql',
  generates: {
    './types/graphql.ts': {
      plugins: [
        {
          add: {
            content: '/* eslint-disable */',
          },
        },
        'typescript',
      ],
      config: {
        immutableTypes: true,
        defaultScalarType: 'unknown',
        enumsAsConst: true,
        skipTypename: true,
        avoidOptionals: true,
      },
    },
    './types': {
      preset: 'near-operation-file',
      presetConfig: {
        baseTypesPath: 'graphql.ts',
      },
      plugins: [
        {
          add: {
            content: '/* eslint-disable */',
          },
        },
        'typescript-operations',
        'typed-document-node',
      ],
      config: {
        immutableTypes: true,
        defaultScalarType: 'unknown',
        enumsAsConst: true,
        skipTypename: true,
        avoidOptionals: true,
      },
    },
  },
  hooks: {
    afterOneFileWrite: ['prettier --write'],
  },
};
export default config;
