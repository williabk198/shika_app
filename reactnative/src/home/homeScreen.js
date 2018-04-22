import React, { Component } from 'react'
import { View, Keyboard, Text, Input, Image, Dimensions, TouchableHighlight, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
const dim = Dimensions.get('window')

class HomeScreen extends Component{
    render(){
        return(
            <View style={{flex:1, flexDirection:'column', backgroundColor:'black',alignItems:'center'}}>
                <Image 
                    source={require('../images/DogParkTransparent.png')}
                    style={{height:dim.height/4}}resizeMode='contain'
                />
                <View style={{flexDirection:'row',justifyContent:'center'}}>
                <HomeButton 
                    onPress={()=>this.props.navigation.navigate('CheckIn')}
                    >Check In</HomeButton></View>
                    <View style={{flexDirection:'row',justifyContent:'center'}}>
                <HomeButton 
                    onPress={()=>this.props.navigation.navigate('Profile')}
                    >Edit Profile</HomeButton>
</View>

            </View>
        )
    }
}

export default HomeScreen