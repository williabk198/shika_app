/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 * @flow
 */

import React, { Component } from 'react';
import {
  Text,
  View,
  AppRegistry,
} from 'react-native';
import {Root} from './routes'

class App extends Component {
  render() {
    return <Root />;
  }
}

AppRegistry.registerComponent('shika', () => App)