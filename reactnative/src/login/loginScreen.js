import React, { Component } from 'react'
import { View, Keyboard, Text, Input, Image, Dimensions, StyleSheet } from 'react-native'
import { TextButton } from '../shared/ui'
import { FBLogin } from 'react-native-facebook-login'
import PropTypes from 'prop-types'
const dim = Dimensions.get('window')
var { FBLoginManager } = require('react-native-facebook-login')

FBLoginManager.setLoginBehavior(FBLoginManager.LoginBehaviors.Web) // defaults to Native

class LoginScreen extends Component {
  _login () {
    console.log('trying to login')
    FBLoginManager.loginWithPermissions([ 'email', 'user_friends' ], function (
      error,
      data
    ) {
      if (!error) {
        console.log('Login data: ', data)
        this.props.navigation.navigate('Home')
      } else {
        console.log('Error: ', error)
      }
    })
  }

  render () {
    return (
      <View
        style={{
          flex: 1,
          flexDirection: 'column',
          backgroundColor: '#C2B280',
          alignItems: 'center'
        }}
      >
        <Image
          source={require('../images/DogParkTransparent.png')}
          style={{ height: dim.height / 2 }}
          resizeMode='contain'
        />
        <View
          style={{ width: dim.width, flexDirection: 'row', justifyContent: 'center' }}
        >
          <TextButton style={styles.buttonStyle} onPress={() => this._login()}>
            Sign in with Facebook
          </TextButton>
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  buttonStyle: {
    padding: 50,
    margin: 50
  }
})

export default LoginScreen
