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
      var me = null
      if (this.props.navigation.state.params){
         me = this.props.navigation.state.params;
      }
        return (
          <ScrollView style={{backgroundColor:'#C2B280'}}>
            <List>
              {BigParkUsers.map((user) => (
                <ListItem
                  key={user.username}
                  roundAvatar
                  avatar={{ uri: user.pets[0].picture }}
                  title={user.pets[0].name}
                  hideChevron
                  rightTitle={user.pets[0].gender}
                  subtitle={`${user.name.first} ${user.name.last}`}
                  onPress={() => {}}
                />
              ))}
              {me && me.pet && me.pet.picture &&
              <ListItem
                roundAvatar
                avatar={{ uri: me.pet.picture }}
                title={me.pet.name}
                subtitle={`${me.name.first} ${me.name.last}`}
                onPress={() => {}}
              />
              }
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