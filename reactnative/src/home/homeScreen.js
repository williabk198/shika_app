import React, { Component } from 'react'
import { View,ScrollView, Keyboard, Text,Image, Dimensions, TouchableHighlight, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
import { ListItem, List } from 'react-native-elements';
import {Item} from 'native-base'


const dim = Dimensions.get('window')

class HomeScreen extends Component{
    render(){
        return(
            <View style={{flex:1, flexDirection:'column', backgroundColor:'#C2B280',alignItems:'center'}}>
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
                    >Profile</HomeButton>
                </View>
                <View style={{flexDirection:'row',justifyContent:'center'}}>
                <HomeButton 
                    onPress={()=>console.log('Settings button is not implemented')}
                    >Settings</HomeButton>
                </View>
                <Item><Text>Hours: Monday-Sunday 6:00AM-8:00PM</Text></Item>
                <View style={{flexDirection:'row',justifyContent:'center'}}>
                <ScrollView>
                <View style={{flexDirection:'row',justifyContent:'center',marginTop:20}}>
                    <Text style={{fontWeight:'bold', fontSize:20}}>Upcoming Events:</Text></View>
                    <List style={{padding:10, flex:1}}>
                        <ListItem
                            title='New Mobile App Coming Soon!'
                            hideChevron
                            subtitle='Want to plan your visits around other guests at our dog park? Now you can!'
                            onPress={() => {}}
                        />
                        <ListItem
                            title='Dog Training Workshops'
                            hideChevron
                            subtitle='April 28th - Puppy Socialization (only for puppies less than 5 months old) '
                            onPress={() => {}}
                        />

                    </List>
                </ScrollView>
                </View>
            </View>
        )
    }
}

export default HomeScreen