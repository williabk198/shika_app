import React, { Component } from 'react';
import {
  Text,
  View,
  ScrollView
} from 'react-native';
import { List, ListItem, Button } from 'react-native-elements';
import { me } from '../config/data';

class myPetListScreen extends Component {
    render() {
        return (
          <ScrollView style={{backgroundColor:'#C2B280'}}>
            <List>
              {me.pets.map((pet, index) => (
                <ListItem
                  key={index}
                  roundAvatar
                  avatar={{ uri: pet.picture }}
                  title={pet.name}
                  hideChevron
                  rightTitle={pet.gender}
                  subtitle={pet.breed}
                  onPress={() => this.props.navigation.navigate('DogList',{name:{first:me.name.first,last:me.name.last,pet}})}
                />
              ))}
            </List>
          </ScrollView>
        );
      }
    }
    
export default myPetListScreen