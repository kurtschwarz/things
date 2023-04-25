import { createStyles, rem, Text, Button, Header, Group } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'

import { client } from '@/graphql'
import { ColorSchemeToggle } from '@/modules/theme'

import { GET_LOCATION } from '../queries'
import { openCreateLocationMutationModal } from '../helpers'
import { LocationBreadcrumbs } from '../components/LocationBreadcrumbs'

const useStyles = createStyles((theme) => ({
  root: {},
  breadcrumbs: {
    height: rem(40),
    background: `${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]}`
  }
}))

export const locationDetailLoader = async ({ params }) => {
  return client.query({
    query: GET_LOCATION,
    variables: { id: params.id },
    fetchPolicy: 'no-cache'
  })
}

export const LocationDetail = () => {
  const { classes } = useStyles()
  const { data } = useLoaderData() as Awaited<ReturnType<typeof locationDetailLoader>>
  const revalidator = useRevalidator()

  return (
    <div className={classes.root}>
      <Header height={{ base: 101 }}>
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

        <Group className={classes.breadcrumbs} pl='md' pr='md'>
          <LocationBreadcrumbs location={data.location} />
        </Group>
      </Header>
    </div>
  )
}
