import { createStyles, Text, Button, Header } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'

import { client } from '@/graphql'

import { GET_LOCATION } from '../queries'
import { openCreateLocationMutationModal } from '../helpers'

const useStyles = createStyles((theme) => ({
  root: {},
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
      <Header height={{ base: 50, md: 70 }} p='md'>
        <div style={{ display: 'flex', alignItems: 'center', height: '100%', justifyContent: 'space-between' }}>
          <Text>{data.locations.edges[0].node.name}</Text>
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
    </div>
  )
}
