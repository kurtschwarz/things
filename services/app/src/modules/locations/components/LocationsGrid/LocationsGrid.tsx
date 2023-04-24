import { Grid, Card, Image, Text, Stack, Group, Button, ActionIcon } from '@mantine/core'
import { BiStar } from 'react-icons/bi'

import { LocationConnection, LocationEdge } from '@/graphql'

import { LocationsGridGroup } from './LocationsGridGroup'

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
    <Stack spacing='xl'>
      {rootLocationsWithChildren.map((edge, index) => (
        <LocationsGridGroup
          name={edge?.node?.name}
          divider={index > 0}
        >
          <Grid>
            <Grid.Col span={4}>
              <Card shadow='sm' padding='lg' radius='sm' withBorder>
                <Card.Section>
                  <Image
                    src='https://images.unsplash.com/photo-1588880331179-bc9b93a8cb5e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2340&q=80'
                    height={160}
                    alt='Norway'
                  />
                </Card.Section>

                <Text weight={800} size='lg' mt='lg'>{edge?.node?.name}</Text>
                <Text size='sm' italic={!edge?.node?.description}>{edge?.node?.description || 'No description provided.'}</Text>

                <Group mt='xs'>
                  <Button radius='md' style={{ flex: 1 }}>
                    View Inventory
                  </Button>
                  <ActionIcon variant='default' radius='md' size={36}>
                    <BiStar />
                  </ActionIcon>
                </Group>
              </Card>
            </Grid.Col>

            <Grid.Col span={8}>
              <Grid>
                {edge?.node?.children?.map((child) => (
                  <Grid.Col span={4}>
                    <Card shadow='sm' padding='lg' radius='sm' withBorder>
                      <Text weight={800} size='lg'>{child.name}</Text>
                      <Text size='sm' italic={!child.description}>{child.description || 'No description provided.'}</Text>
                      <Group mt='md'>
                        <Text fz='xs' opacity={0.45}>{edge.node?.name} → {child.name}</Text>
                      </Group>
                    </Card>
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
        {rootLocationsWithoutChildren?.map((edge) => (
          <Grid>
            <Grid.Col span={4}>
              <Card shadow='sm' padding='lg' radius='sm' withBorder>
                <Card.Section>
                  <Image
                    src='https://images.unsplash.com/photo-1588880331179-bc9b93a8cb5e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2340&q=80'
                    height={160}
                    alt='Norway'
                  />
                </Card.Section>

                <Text weight={800} size='lg' mt='lg'>{edge?.node?.name}</Text>
                <Text size='sm' italic={!edge?.node?.description}>{edge?.node?.description || 'No description provided.'}</Text>
              </Card>
            </Grid.Col>
          </Grid>
        ))}
      </LocationsGridGroup>
    </Stack>
  )
}

export default LocationsGrid
