import { useQuery } from '@apollo/client'
import { useEffect, useState } from 'react'

import { Location } from '@/graphql/types'
import { GET_ALL_LOCATION_PARENTS_RECURSIVE } from '../../queries'

export const useLocationSelectLogic = () => {
  const [value, setValue] = useState('')
  const [data, setData] = useState<{ value: string, label: string, description: string }[]>([])
  const { data: allLocations, loading } = useQuery(GET_ALL_LOCATION_PARENTS_RECURSIVE, {})

  const buildLocationHierarchy = (location: Location, hierarchy: string = ''): string => {
    if (location.parent) {
      return buildLocationHierarchy(location.parent, `${location.parent.name} → ${hierarchy}`)
    }

    return hierarchy
  }

  useEffect(() => {
    if (allLocations) {
      setData(
        allLocations?.locations?.edges?.map(({ node: location }) => ({
          value: location.id,
          label: location.name,
          hierarchy: buildLocationHierarchy(location, location.name)
        }))
      )
    }
  }, [
    allLocations,
    setData
  ])

  return {
    value,
    data,
    loading
  }
}
