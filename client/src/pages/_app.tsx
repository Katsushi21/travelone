import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { createClient, Provider } from 'urql'

console.log(process.env.NEXT_PUBLIC_URQL_REQUEST_DEST)
const client = createClient({
  url: process.env.NEXT_PUBLIC_URQL_REQUEST_DEST,
})

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider value={client}>
      <Component {...pageProps} />
    </Provider>
  )
}

export default MyApp
