import { gql } from '@/graphql'

export const GET_ALL_LOCATIONS = gql`
  query GetAllLocations {
    locations {
      edges {
        node {
          id
          name
          description
          parentID
          children {
            id
            name
            description
            children {
              id
              name
              description
            }
            children {
              id
              name
              description
            }
          }
        }
      }
    }
  }
`

export const GET_ALL_LOCATION_PARENTS_RECURSIVE = gql`
  query GetAllLocationsWithParents {
    locations {
      edges {
        node {
          id
          name
          parentID
          parent {
            id
            name
            parentID
            parent {
              id
              name
              parentID
              parent {
                id
                name
                parentID
                parent {
                  id
                  name
                  parentID
                }
              }
            }
          }
        }
      }
    }
  }
`
