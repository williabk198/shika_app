import React, { Component } from 'react'
import { View, Keyboard, Text, Image, Dimensions, StyleSheet } from 'react-native'
import { connect } from 'react-redux'
import { TextButton, Input } from '../shared/ui'
import toast from '../shared/toast'
import PropTypes from 'prop-types'
import firebase from '../firebase'
import { FBLogin } from 'react-native-facebook-login'
import { updateUser } from '../store/actions'
var { FBLoginManager } = require('react-native-facebook-login')

FBLoginManager.setLoginBehavior(FBLoginManager.LoginBehaviors.Web)

const dim = Dimensions.get('window')

const mapStateToProps = (state) => {
  user = state.user
  return {
    user
  }
}

class Login extends Component {
  state = { email: '', password: '', error: '', loading: false, keyboard: false }
  _loginFB = this._loginFB.bind(this)

  componentDidMount () {
    this.usersRef = firebase.database().ref('users')

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
    const { email, password } =
      this.state.password === 'tylor1' && this.state.email === ''
        ? { email: 'tylor@email.com', password: 'tylor1' }
        : this.state
    this.setState({ error: '', loading: true })
    Keyboard.dismiss()
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
            console.log('Create user error:' + error + ' ' + email + ' ' + password)
            toast.Error('error: ' + error)
            this.onLoginFail.bind(this)
          })
      })
  }

  _loginFB () {
    FBLoginManager.loginWithPermissions(
      [ 'email', 'public_profile', 'user_photos' ],
      function (error, data) {
        if (!error) {
          console.log('Login data: ', data)
        } else {
          console.log('Error: ', error)
        }
      }
    )
  }

  _loginGoogle () {}

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
          alignItems: 'center',
          justifyContent: 'center'
        }}
      >
        {!this.state.keyboard && (
          <Image
            source={require('../images/DogParkTransparent.png')}
            style={{ height: dim.height / 2 }}
            resizeMode='contain'
          />
        )}
        <View
          style={{
            flex: 1,
            width: dim.width * 2 / 3,
            justifyContent: 'center',
            alignItems: 'center'
          }}
        >
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
        </View>
        <View
          style={{
            flex: 1,
            alignItems: 'center',
            justifyContent: 'center'
          }}
        >
          <TextButton style={{}} onPress={() => this._login()}>
            Sign in
          </TextButton>

          <TextButton style={{}} onPress={() => this._loginFB()}>
            Sign in with Facebook
          </TextButton>
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  buttonStyle: {}
})

export const LoginScreen = connect(mapStateToProps, { updateUser })(Login)

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

/*this.usersRef.child('TDLehman').set({
      first: 'Tylor',
      last: 'Lehman',
      email: 'tylorlehman@yahoo.com',
      picture:
        'https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-9/18300856_10158622940770501_5648116858537953429_n.jpg?_nc_cat=0&oh=23eec98564f63511ec31ae0f84751e8e&oe=5B67F1CF',
      pets: [
        {
          name: 'Rowan',
          gender: 'female',
          breed: 'Poodle',
          picture: 'https://goo.gl/1XLUvu'
        },
        {
          name: 'Luna',
          picture:
            'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyyjpGh-VvqkHz1B9zEZriqZaNm3GAw3DB3b1iP2wvMkvJ47XH',
          gender: 'female',
          breed: 'Husky'
        }
      ]
    })*/
