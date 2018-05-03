import React, { Component } from 'react'
import { View, Keyboard, Text, Image, Dimensions, StyleSheet } from 'react-native'
import { TextButton, Input } from '../shared/ui'
import PropTypes from 'prop-types'
import firebase from '../firebase'

const dim = Dimensions.get('window')

class LoginScreen extends Component {
  state = { email: '', password: '', error: '', loading: false, keyboard: false }
  componentDidMount () {
    this.keyboardDidShowListener = Keyboard.addListener(
      'keyboardDidShow',
      this._keyboardDidShow.bind(this)
    )
    this.keyboardDidHideListener = Keyboard.addListener(
      'keyboardDidHide',
      this._keyboardDidHide.bind(this)
    )
  }

  componentWillUnmount () {
    this.keyboardDidShowListener.remove()
    this.keyboardDidHideListener.remove()
  }

  _keyboardDidShow () {
    this.setState({ keyboard: true })
  }

  _keyboardDidHide () {
    this.setState({ keyboard: false })
  }

  _login () {
    const { email, password } = this.state
    this.setState({ error: '', loading: true })
    firebase
      .auth()
      .signInWithEmailAndPassword(email, password)
      .then(this.onLoginSuccess.bind(this))
      .catch((error) => {
        console.log('login error:' + error)
        error = null
        firebase
          .auth()
          .createUserWithEmailAndPassword(email, password)
          .then(this.onLoginSuccess.bind(this))
          .catch((error) => {
            console.log('Create user error:' + error)
            this.onLoginFail.bind(this)
          })
      })
  }

  onLoginFail () {
    console.log('Login failed: ' + error)
    this.setState({
      email: this.state.email,
      password: '',
      error: 'Authentication Failed',
      loading: false
    })
  }

  onLoginSuccess () {
    this.setState({
      email: '',
      password: '',
      loading: false,
      error: ''
    })
    this.props.navigation.navigate('Home')
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
        {!this.state.keyboard && (
          <Image
            source={require('../images/DogParkTransparent.png')}
            style={{ height: dim.height / 2 }}
            resizeMode='contain'
          />
        )}
        <Input
          placeholder='user@email.com'
          label='Email: '
          value={this.state.email}
          onChangeText={(email) => this.setState({ email })}
        />
        <Input
          secureTextEntry
          placeholder='password'
          label='Password: '
          value={this.state.password}
          onChangeText={(password) => this.setState({ password })}
        />
        <View
          style={{ width: dim.width, flexDirection: 'row', justifyContent: 'center' }}
        >
          <TextButton style={styles.buttonStyle} onPress={() => this._login()}>
            Sign in
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

/*var { FBLoginManager } = require('react-native-facebook-login')
//FBLoginManager.setLoginBehavior(FBLoginManager.LoginBehaviors.Web) // defaults to Native
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
*/
