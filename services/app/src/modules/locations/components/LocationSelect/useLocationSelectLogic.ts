import { useQuery } from '@apollo/client'
import { useEffect, useState } from 'react'

import type { Location } from '@/graphql'

import { GET_ALL_LOCATION_PARENTS_RECURSIVE } from '../../queries'

type LocationSelectOption = { value: string, label: string, hierarchy: string }

export const useLocationSelectLogic = () => {
  const [options, setOptions] = useState<LocationSelectOption[]>([])
  const { data, loading } = useQuery(GET_ALL_LOCATION_PARENTS_RECURSIVE, {})

  const buildLocationHierarchy = (location: Location | null | undefined, hierarchy: string = ''): string => {
    if (location?.parent) {
      return buildLocationHierarchy(location.parent, `${location.parent.name} → ${hierarchy}`)
    }

    return hierarchy
  }

  useEffect(() => {
    if (data && data.locations) {
      setOptions(
        (data.locations.edges ?? [])
          .map<LocationSelectOption>((edge) => ({
            value: edge!.node!.id,
            label: edge!.node!.name ?? edge!.node!.id,
            hierarchy: buildLocationHierarchy(edge!.node, edge!.node?.name ?? '')
          }))
      )
    }
  }, [
    data,
    setOptions
  ])

  return {
    options,
    loading
  }
}
