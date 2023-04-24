import { Grid, Card, Image, Text, Stack, Group, Button, ActionIcon, Avatar } from '@mantine/core'
import { BiStar } from 'react-icons/bi'

import { Location, LocationConnection } from '@/graphql'

import { LocationsGridGroup } from './LocationsGridGroup'
import { LocationCard } from './LocationCard/LocationCard'

type LocationsGridProps = {
  locations: LocationConnection
}

export const LocationsGrid = (props: LocationsGridProps) => {
  const rootLocationsWithChildren = (props.locations.edges || [])
    .filter((edge) => edge?.node?.parentID == null && (edge?.node?.children?.length || 0) > 0)
    .sort((a, b) => (b?.node?.children?.length || 0) - (a?.node?.children?.length || 0))

  const rootLocationsWithoutChildren = (props.locations.edges || [])
    .filter((edge) => edge?.node?.parentID == null && (edge?.node?.children?.length || 0) < 1)

  return (
    <Stack spacing='xs'>
      {rootLocationsWithChildren.map((edge, index) => (
        <LocationsGridGroup
          name={edge?.node?.name}
          divider={index > 0}
        >
          <Grid>
            <Grid.Col span={4}>
              <LocationCard
                location={edge?.node as Location}
              />
            </Grid.Col>

            <Grid.Col span={8}>
              <Grid>
                {edge?.node?.children?.map((child) => (
                  <Grid.Col span={4}>
                    <LocationCard location={child} compact />
                  </Grid.Col>
                ))}
              </Grid>
            </Grid.Col>
          </Grid>
        </LocationsGridGroup>
      ))}

      <LocationsGridGroup
        name={'All Other Locations'}
        divider={rootLocationsWithChildren.length > 0}
      >
        <Grid>
          {rootLocationsWithoutChildren?.map((edge) => (
            <Grid.Col span={4}>
              <LocationCard location={edge?.node as Location} />
            </Grid.Col>
          ))}
        </Grid>
      </LocationsGridGroup>
    </Stack>
  )
}

export default LocationsGrid
