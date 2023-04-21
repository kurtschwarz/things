import { gql } from '@apollo/client'

export const generateLocationParentsRecursive = (depth = 5): string => `
  parent {
    id
    name
    ${depth - 1 ? generateLocationParentsRecursive(depth - 1) : ''}
  }
`

export const GET_ALL_LOCATIONS = gql`
  query {
    locations {
      edges {
        node {
          id
          name
          parentID
        }
      }
    }
  }
`

export const GET_ALL_LOCATION_PARENTS_RECURSIVE = gql`
  query {
    locations {
      edges {
        node {
          id
          name
          parentID
          ${generateLocationParentsRecursive()}
        }
      }
    }
  }
`
