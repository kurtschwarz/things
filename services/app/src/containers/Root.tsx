import { useState } from 'react'
import { RouterProvider } from 'react-router-dom'
import { MantineProvider, MantineThemeOverride } from '@mantine/core'
import { ModalsProvider } from '@mantine/modals'
import { ApolloProvider } from '@apollo/client'

import { client } from '@/graphql/client'
import { screens } from '@/screens'
import { modals, defaultModalProps } from '@/modules/modals'
import { ColorSchemeProvider } from '@/modules/theme'

const theme: MantineThemeOverride = {
  primaryColor: 'teal',
  primaryShade: 5,
  fontFamily: '"Inter"'
}

export const Root = () => {
  return (
    <ApolloProvider client={client}>
      <ColorSchemeProvider>
        {(colorScheme) => (
          <MantineProvider theme={{ ...theme, colorScheme }} withGlobalStyles withNormalizeCSS>
            <ModalsProvider modals={modals} modalProps={defaultModalProps}>
              <RouterProvider router={screens} />
            </ModalsProvider>
          </MantineProvider>
        )}
      </ColorSchemeProvider>
    </ApolloProvider>
  )
}
