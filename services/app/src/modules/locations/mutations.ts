import { gql } from '@apollo/client'

export const UPDATE_LOCATION = gql`
  mutation ($id: ID!, $input: UpdateLocationInput!) {
    updateLocation (id: $id, input: $input) {
      id
    }
  }
`

export const CREATE_LOCATION = gql`
  mutation ($input: CreateLocationInput!) {
    createLocation (input: $input) {
      id
    }
  }
`
