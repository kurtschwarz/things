import { createStyles, rem, Group, ActionIcon, useMantineColorScheme } from '@mantine/core'
import { BiMoon, BiSun } from 'react-icons/bi'

export const ColorSchemeToggle = () => {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme()
  const Icon = colorScheme === 'dark' ? BiSun : BiMoon

  return (
    <Group>
      <ActionIcon
        onClick={() => toggleColorScheme()}
        size='lg'
        sx={(theme) => ({
          backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
          color: theme.colorScheme === 'dark' ? theme.colors.yellow[4] : theme.colors.blue[6],
        })}
      >
        <Icon />
      </ActionIcon>
    </Group>
  )
}
