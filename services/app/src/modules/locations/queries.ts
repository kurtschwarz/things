import { gql } from '@apollo/client'

export const GET_LOCATIONS = gql`
  query {
    locations {
      edges {
        node {
          id
          name
        }
      }
    }
  }
`
