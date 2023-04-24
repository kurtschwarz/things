import { createStyles, rem, Card, Image, Text } from '@mantine/core'

import { Location } from '@/graphql'
import { LocationCardStats } from './LocationCardStats'

type LocationCardProps = {
  location: Location
  compact?: boolean
}

export const LocationCard = (props: LocationCardProps) => {
  return (
    <Card shadow='sm' padding='lg' radius='sm' withBorder>
      {!props.compact ? (
        <Card.Section>
          <Image
            src='https://images.unsplash.com/photo-1588880331179-bc9b93a8cb5e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2340&q=80'
            height={160}
            alt='Norway'
          />
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
