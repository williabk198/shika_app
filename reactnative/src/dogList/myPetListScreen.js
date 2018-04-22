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
          <ScrollView>
            <List>
              {me.pets.map((pet, index) => (
                <ListItem
                  key={index}
                  roundAvatar
                  avatar={{ uri: pet.picture }}
                  title={pet.name}
                  subtitle={pet.breed}
                  onPress={() => {}}
                />
              ))}
            </List>
          </ScrollView>
        );
      }
    }
    
export default myPetListScreen