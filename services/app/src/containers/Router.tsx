import { Title } from '@mantine/core'
import { createBrowserRouter, RouteObject } from 'react-router-dom'

import { locationsDashboardLoader, LocationsDashboard, locationDetailLoader, LocationDetail } from '@/modules/locations'
import { ProtectedLayout } from '@/components/Layout/ProtectedLayout'

import { routerRootLoader, RouterRoot } from './RouterRoot'

const publicViews: RouteObject[] = [
  {
    path: '/login',
    loader: async () => ({}),
    element: <Title order={2}>Login</Title>
  }
]

const protectedViews: RouteObject[] = [
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
    path: '/location/:id',
    loader: locationDetailLoader,
    element: <LocationDetail />
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

export const router = createBrowserRouter([
  {
    path: '/',
    loader: routerRootLoader,
    element: <RouterRoot />,
    children: [
      ...publicViews,
      {
        element: <ProtectedLayout />,
        children: [
          ...protectedViews
        ]
      }
    ]
  }
])
