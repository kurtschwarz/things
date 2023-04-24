import { createStyles, Button, Text, Header } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'

import { client } from '@/graphql/client'
import { Location } from '@/graphql/types'

import { GET_ALL_LOCATIONS } from '../queries'
import { openCreateLocationMutationModal, openUpdateLocationMutationModal } from '../helpers'
import LocationsGrid from '../components/LocationsGrid/LocationsGrid'

const useStyles = createStyles((theme) => ({
  root: {},
}))

export const locationsDashboardLoader = async () =>
  client.query({
    query: GET_ALL_LOCATIONS
  })

export const LocationsDashboard = () => {
  const { classes } = useStyles()
  const { data } = useLoaderData() as { data: { locations: { edges: { node: Location }[] } } }
  const revalidator = useRevalidator()

  return (
    <div className={classes.root}>
      <Header height={{ base: 50, md: 70 }} p='md'>
        <div style={{ display: 'flex', alignItems: 'center', height: '100%', justifyContent: 'space-between' }}>
          <Text>Locations</Text>
          <Button
            onClick={() => {
              openCreateLocationMutationModal({
                onCompleted: () => revalidator.revalidate()
              })
            }
          }>
            New Location
          </Button>
        </div>
      </Header>

      <LocationsGrid locations={data.locations?.edges} />
    </div>
  )
}
