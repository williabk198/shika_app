/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 * @flow
 */

import React, { Component } from 'react'
import { Text, View, AppRegistry } from 'react-native'
import { persistStore, autoRehydrate } from 'redux-persist'
import { Root } from './routes'

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
