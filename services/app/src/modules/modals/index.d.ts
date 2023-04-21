import { modals } from './modals'

declare module '@mantine/modals' {
  export interface MantineModalsOverride {
    modals: typeof modals
  }
}
