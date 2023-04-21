import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { MantineProvider } from '@mantine/core'
import { ModalsProvider } from '@mantine/modals'
import { ApolloProvider } from '@apollo/client'

import '@/index.css'
import { client } from '@/graphql/client'
import { screens } from '@/screens'
import { modals, defaultModalProps } from '@/modules/modals'

ReactDOM
  .createRoot(document.getElementById('root') as HTMLElement)
  .render(
    <React.StrictMode>
      <ApolloProvider client={client}>
        <MantineProvider withGlobalStyles withNormalizeCSS>
          <ModalsProvider modals={modals} modalProps={defaultModalProps}>
            <RouterProvider router={screens} />
          </ModalsProvider>
        </MantineProvider>
      </ApolloProvider>
    </React.StrictMode>,
  )
