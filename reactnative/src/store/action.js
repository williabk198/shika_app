import moment from 'moment'
import { REFRESH_VISITORS, UPDATE_PROFILE, LOADING } from './type'

export const refreshVisitors = () => {
  return async (dispatch) => {
    dispatch({ type: LOADING })
    try {
      const visitors = await api.visitors.get()
      dispatch(refreshVisitorsSuccess(visitors))
    } catch (error) {
      console.log(`error refreshing visitors: ${error.message}`)
      dispatch({ type: REFRESH_VISITORS })
    }
  }
}

export const loading = () =>  {
  console.log('loading')
  return({type:LOADING})}

export const updateUser = (user) => {
  console.log('should be updating user')
  return async (dispatch) => {
    dispatch({ type: LOADING })
    delete user.picture
    dispatch({ type: UPDATE_PROFILE, payload: user })
    try {
      //update firebase here. Should actually be before profile stuff
    } catch (error) {
      console.log(`error updating profile: ${error.message}`)
    }
  }
}
