import React from 'react'
import { StyleSheet, TouchableWithoutFeedback } from 'react-native'
import { Button, Text } from 'native-base'
import PropTypes from 'prop-types'

export const TextButton = ({ children, onPress, style }) => (
  <Button
    onPress={onPress}
    style={[
      {
        width: 200,
        marginBottom: 10,
        backgroundColor: '#334477',
        justifyContent: 'center'
      },
      style
    ]}
  >
    <Text style={{ color: 'white', fontWeight: 'bold' }}>{children}</Text>
  </Button>
)
TextButton.propTypes = {
  children: PropTypes.any,
  onPress: PropTypes.func.isRequired
}
TextButton.defaultProps = {
  children: 'Click Me'
}

export const HomeButton = ({ children, onPress, style }) => (
  <Button
    onPress={onPress}
    style={{
      width: 200,
      marginBottom: 10,
      backgroundColor: 'yellow',
      justifyContent: 'center'
    }}
  >
    <Text style={{ color: 'black', fontWeight: 'bold' }}>{children}</Text>
  </Button>
)
TextButton.propTypes = {
  children: PropTypes.any,
  onPress: PropTypes.func.isRequired
}
TextButton.defaultProps = {
  children: 'Click Me'
}
