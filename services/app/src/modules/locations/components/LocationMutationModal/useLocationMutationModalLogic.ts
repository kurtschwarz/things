import { useForm } from '@mantine/form'
import type { ContextModalProps } from '@mantine/modals/lib/context'

import { Location } from '@/graphql/types'
import { useLocationMutations } from '@/modules/locations'
import { GET_LOCATIONS } from '../../queries'

export const useLocationMutationModalLogic = (
  modal: ContextModalProps,
  location?: Location,
  onCompleted?: () => void
) => {
  const form = useForm({
    initialValues: {
      name: location?.name ?? '',
      description: location?.description ?? '',
    }
  })

  const {
    createLocationMutation,
    createLocationMutationLoading,
    updateLocationMutation,
    updateLocationMutationLoading
  } = useLocationMutations()

  const loading = createLocationMutationLoading || updateLocationMutationLoading

  const handleCancel = (event?: React.MouseEvent<HTMLButtonElement>) => {
    event?.preventDefault()
    modal.context.closeContextModal(modal.id, true)
  }

  const handleMutationCompleted = () => {
    onCompleted?.()

    setTimeout(() => {
      modal.context.closeContextModal(modal.id, false)
    }, 200)
  }

  const handleSave = async (values: Pick<Location, 'name' | 'description'>) => {
    if (location?.id) {
      await updateLocationMutation({
        variables: {
          id: location.id,
          input: {
            name: values.name,
            // description: values.description
          }
        },
        // refetchQueries: [
        //   { query: GET_LOCATIONS }
        // ],
        // awaitRefetchQueries: true,
        onCompleted: handleMutationCompleted
      })

      return
    }

    await createLocationMutation({
      variables: {
        input: {
          name: values.name,
          // description: values.description
        }
      },
      // refetchQueries: [
      //   { query: GET_LOCATIONS }
      // ],
      // awaitRefetchQueries: true,
      onCompleted: handleMutationCompleted
    })
  }

  return {
    form,
    loading,
    handleCancel,
    handleSave
  }
}
