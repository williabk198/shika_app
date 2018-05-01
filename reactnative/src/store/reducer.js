import { REFRESH_VISITORS, UPDATE_PROFILE, LOADING } from './type'

const INITIAL_STATE = {
  loading: false,
  user: {
    userName: 'TDLehman',
    name: {
      first: 'Tylor',
      last: 'Lehman'
    },
    pets: {
      Rowan: {
        gender: 'female',
        type: 'dog',
        breed: 'Poodle',
        picture: 'https://goo.gl/1XLUvu'
      },
      Luna: {
        gender: 'female',
        type: 'dog',
        breed: 'Husky',
        picture:
          'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyyjpGh-VvqkHz1B9zEZriqZaNm3GAw3DB3b1iP2wvMkvJ47XH'
      }
    }
  },
  visitors: {
    lastRefresh: '04/30/2018',
    LargePark: {
      williabk198: {
        name: {
          first: 'Brandon',
          last: 'Williams'
        },
        pets: {
          name: 'Harley',
          picture: 'https://goo.gl/PPefEQ',
          gender: 'male',
          breed: 'pitbull',
          type: 'dog',
          park: 'large'
        }
      }
    },
    LittlePark: {}
  }
}

export const reducer = (state = INITIAL_STATE, { type, payload }) => {
  switch (type) {
    case LOADING:
      return { ...state, loading: true }
    case UPDATE_PROFILE:
      return { ...state, loading: false }
    case REFRESH_VISITORS:
      return { ...state, loading: false }
    default:
      return state
  }
}
