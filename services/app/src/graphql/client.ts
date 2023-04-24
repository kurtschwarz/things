import { ApolloClient, InMemoryCache } from '@apollo/client'
import { relayStylePagination } from '@apollo/client/utilities'

import { StrictTypedTypePolicies } from '@/graphql'

export const typePolicies: StrictTypedTypePolicies = {
  Query: {
    fields: {
      locations: relayStylePagination()
    }
  }
}

export const client = new ApolloClient({
  uri: 'http://127.0.0.1:8081/graphql',
  cache: new InMemoryCache({
    typePolicies
  })
})
