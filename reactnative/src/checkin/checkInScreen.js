import React, { Component } from 'react'
import { View, TouchableOpacity, Keyboard, Text, Input, Image, Dimensions, Button, StyleSheet} from 'react-native'

const dim = Dimensions.get('window')

class CheckInScreen extends Component{
    render(){
        return(
            <View style={{flex:1, flexDirection:'column', backgroundColor:'black',alignItems:'center'}}>
                <Image 
                source={require('../images/DogParkTransparent.png')}
                style={{height:dim.height/2}}resizeMode='contain'
                />
                

            </View>
            )
    }
}
const styles = StyleSheet.create({
    buttonStyle:{
        padding:50,
        margin:50
    }
})
export default CheckInScreen