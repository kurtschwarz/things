import { useMutation } from '@apollo/client'

import { useGetAllLocationsQuery, useGetAllLocationsWithParentsQuery } from '@/graphql'
import { CREATE_LOCATION, UPDATE_LOCATION } from './mutations'

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

export {
  useGetAllLocationsQuery,
  useGetAllLocationsWithParentsQuery
}
