import React, { Component } from 'react';
import {
  Text,
  View,
  ScrollView
} from 'react-native';
import { List, ListItem, Button } from 'react-native-elements';
import { BigParkUsers } from '../config/data';

class DogListScreen extends Component {
    render() {
        return (
          <ScrollView>
            <List>
              {BigParkUsers.map((user) => (
                <ListItem
                  key={user.username}
                  roundAvatar
                  avatar={{ uri: user.pets[0].picture }}
                  title={user.pets[0].name}
                  subtitle={`${user.name.first} ${user.name.last}`}
                  onPress={() => {}}
                />
              ))}
            </List>
            <Button
              title="Check In A Pet!"
              buttonStyle={{ marginTop: 20, marginBottom:50 }}
              onPress={()=>this.props.navigation.navigate('MyPetList')}
            />
            
          </ScrollView>
        );
      }
    }

export default DogListScreen