import path from 'path'
import { fileURLToPath, URL } from 'url'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import codegen from 'vite-plugin-graphql-codegen'

export default defineConfig({
  plugins: [
    react(),
    codegen({
      configFilePathOverride: path.join(__dirname, 'codegen.ts')
    })
  ],
  server: {
    host: '0.0.0.0',
    port: 3000
  },
  resolve: {
    alias: [
      { find: '@', replacement: fileURLToPath(new URL('./src', import.meta.url)) },
    ]
  },
  define: {
    APP_VERSION: JSON.stringify('0.0.0-alpha')
  }
})
