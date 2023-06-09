import { useForm } from '@mantine/form'
import type { ContextModalProps } from '@mantine/modals/lib/context'

import type { Location } from '@/graphql'
import { useLocationMutations, GET_ALL_LOCATIONS } from '@/modules/locations'

export const useLocationMutationModalLogic = (
  modal: ContextModalProps,
  location?: Location,
  parentID?: string,
  onCompleted?: () => void
) => {
  const form = useForm({
    initialValues: {
      name: location?.name ?? '',
      description: location?.description ?? '',
      parentID: (parentID || location?.parentID) ?? '',
    }
  })

  const {
    createLocationMutation,
    createLocationMutationLoading,
    updateLocationMutation,
    updateLocationMutationLoading,
    deleteLocationMutation,
    deleteLocationMutationLoading
  } = useLocationMutations()

  const loading = createLocationMutationLoading || updateLocationMutationLoading || deleteLocationMutationLoading

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

  const handleSave = async (values: Pick<Location, 'name' | 'description' | 'parentID'>) => {
    if (location?.id) {
      await updateLocationMutation({
        variables: {
          id: location.id,
          input: {
            name: values.name,
            parentID: values.parentID === '' ? null : values.parentID,
            description: values.description === '' ? null : values.description
          }
        },
        refetchQueries: [
          { query: GET_ALL_LOCATIONS }
        ],
        awaitRefetchQueries: true,
        onCompleted: handleMutationCompleted
      })

      return
    }

    await createLocationMutation({
      variables: {
        input: {
          name: values.name,
          parentID: values.parentID === '' ? null : values.parentID,
          description: values.description === '' ? null : values.description
        }
      },
      refetchQueries: [
        { query: GET_ALL_LOCATIONS }
      ],
      awaitRefetchQueries: true,
      onCompleted: handleMutationCompleted
    })
  }

  const handleDelete = async () => {
    if (!location?.id) {
      return
    }

    await deleteLocationMutation({
      variables: {
        id: location.id
      },
      refetchQueries: [
        { query: GET_ALL_LOCATIONS }
      ],
      awaitRefetchQueries: true,
      onCompleted: handleMutationCompleted
    })
  }

  return {
    form,
    loading,
    handleCancel,
    handleSave,
    handleDelete
  }
}
