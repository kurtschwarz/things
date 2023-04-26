import { useMutation } from '@apollo/client'

import { useGetAllLocationsQuery, useGetAllLocationsWithParentsQuery } from '@/graphql'

import { CREATE_LOCATION, UPDATE_LOCATION, DELETE_LOCATION } from './mutations'

export const useLocationMutations = () => {
  const [createLocationMutation, { loading: createLocationMutationLoading }] = useMutation(CREATE_LOCATION)
  const [updateLocationMutation, { loading: updateLocationMutationLoading }] = useMutation(UPDATE_LOCATION)
  const [deleteLocationMutation, { loading: deleteLocationMutationLoading }] = useMutation(DELETE_LOCATION)

  return {
    createLocationMutation,
    createLocationMutationLoading,
    updateLocationMutation,
    updateLocationMutationLoading,
    deleteLocationMutation,
    deleteLocationMutationLoading
  }
}

export {
  useGetAllLocationsQuery,
  useGetAllLocationsWithParentsQuery
}
