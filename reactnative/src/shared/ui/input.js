import React from 'react'
import { TextInput, View, Text, Dimensions } from 'react-native'
const dim = Dimensions.get('window')
export const Input = ({ label, value, onChangeText, placeholder, secureTextEntry }) => {
  return (
    <View
      style={{
        flex: 1,
        flexDirection: 'column'
      }}
    >
      <Text
        style={{
          fontSize: 18,
          paddingLeft: 20,
          flex: 1
        }}
      >
        {label}
      </Text>
      <TextInput
        secureTextEntry={secureTextEntry}
        placeholder={placeholder}
        autoCorrect={false}
        style={{
          color: '#000',
          paddingRight: 15,
          paddingLeft: 15,
          fontSize: 15,
          lineHeight: 23,
          width: dim.width
        }}
        value={value}
        onChangeText={onChangeText}
      />
    </View>
  )
}
