import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: 'http://host.docker.internal:8000/graphql',
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
        extension: '.generated.ts',
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
