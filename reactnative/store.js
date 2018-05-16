import { createStore, combineReducers, applyMiddleware } from 'redux'
import { autoRehydrate } from 'redux-persist'
import { reducer } from './src/store/reducer'
import { composeWithDevTools } from 'redux-devtools-extension' // maybe
import ReduxThunk from 'redux-thunk'

const middleware = applyMiddleware(ReduxThunk)
const appReducer = combineReducers({reducer})
const enhancer = composeWithDevTools( autoRehydrate()) //maybe

export const store = createStore(appReducer, {}, enhancer)
