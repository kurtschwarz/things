import { AppShell } from '@mantine/core'
import { useOutlet } from 'react-router'

import { Sidebar } from '../Sidebar/Sidebar'

export const ProtectedLayout = () => {
  const outlet = useOutlet()

  return (
    <AppShell
      layout='alt'
      navbar={<Sidebar />}
    >
      {outlet}
    </AppShell>
  )
}
