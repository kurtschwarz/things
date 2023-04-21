import { useMutation } from '@apollo/client'

import { CREATE_LOCATION, UPDATE_LOCATION } from './mutations'
import { GET_LOCATIONS } from './queries'

export const useLocationMutations = () => {
  const [createLocationMutation, { loading: createLocationMutationLoading }] = useMutation(CREATE_LOCATION)
  const [updateLocationMutation, { loading: updateLocationMutationLoading }] = useMutation(UPDATE_LOCATION)

  return {
    createLocationMutation,
    createLocationMutationLoading,
    updateLocationMutation,
    updateLocationMutationLoading
  }
}
