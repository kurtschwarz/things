import { createBrowserRouter } from 'react-router-dom'

import RootScreen, { rootLoader } from './RootScreen/RootScreen'
import DashboardScreen, { dashboardLoader } from './DashboardScreen/DashboardScreen'
import LoginScreen, { loginLoader } from './LoginScreen/LoginScreen'
import { ProtectedLayout } from '../components/Layout/ProtectedLayout'
import LocationsScreen, { locationsLoader } from './LocationsScreen/LocationsScreen'

export const screens = createBrowserRouter([
  {
    path: '/',
    loader: rootLoader,
    element: <RootScreen />,
    children: [
      {
        path: '/login',
        loader: loginLoader,
        element: <LoginScreen />
      },
      {
        element: <ProtectedLayout />,
        children: [
          {
            path: '/',
            loader: dashboardLoader,
            element: <DashboardScreen />
          },
          {
            path: '/locations',
            loader: locationsLoader,
            element: <LocationsScreen />
          },
          {
            path: '/tags',
            loader: async () => ({}),
            element: <div>Tags Screen</div>
          },
          {
            path: '/settings',
            loader: async () => ({}),
            element: <div>Settings Screen</div>
          }
        ]
      }
    ]
  }
])
