declare namespace NodeJS {
  interface ProcessEnv {
    readonly NODE_ENV: 'development' | 'production' | 'test'
    readonly NEXT_PUBLIC_URQL_REQUEST_DEST: 'string'
  }
}
