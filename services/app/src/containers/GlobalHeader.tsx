import { rem, createStyles, Header, Group, Text, Kbd, UnstyledButton } from '@mantine/core'
import { useSpotlight } from '@mantine/spotlight'
import { BiSearch } from 'react-icons/bi'

import { ColorSchemeToggle } from '@/modules/theme'

const useStyles = createStyles((theme) => ({
  searchButton: {
    userSelect: 'none',
    cursor: 'pointer',
    border: `${rem(1)} solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[4]}`,
    borderRadius: theme.radius.md,
    padding: `0 ${theme.spacing.sm}`,
    height: rem(34),
    fontSize: theme.fontSizes.sm,

    '& svg': {
      marginTop: rem(2)
    }
  },

  searchButtonLabel: {
    width: rem(238)
  }
}))

type GlobalHeaderProps = {
  children?: React.ReactNode
}

export const GlobalHeader = (props: GlobalHeaderProps) => {
  const { classes } = useStyles()
  const spotlight = useSpotlight()

  const handleSearchClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    spotlight.openSpotlight()
  }

  return (
    <Header height={{ base: 61 }}>
      <Group style={{ height: 61 }} pl='md' pr='md' position='right'>
        <UnstyledButton className={classes.searchButton} onClick={handleSearchClick}>
          <Group position='left' align='center' spacing='xs'>
            <BiSearch />
            <Text size='sm' className={classes.searchButtonLabel}>Search</Text>
            <Kbd size='xs'>/</Kbd>
          </Group>
        </UnstyledButton>

        <ColorSchemeToggle />
      </Group>
    </Header>
  )
}
