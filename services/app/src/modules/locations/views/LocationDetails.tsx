import { createStyles, rem, Image, Text, Header, Grid, Group, TextInput, Textarea, Stack, Title } from '@mantine/core'
import { useLoaderData, useRevalidator } from 'react-router'
import { useClipboard } from '@mantine/hooks'

import { client } from '@/graphql'
import { ColorSchemeToggle } from '@/modules/theme'

import { GET_LOCATION } from '../queries'
import { openCreateLocationMutationModal } from '../helpers'
import { LocationBreadcrumbs } from '../components/LocationBreadcrumbs'

const useStyles = createStyles((theme) => ({
  root: {
    width: '100%',
    height: '100%'
  },

  breadcrumbs: {
    height: rem(40),
    background: `${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]}`
  },

  detailsPane: {
    background: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
    borderRight: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[2]
    }`,
    height: '100%'
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
  const clipboard = useClipboard({ timeout: 500 })

  return (
    <div className={classes.root}>
      <Grid style={{ width: '100%', height: '100%' }} m={0} p={0}>
        <Grid.Col span={3} className={classes.detailsPane} p={0}>
          <Image
            src='https://images.unsplash.com/photo-1588880331179-bc9b93a8cb5e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2340&q=80'
          />

          <Group pl='md' pr='md' mt='md' position='left' align='top'>
            <Group>
              <TextInput value={data.location.id} label='ID' style={{ width: '100%' }} readOnly onFocus={() => clipboard.copy(data.location.id)} />
              <TextInput value={data.location.name} label='Name' style={{ width: '100%' }} />
              <Textarea value={data.location.description} label='Description' style={{ width: '100%' }} placeholder='Location description'></Textarea>
            </Group>
          </Group>
        </Grid.Col>
        <Grid.Col span='auto'>
          <Stack>
            <Stack>
              <Title order={2} size='h4'>Location Statistics</Title>
              <Grid>
                <Grid.Col span='auto'>a</Grid.Col>
                <Grid.Col span='auto'>b</Grid.Col>
                <Grid.Col span='auto'>c</Grid.Col>
              </Grid>
            </Stack>

            <Stack>
              <Title order={2} size='h4'>Location Children</Title>
              <Grid>
                <Grid.Col span='auto'>a</Grid.Col>
                <Grid.Col span='auto'>b</Grid.Col>
                <Grid.Col span='auto'>c</Grid.Col>
              </Grid>
            </Stack>

            <Stack>
              <Title order={2} size='h4'>Location Assets</Title>
              <Grid>
                <Grid.Col span='auto'>a</Grid.Col>
                <Grid.Col span='auto'>b</Grid.Col>
                <Grid.Col span='auto'>c</Grid.Col>
              </Grid>
            </Stack>
          </Stack>
        </Grid.Col>
      </Grid>
    </div>
  )
}
