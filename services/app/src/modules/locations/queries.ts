import { gql } from '@/graphql'

export const GET_LOCATION = gql`
  query GetLocation ($id: ID!) {
    locations (where: { id: $id }) {
      edges {
        node {
          id
          name
        }
      }
    }
  }
`

export const GET_ALL_LOCATIONS = gql`
  query GetAllLocations {
    locations {
      edges {
        node {
          id
          name
          description
          parentID

          stats {
            totalLocations
            totalItems
            totalValue
          }

          children {
            id
            name
            description

            stats {
              totalLocations
              totalItems
              totalValue
            }

            children {
              id
              name
              description

              stats {
                totalLocations
                totalItems
                totalValue
              }

              children {
                id
                name
                description

                stats {
                  totalLocations
                  totalItems
                  totalValue
                }
              }
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
