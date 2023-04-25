import { createStyles, rem, Card, Image, Text, Button, Group, ActionIcon } from '@mantine/core'
import { useNavigate } from 'react-router-dom'

import { Location } from '@/graphql'

import { openUpdateLocationMutationModal } from '../../../helpers'
import { LocationCardStats } from './LocationCardStats'
import { BiStar } from 'react-icons/bi'

type LocationCardProps = {
  location: Location
  compact?: boolean
}

const useStyles = createStyles((theme) => ({
  root: {
    cursor: 'pointer'
  },

  imageButtons: {
    top: rem(18),
    right: rem(18),
    position: 'absolute'
  }
}))

export const LocationCard = (props: LocationCardProps) => {
  const { classes } = useStyles()
  const navigate = useNavigate()

  const handleCardClick = (event: React.MouseEvent<HTMLDivElement>) => {
    event.preventDefault()

    navigate(`/location/${props.location.id}`)
  }

  const handleEditClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault()
    event.stopPropagation()

    openUpdateLocationMutationModal(props.location)
  }

  const handleStarClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault()
    event.stopPropagation()
  }

  return (
    <Card shadow='sm' padding='lg' radius='sm' withBorder className={classes.root} onClick={handleCardClick}>
      {!props.compact ? (
        <Card.Section style={{ position: 'relative' }}>
          <Image
            src='https://images.unsplash.com/photo-1588880331179-bc9b93a8cb5e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2340&q=80'
            height={160}
            alt='Norway'
          />

          <Group className={classes.imageButtons} spacing='xs'>
            <ActionIcon variant='default' size='lg' onClick={handleStarClick}>
              <BiStar />
            </ActionIcon>

            <Button variant='white' color='dark' onClick={handleEditClick}>Edit</Button>
          </Group>
        </Card.Section>
      ) : null}

      <Text weight={800} size='lg' mt={!props.compact ? 'lg' : undefined}>{props.location?.name}</Text>
      <Text size='sm' italic={!props.location?.description}>{props.location?.description || 'No description provided.'}</Text>

      {/* <Group mt='xs'>
        <Button radius='md' style={{ flex: 1 }}>
          View Inventory
        </Button>
      </Group> */}

      {/* <Group mt='md'>
        <Text fz='xs' opacity={0.45}>{edge.node?.name} → {child.name}</Text>
      </Group> */}

      <LocationCardStats location={props.location} />
    </Card>
  )
}
