import { useState } from 'react'
import { ColorScheme, ColorSchemeProvider as MantineColorSchemeProvider } from '@mantine/core'

type ColorSchemeProviderProps = {
  children: (colorScheme: ColorScheme) => React.ReactNode
}

export const ColorSchemeProvider = (props: ColorSchemeProviderProps) => {
  const [colorScheme, setColorScheme] = useState<ColorScheme>('light')
  const toggleColorScheme = (newColorScheme?: ColorScheme) =>
    setColorScheme(newColorScheme || (colorScheme === 'dark' ? 'light' : 'dark'))

  return (
    <MantineColorSchemeProvider
      colorScheme={colorScheme}
      toggleColorScheme={toggleColorScheme}
    >
      {props.children(colorScheme)}
    </MantineColorSchemeProvider>
  )
}
