import { gql } from '@apollo/client'

export const generateLocationRelationRecursive = (relation: string, depth = 5): string => `
  ${relation} {
    id
    name
    ${depth - 1 ? generateLocationRelationRecursive(relation, depth - 1) : ''}
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
          ${generateLocationRelationRecursive('children', 3)}
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
          ${generateLocationRelationRecursive('parent')}
        }
      }
    }
  }
`
