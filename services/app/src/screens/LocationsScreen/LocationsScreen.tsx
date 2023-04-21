import { Text } from '@mantine/core'
import { useLoaderData } from 'react-router'

import { client } from '@/graphql/client'
import { GET_LOCATIONS } from '@/graphql/queries/locations'

export const locationsLoader = async () => {
  return client.query({
    query: GET_LOCATIONS
  })
}

export const LocationsScreen = () => {
  const { data } = useLoaderData()

  const locationsList = data.locations?.edges?.map(({ node: location }) => (
    <li>{location.name}</li>
  ))

  return (
    <div>
      <Text>Locations Screen</Text>

      <ul>
        {locationsList}
      </ul>
    </div>
  )
}

export default LocationsScreen
