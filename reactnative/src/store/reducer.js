import { REFRESH_VISITORS, UPDATE_PROFILE, LOADING } from './type'

const INITIAL_STATE = {
  loading: false,
  user: {
    first: 'Tylor',
    last: 'Lehman',
    email: 'tylorlehman@yahoo.com',
    picture:
      'https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-1/p160x160/18300856_10158622940770501_5648116858537953429_n.jpg?_nc_cat=0&oh=866174ed629b9898424eea41de4093db&oe=5B988A90',
    pets: [
      {
        name: 'Rowan',
        sex: 'female',
        breed: 'Poodle',
        picture: 'https://goo.gl/1XLUvu'
      },
      {
        name: 'Luna',
        sex: 'male',
        breed: 'Husky',
        picture:
          'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyyjpGh-VvqkHz1B9zEZriqZaNm3GAw3DB3b1iP2wvMkvJ47XH'
      }
    ]
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
          sex: 'male',
          breed: 'pitbull',
          park: 'large'
        }
      }
    },
    LittlePark: {}
  }
}

export const reducer = (state = INITIAL_STATE, { type, payload }) => {
  console.log('in reducer')
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
