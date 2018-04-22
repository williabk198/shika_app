import React, { Component } from 'react'
import { View, TouchableOpacity, Keyboard, Text, Input, Image, Dimensions, Button, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
const dim = Dimensions.get('window')

class CheckInScreen extends Component{
    render(){
        return(
            <View style={{flex:1, flexDirection:'column', backgroundColor:'#C2B280',alignItems:'center'}}>
                <Image 
                    source={require('../images/DogParkTransparent.png')}
                    style={{height:dim.height/4}}resizeMode='contain'
                />
                <View style={{flexDirection:'row',justifyContent:'center'}}>
                <HomeButton 
                    onPress={()=>this.props.navigation.navigate('DogList','Large')}
                    >Large Dog Park</HomeButton></View>
                    <View style={{flexDirection:'row',justifyContent:'center'}}>
                <HomeButton 
                    onPress={()=>this.props.navigation.navigate('DogList','Small')}
                    >Small Dog Park</HomeButton>
</View>

            </View>
        )
    }
}
export default CheckInScreen