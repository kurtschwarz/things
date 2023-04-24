import { createStyles, rem, getStylesRef, Navbar, Group, Code, Text } from '@mantine/core'
import { BiArchive, BiMap, BiRocket, BiCog, BiLabel } from 'react-icons/bi'
import { useLocation } from 'react-router'
import { Link } from 'react-router-dom'

const useStyles = createStyles((theme) => ({
  header: {
    paddingBottom: theme.spacing.md,
    marginBottom: `calc(${theme.spacing.md} * 1.5)`,
    borderBottom: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]
    }`
  },

  logo: {
    display: 'flex',
    alignItems: 'center',
    userSelect: 'none',
    'svg': {
      marginRight: `4px`
    }
  },

  link: {
    ...theme.fn.focusStyles(),
    display: 'flex',
    alignItems: 'center',
    textDecoration: 'none',
    fontSize: theme.fontSizes.sm,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
    padding: `${theme.spacing.xs} ${theme.spacing.sm}`,
    borderRadius: theme.radius.sm,
    fontWeight: 500,

    '&:hover': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
      color: theme.colorScheme === 'dark' ? theme.white : theme.black,

      [`& .${getStylesRef('icon')}`]: {
        color: theme.colorScheme === 'dark' ? theme.white : theme.black,
      },
    },
  },

  linkIcon: {
    ref: getStylesRef('icon'),
    color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[6],
    marginRight: theme.spacing.sm,
  },

  linkActive: {
    '&, &:hover': {
      backgroundColor: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).background,
      color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
      [`& .${getStylesRef('icon')}`]: {
        color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
      },
    },
  }
}))

const navItems = [
  {
    icon: BiRocket,
    label: 'Dashboard',
    link: '/'
  },
  {
    icon: BiMap,
    label: 'Locations',
    link: '/locations',
    active: (url: string) => url.startsWith('/location')
  },
  {
    icon: BiLabel,
    label: 'Tags',
    link: '/tags'
  },
  {
    icon: BiCog,
    label: 'Settings',
    link: '/settings'
  }
]

export const Sidebar = () => {
  const { classes, cx } = useStyles()
  const location = useLocation()

  const navLinks = navItems.map(navItem => (
    <Link
      key={navItem.label}
      to={navItem.link}
      className={cx(classes.link, { [classes.linkActive]: navItem.active ? navItem.active(location.pathname) : navItem.link === location.pathname })}
    >
      {navItem.icon && <navItem.icon className={classes.linkIcon} />}
      <span>{navItem.label}</span>
    </Link>
  ))

  return (
    <Navbar width={{ sm: 300 }} p="md">
      <Navbar.Section>
        <Group className={classes.header} position='apart'>
          <Text className={classes.logo} fz='lg' fw={700}>
            <BiArchive />
            Things
          </Text>

          <Code sx={{ fontWeight: 700 }}>v{APP_VERSION}</Code>
        </Group>
      </Navbar.Section>

      <Navbar.Section grow>
        {navLinks}
      </Navbar.Section>
    </Navbar>
  )
}
