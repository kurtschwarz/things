import { Breadcrumbs, Anchor } from '@mantine/core'
import { Link } from 'react-router-dom'

import { Location } from '@/graphql'
import { useMemo } from 'react'

type LocationBreadcrumbsProps = {
  location: Location
}

export const LocationBreadcrumbs = (props: LocationBreadcrumbsProps) => {
  const breadcrumbs = useMemo<React.ReactNode[]>(() => {
    const buildHierarchy = (location: Location, hierarchy: Location[]): Location[] => {
      if (location.parent) {
        return buildHierarchy(location.parent, [
          location.parent,
          ...hierarchy
        ])
      }

      return hierarchy
    }

    if (!props.location) {
      return []
    }

    return [
      <Link to='/locations'>All Locations</Link>,
      ...buildHierarchy(props.location, [props.location])
        .map((location) => (
          <Link to={`/location/${location.id}`} key={location.id} color='black'>
            {location.name}
          </Link>
        ))
    ]
  }, [
    props.location
  ])

  return (
    <Breadcrumbs separator='→'>
      {breadcrumbs}
    </Breadcrumbs>
  )
}
