import { Title } from '@mantine/core'
import { createBrowserRouter } from 'react-router-dom'

import { locationsDashboardLoader, LocationsDashboard } from '@/modules/locations'

import RootScreen, { rootLoader } from './RootScreen/RootScreen'
import { ProtectedLayout } from '../components/Layout/ProtectedLayout'

export const screens = createBrowserRouter([
  {
    path: '/',
    loader: rootLoader,
    element: <RootScreen />,
    children: [
      {
        path: '/login',
        loader: async () => ({}),
        element: <Title order={2}>Login</Title>
      },
      {
        element: <ProtectedLayout />,
        children: [
          {
            path: '/',
            loader: async () => ({}),
            element: <Title order={2}>Dashboard</Title>
          },
          {
            path: '/locations',
            loader: locationsDashboardLoader,
            element: <LocationsDashboard />
          },
          {
            path: '/tags',
            loader: async () => ({}),
            element: <Title order={2}>Tags</Title>
          },
          {
            path: '/settings',
            loader: async () => ({}),
            element: <Title order={2}>Settings</Title>
          }
        ]
      }
    ]
  }
])
