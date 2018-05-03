import { createStore, combineReducers, applyMiddleware } from 'redux'
import { autoRehydrate } from 'redux-persist'
import { reducer } from './src/store/reducer'
import { composeWithDevTools } from 'redux-devtools-extension' // maybe

const enhancer = composeWithDevTools(middleware, autoRehydrate()) //maybe
export const store = createStore(reducer, {}, enhancer)
