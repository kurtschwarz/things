import { ActionIcon, Box, Group, Text, createStyles } from '@mantine/core'
import { useState } from 'react'
import { BiChevronUp, BiChevronDown } from 'react-icons/bi'

export type LocationsGridGroupProps = {
  name?: string | null
  divider?: boolean
  children: React.ReactNode
}

const useStyles = createStyles<string, LocationsGridGroupProps>((theme, props) => ({
  root: {
    width: '100%',
    userSelect: 'none',
    cursor: 'pointer',
    borderTop: props.divider ? `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[2]}` : 'none'
  }
}))

export const LocationsGridGroup = (props: LocationsGridGroupProps) => {
  const { classes } = useStyles(props)
  const [visible, setVisible] = useState(true)

  const handleClick = (event: React.MouseEvent<HTMLDivElement>) => {
    event.preventDefault()
    event.stopPropagation()

    setVisible(!visible)
  }

  return (
    <Box mb={visible ? 'lg' : undefined}>
      <Group
        pb={visible ? 'md' : undefined}
        pt={props.divider ? 'sm' : undefined}
        className={classes.root}
        position='apart'
        onClick={handleClick}
      >
        <Text transform='uppercase' opacity={0.35} fz={'sm'} fw={500}>{props.name}</Text>
        <ActionIcon>
          {visible ? <BiChevronDown /> : <BiChevronUp/>}
        </ActionIcon>
      </Group>

      {visible ? props.children : null}
    </Box>
  )
}
