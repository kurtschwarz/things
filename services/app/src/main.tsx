import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { MantineProvider } from '@mantine/core'

import './index.css'
import { screens } from './screens'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <RouterProvider router={screens} />
    </MantineProvider>
  </React.StrictMode>,
)
