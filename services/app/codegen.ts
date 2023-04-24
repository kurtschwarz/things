import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: '../api/graph/**/*.graphqls',
  generates: {
    'src/graphql/generated/types.ts': {
      documents: [
        'src/**/*.ts'
      ],
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-react-apollo'
      ],
      config: {
        useTypeImports: true,
        flattenGeneratedTypes: true,
        futureProofUnions: true,
        futureProofEnums: true,
        withHooks: true,
        withComponents: false
      }
    },
    'src/graphql/generated/helpers.ts': {
      plugins: [
        'typescript-apollo-client-helpers'
      ]
    }
  },
  verbose: true,
  overwrite: true,
}

export default config
