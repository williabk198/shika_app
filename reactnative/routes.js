import React, { Component } from 'react'
import { StackNavigator } from 'react-navigation'
import  LoginScreen  from './src/login/loginScreen'
import HomeScreen from './src/home/homeScreen'
import CheckInScreen from './src/checkin/checkInScreen'
import ProfileScreen from './src/profile/profileScreen'
export const Root = StackNavigator({
    Login:{
        screen: LoginScreen,
        navigationOptions: {
            header: null
        }
    },
    Home:{
        screen: HomeScreen,
        navigationoptions:{
            header: null
        }
    },
        Profile:{
        screen:ProfileScreen,
        navigationOptions:{
            header:null,
        }
    }
/*    CheckIn:{
        screen:CheckInScreen,
        navigationOptions:{
            title:'Check In',
        },
    },
*/
})
