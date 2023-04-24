import { Select, SelectProps } from '@mantine/core'

import { useLocationSelectLogic } from './useLocationSelectLogic'
import { LocationSelectItem } from './LocationSelectItem'

export type LocationSelectProps = Pick<SelectProps, 'mt' | 'label' | 'classNames'>

export const LocationSelect = (props: LocationSelectProps) => {
  const { options } = useLocationSelectLogic()

  return (
    <Select
      data={options}
      itemComponent={LocationSelectItem}
      dropdownPosition='bottom'
      searchable
      clearable
      withinPortal
      {...props}
    />
  )
}
