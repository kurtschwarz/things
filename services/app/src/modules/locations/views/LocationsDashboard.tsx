import { Button, Title } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'

import { client } from '@/graphql/client'
import { Location } from '@/graphql/types'

import { GET_LOCATIONS } from '../queries'
import { openCreateLocationMutationModal, openUpdateLocationMutationModal } from '../helpers'


export const locationsDashboardLoader = async () =>
  client.query({
    query: GET_LOCATIONS
  })

export const LocationsDashboard = () => {
  const { data } = useLoaderData() as { data: { locations: { edges: { node: Location }[] } } }
  const revalidator = useRevalidator()

  const locationsList = data.locations?.edges?.map(({ node: location }) => {
    const handleEditClick = (event: React.MouseEvent<HTMLAnchorElement>) => {
      event.preventDefault()

      openUpdateLocationMutationModal(location, {
        onCompleted: () => revalidator.revalidate()
      })
    }

    return (
      <li key={location.id}>{location.name} <a href='#' onClick={handleEditClick}>Edit</a></li>
    )
  })

  return (
    <div>
      <Title order={2}>Locations</Title>

      <Button
        onClick={() => {
          openCreateLocationMutationModal({
            onCompleted: () => revalidator.revalidate()
          })
        }
      }>
        New Location
      </Button>

      <ul>
        {locationsList}
      </ul>
    </div>
  )
}
