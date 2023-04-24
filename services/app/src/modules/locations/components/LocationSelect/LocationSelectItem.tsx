import { forwardRef } from 'react'
import { Text } from '@mantine/core'

type LocationSelectItemProps = React.ComponentPropsWithoutRef<'div'> & {
  label: string
  hierarchy: string
}

export const LocationSelectItem = forwardRef<HTMLDivElement, LocationSelectItemProps>(
  ({ label, hierarchy, ...others }: LocationSelectItemProps,
  ref
) => (
    <div ref={ref} {...others}>
      <Text size='sm'>{label}</Text>
      <Text size='xs' opacity={0.65}>
        {hierarchy}
      </Text>
    </div>
  )
)
