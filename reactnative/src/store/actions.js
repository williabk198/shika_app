import moment from 'moment'
import { REFRESH_VISITORS, UPDATE_PROFILE, LOADING } from './types'

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

export const updateProfile = (profile) => {
  return async (dispatch) => {
    dispatch({ type: LOADING })
    try {
      const response = await api.profile.get(profile)
    } catch (error) {
      console.log(`error updating profile: ${error.message}`)
      dispatch({ type: UPDATE_PROFILE, payload: { profile } })
    }
  }
}
