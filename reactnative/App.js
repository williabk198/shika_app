/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 * @flow
 */

import React, { Component } from 'react'
import { Text, View, AppRegistry, AsyncStorage } from 'react-native'
import { persistStore, autoRehydrate } from 'redux-persist'
import { createStore } from 'redux'
import { Provider } from 'react-redux'
import { Root } from './routes'

import { reducer } from './src/store/reducer'

const store = createStore(reducer, undefined, autoRehydrate())

persistStore(store, { storage: AsyncStorage })

class App extends Component {
  render () {
    return (
      <Provider store={store}>
        <Root />
      </Provider>
    )
  }
}

AppRegistry.registerComponent('shika', () => App)
