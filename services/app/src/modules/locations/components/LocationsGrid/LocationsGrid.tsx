import { Grid, Card, Image, Text, Stack, Group, Button, ActionIcon } from '@mantine/core'
import { BiStar } from 'react-icons/bi'

import { Location } from '@/graphql/types'

import { LocationsGridGroup } from './LocationsGridGroup'

type LocationsGridProps = {
  locations: { node: Location }[]
}

export const LocationsGrid = (props: LocationsGridProps) => {
  const rootLocationsWithChildren = props.locations
    .filter(({ node: location }) => location.parentID == null && location.children?.length > 0)
    .sort((a, b) => b.node.children?.length - a.node.children?.length) as LocationsGridProps['locations']

  const rootLocationsWithoutChildren = props.locations
    .filter(({ node: location }) => location.parentID == null && location.children?.length < 1) as LocationsGridProps['locations']

  return (
    <Stack spacing='xl'>
      {rootLocationsWithChildren.map(({ node: rootLocation }, index) => (
        <LocationsGridGroup
          name={rootLocation.name}
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

                <Text weight={800} size='lg' mt='lg'>{rootLocation.name}</Text>
                <Text size='sm' italic={!rootLocation.description}>{rootLocation.description || 'No description provided.'}</Text>

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
                {rootLocation.children.map((childLocation) => (
                  <Grid.Col span={4}>
                    <Card shadow='sm' padding='lg' radius='sm' withBorder>
                      <Text weight={800} size='lg'>{childLocation.name}</Text>
                      <Text size='sm' italic={!childLocation.description}>{childLocation.description || 'No description provided.'}</Text>
                      <Group mt='md'>
                        <Text fz='xs' opacity={0.45}>{rootLocation.name} → {childLocation.name}</Text>
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
        {rootLocationsWithoutChildren.map(({ node: location }) => (
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

                <Text weight={800} size='lg' mt='lg'>{location.name}</Text>
                <Text size='sm' italic={!location.description}>{location.description || 'No description provided.'}</Text>
              </Card>
            </Grid.Col>
          </Grid>
        ))}
      </LocationsGridGroup>
    </Stack>
  )
}

export default LocationsGrid
