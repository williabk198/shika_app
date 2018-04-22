import React, { Component } from 'react';
import {
  Text,
  View,
  ScrollView
} from 'react-native';
import { List, ListItem, Button } from 'react-native-elements';
import { BigParkUsers } from '../config/data';
import {goBackTo} from '../config/util';

class DogListScreen extends Component {
    render() {
      console.log(this.props)
      var me = null
      if (this.props.navigation.state.params){
         me = this.props.navigation.state.params;
      }
        return (
          <ScrollView style={{backgroundColor:'#C2B280'}}>
            <List>
              {BigParkUsers.map((user,index) => (
                <ListItem
                  key={index}
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
                style={{color:'#C2B280'}}
                avatar={{ uri: me.pet.picture }}
                title={me.pet.name}
                rightTitle={me.pet.gender}
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
              {me && me.pet &&<Button
              title="Check Out"
              buttonStyle={{ marginTop: 20, marginBottom:50 }}
              onPress={()=>goBackTo('DogList',this.props.navigation)}
            />}
          </ScrollView>
        );
      }
    }

export default DogListScreen