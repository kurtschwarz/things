import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { MantineProvider } from '@mantine/core'
import { ApolloProvider } from '@apollo/client'

import '@/index.css'
import { client } from '@/graphql/client'
import { screens } from '@/screens'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <ApolloProvider client={client}>
        <RouterProvider router={screens} />
      </ApolloProvider>
    </MantineProvider>
  </React.StrictMode>,
)
