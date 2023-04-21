import type { ModalSettings } from '@mantine/modals/lib/context'

import { LocationMutationModal } from '@/modules/locations'

export const LOCATION_MUTATION_MODAL = 'locationMutationModal'

export const modals = {
  [LOCATION_MUTATION_MODAL]: LocationMutationModal
}

export const defaultModalProps: ModalSettings = {
  centered: true
}
