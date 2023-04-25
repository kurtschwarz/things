import { createStyles, Button, Text, Header, Group } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'

import { client } from '@/graphql'
import { ColorSchemeToggle } from '@/modules/theme'

import { GET_ALL_LOCATIONS } from '../queries'
import { openCreateLocationMutationModal } from '../helpers'
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
  const { data } = useLoaderData() as Awaited<ReturnType<typeof locationsDashboardLoader>>
  const revalidator = useRevalidator()

  return (
    <div className={classes.root}>
      <Header height={{ base: 61 }}>
        <Group style={{ height: 61 }} pl='md' pr='md'>
          <Text weight={600}>Locations</Text>

          <Button
            style={{ marginLeft: 'auto' }}
            onClick={() => {
              openCreateLocationMutationModal({
                onCompleted: () => revalidator.revalidate()
              })
            }
          }>
            New Location
          </Button>

          <ColorSchemeToggle />
        </Group>
      </Header>

      <LocationsGrid locations={data.locations} />
    </div>
  )
}
