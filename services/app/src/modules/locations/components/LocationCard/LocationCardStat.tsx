import { rem, Box, Text } from '@mantine/core'

type LocationCardStatProps = {
  title: string
  value: string
}

export const LocationCardStat = (props: LocationCardStatProps) => {
  return (
    <Box key={props.title} style={{ width: rem(100) }}>
      <Text size='xs' color='dimmed'>
        {props.title}
      </Text>
      <Text weight={500} size='sm'>
        {props.value}
      </Text>
    </Box>
  )
}
