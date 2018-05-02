import React, { Component } from 'react'
import { View, Keyboard, Text, Input, Image, Dimensions, StyleSheet } from 'react-native'
import { TextButton } from '../shared/ui'

const dim = Dimensions.get('window')

class LoginScreen extends Component {
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
          <TextButton
            style={styles.buttonStyle}
            onPress={() => this.props.navigation.navigate('Home')}
          >
            Log in with Facebook
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
