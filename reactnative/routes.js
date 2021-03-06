import React, { Component } from 'react'
import { StackNavigator } from 'react-navigation'
import { LoginScreen } from './src/login/loginScreen'
import HomeScreen from './src/home/homeScreen'
import CheckInScreen from './src/checkin/checkInScreen'
import { ProfileScreen } from './src/profile/profileScreen'
import DogListScreen from './src/dogList/dogListScreen'
import MyPetListScreen from './src/dogList/myPetListScreen'
import AddPetScreen from './src/profile/addPetScreen'

export const Shika = StackNavigator({
  Login: {
    screen: LoginScreen,
    navigationOptions: {
      header: null
    }
  },
  Home: {
    screen: HomeScreen,
    navigationOptions: {
      header: null
    }
  },
  Profile: {
    screen: ProfileScreen,
    navigationOptions: {
      header: null
    }
  },
  CheckIn: {
    screen: CheckInScreen,
    navigationOptions: {
      header: null
    }
  },
  DogList: {
    screen: DogListScreen,
    navigationOptions: {
      header: null
    }
  },
  MyPetList: {
    screen: MyPetListScreen,
    navigationOptions: {
      header: null
    }
  },
  AddPet: {
    screen: AddPetScreen,
    navigationOptions: {
      header: null
    }
  }
})
