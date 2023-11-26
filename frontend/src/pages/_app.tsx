import React from 'react'
import '../styles/globals.css'
import type { AppProps } from 'next/app'

import { ModalProvider } from '../contexts/modal-context'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ModalProvider>
      <Component {...pageProps} />
    </ModalProvider>
  )
}
