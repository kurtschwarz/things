import { AppShell, createStyles } from '@mantine/core'
import { ModalsProvider } from '@mantine/modals'
import { useLocation, useNavigation, useOutlet } from 'react-router'
import { BiSearch } from 'react-icons/bi'

import { modals, defaultModalProps } from '@/modules/modals'
import { SpotlightProvider } from '@/modules/spotlight'
import { GlobalHeader } from '@/containers/GlobalHeader'
import { GlobalSidebar } from '@/containers/GlobalSidebar'

const useStyles = createStyles((theme) => ({
  shell: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] :  theme.colors.gray[0]
  }
}))

export const ProtectedLayout = () => {
  const { classes } = useStyles()
  const location = useLocation()
  const outlet = useOutlet()

  return (
    <ModalsProvider
      modals={modals}
      modalProps={defaultModalProps}
    >
      <SpotlightProvider>
        <AppShell
          layout='alt'
          header={<GlobalHeader />}
          navbar={<GlobalSidebar />}
          className={classes.shell}
          padding={location.pathname.startsWith('/location/') ? 0 : 'md'}
        >
          {outlet}
        </AppShell>
      </SpotlightProvider>
    </ModalsProvider>
  )
}
