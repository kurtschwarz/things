import { modals } from '@mantine/modals'

import type { Location } from '@/graphql'
import { LOCATION_MUTATION_MODAL } from '@/modules/modals'

export const openCreateLocationMutationModal = (
  options?: {
    parentID?: string,
    onCompleted?: () => void
  }
) =>
  modals.openContextModal({
    modal: LOCATION_MUTATION_MODAL,
    title: 'New Location',
    innerProps: {
      parentID: options?.parentID,
      onCompleted: options?.onCompleted
    }
  })

export const openUpdateLocationMutationModal = (
  location: Location,
  options?: {
    onCompleted?: () => void
  }
) =>
  modals.openContextModal({
    modal: LOCATION_MUTATION_MODAL,
    title: 'Update Location',
    innerProps: {
      location,
      onCompleted: options?.onCompleted
    }
  })
