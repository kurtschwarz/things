import { createStyles, rem, TextInput, Textarea, Button, Group } from '@mantine/core'
import { ContextModalProps } from '@mantine/modals'

import type { Location } from '@/graphql'

import { useLocationMutationModalLogic } from './useLocationMutationModalLogic'
import { LocationSelect } from '../LocationSelect/LocationSelect'

const useStyles = createStyles((theme) => ({
  root: {
    position: 'relative',
  },

  input: {
    height: rem(54),
    paddingTop: rem(18),
  },

  textarea: {
    paddingTop: rem(18),
  },

  label: {
    position: 'absolute',
    pointerEvents: 'none',
    fontSize: theme.fontSizes.xs,
    paddingLeft: theme.spacing.sm,
    paddingTop: `calc(${theme.spacing.sm} / 2)`,
    zIndex: 1,
  }
}))

export const LocationMutationModal = ({ context, id, innerProps }: ContextModalProps<{ location: Location, onCompleted?: () => void }>) => {
  const { classes } = useStyles()
  const { form, loading, handleCancel, handleSave } = useLocationMutationModalLogic({ context, id, innerProps }, innerProps.location, innerProps.onCompleted)

  return (
    <form onSubmit={form.onSubmit((values) => handleSave(values))}>
      <TextInput
        label='Name'
        placeholder='Name'
        classNames={classes}
        withAsterisk
        data-autofocus
        {...form.getInputProps('name')}
      />

      <Textarea
        label='Description'
        placeholder='Add an optional description'
        classNames={classes}
        minRows={3}
        mt='sm'
        autosize
        {...form.getInputProps('description')}
      />

      <LocationSelect
        label='Parent Location'
        classNames={classes}
        mt='sm'
        {...form.getInputProps('parentID')}
      />

      <Group position='right' mt='md'>
        <Button variant='default' onClick={handleCancel}>Cancel</Button>
        <Button type='submit' loading={loading}>Save Location</Button>
      </Group>
    </form>
  )
}

export default LocationMutationModal
