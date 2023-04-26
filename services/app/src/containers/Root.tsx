import { RouterProvider } from 'react-router-dom'
import { MantineProvider, MantineThemeOverride } from '@mantine/core'
import { ApolloProvider } from '@apollo/client'

import { client } from '@/graphql/client'
import { screens } from '@/screens'
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
            <RouterProvider router={screens} />
          </MantineProvider>
        )}
      </ColorSchemeProvider>
    </ApolloProvider>
  )
}
