import { RouterProvider } from 'react-router-dom'
import { MantineProvider, MantineThemeOverride } from '@mantine/core'
import { ModalsProvider } from '@mantine/modals'
import { ApolloProvider } from '@apollo/client'

import { client } from '@/graphql/client'
import { screens } from '@/screens'
import { modals, defaultModalProps } from '@/modules/modals'

const theme: MantineThemeOverride = {
  primaryColor: 'teal',
  primaryShade: 5,
  fontFamily: '"Inter"'
}

export const Root = () => {
  return (
    <ApolloProvider client={client}>
      <MantineProvider theme={theme} withGlobalStyles withNormalizeCSS>
        <ModalsProvider modals={modals} modalProps={defaultModalProps}>
          <RouterProvider router={screens} />
        </ModalsProvider>
      </MantineProvider>
    </ApolloProvider>
  )
}
