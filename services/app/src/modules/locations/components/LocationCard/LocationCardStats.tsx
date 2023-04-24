import { createStyles, rem, Card, Image, Text } from '@mantine/core'

import { LocationCardStat } from './LocationCardStat'
import { Location } from '@/graphql'

const useStyles = createStyles((theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'space-after',
    borderTop: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]
    }`,
  }
}))

type LocationCardStatsProps = {
  location: Location
}

export const LocationCardStats = (props: LocationCardStatsProps) => {
  const { classes } = useStyles()

  return (
    <Card.Section className={classes.root} mt='md' pt='md' pl='lg' pr='lg'>
      <LocationCardStat title='Total Items' value={`${props?.location?.stats?.totalItems || 0}`} />
      <LocationCardStat title='Total Value' value='$0.00' />
    </Card.Section>
  )
}
