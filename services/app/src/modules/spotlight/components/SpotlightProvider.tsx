import { openCreateLocationMutationModal } from '@/modules/locations/helpers'
import { SpotlightProvider as MantineSpotlightProvider, SpotlightAction } from '@mantine/spotlight'
import { BiSearch, BiMap, BiLabel, BiPlusCircle } from 'react-icons/bi'
import { useNavigate, useRevalidator } from 'react-router-dom'

type SpotlightProviderProps = {
  children: React.ReactNode
}

export const SpotlightProvider = ({
  children
}: SpotlightProviderProps) => {
  const navigate = useNavigate()
  const revalidator = useRevalidator()

  const shortcut = ['mod + P', '/']
  const actions: SpotlightAction[] = [
    {
      icon: <BiMap />,
      group: 'pages',
      title: 'Locations',
      description: 'See all locations',
      onTrigger: () => navigate('/locations'),
    },
    {
      icon: <BiLabel />,
      group: 'pages',
      title: 'Tags',
      description: 'See all tags',
      onTrigger: () => navigate('/tags'),
    },

    {
      icon: <BiPlusCircle />,
      group: 'actions',
      title: 'New Location',
      description: 'Add a new location',
      onTrigger: () => openCreateLocationMutationModal({
        onCompleted: () => revalidator.revalidate()
      }),
    },
  ]

  return (
    <MantineSpotlightProvider
      shortcut={shortcut}
      actions={actions}
      searchIcon={<BiSearch />}
      searchPlaceholder='Search...'
    >
      {children}
    </MantineSpotlightProvider>
  )
}
