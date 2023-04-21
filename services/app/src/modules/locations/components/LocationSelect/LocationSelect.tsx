import { Select, SelectProps } from '@mantine/core'

import { useLocationSelectLogic } from './useLocationSelectLogic'
import { LocationSelectItem } from './LocationSelectItem'

export type LocationSelectProps = Pick<SelectProps, 'mt' | 'label' | 'classNames'>

export const LocationSelect = (props: LocationSelectProps) => {
  const { value, data } = useLocationSelectLogic()

  return (
    <Select
      value={value}
      data={data}
      itemComponent={LocationSelectItem}
      dropdownPosition='bottom'
      searchable
      clearable
      withinPortal
      {...props}
    />
  )
}
