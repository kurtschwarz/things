import { AppShell } from '@mantine/core'
import { useOutlet } from 'react-router'

import { GlobalSidebar } from '@/containers/GlobalSidebar'

export const ProtectedLayout = () => {
  const outlet = useOutlet()

  return (
    <AppShell
      layout='alt'
      navbar={<GlobalSidebar />}
      styles={(theme) => ({
        main: {
          backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] :  theme.colors.gray[0] 
        }
      })}
    >
      {outlet}
    </AppShell>
  )
}
