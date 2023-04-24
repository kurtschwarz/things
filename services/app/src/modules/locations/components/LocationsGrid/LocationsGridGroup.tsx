import { ActionIcon, Group, Text, createStyles } from '@mantine/core'
import { useState } from 'react'
import { BiChevronUp, BiChevronDown } from 'react-icons/bi'

export type LocationsGridGroupProps = {
  name?: string | null
  divider?: boolean
  children: React.ReactNode
}

const useStyles = createStyles((theme) => ({
  root: {
    width: '100%',
    userSelect: 'none',
    cursor: 'pointer'
  }
}))

export const LocationsGridGroup = (props: LocationsGridGroupProps) => {
  const { classes } = useStyles()
  const [visible, setVisible] = useState(true)

  const handleClick = (event: React.MouseEvent<HTMLDivElement>) => {
    event.preventDefault()
    event.stopPropagation()

    setVisible(!visible)
  }

  return (
    <div>
      <Group
        pb={visible ? 'md' : undefined}
        pt={props.divider ? 'sm' : undefined}
        className={classes.root}
        style={{ borderTop: props.divider ? '1px solid rgb(233, 236, 239)' : 'none' }}
        position='apart'
        onClick={handleClick}
      >
        <Text transform='uppercase' opacity={0.35} fz={'sm'} fw={500}>{props.name}</Text>
        <ActionIcon>
          {visible ? <BiChevronDown /> : <BiChevronUp/>}
        </ActionIcon>
      </Group>

      {visible ? props.children : null}
    </div>
  )
}
