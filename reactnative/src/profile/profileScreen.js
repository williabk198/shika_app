import React, { Component } from 'react'
import { View, Keyboard, Text, Input, Image, Dimensions, ScrollView, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
import { Tile, List, ListItem, Button } from 'react-native-elements';
import { me } from '../config/data';
const dim = Dimensions.get('window')

class ProfileScreen extends Component{
    render(){
        return(
            <ScrollView>
            <Tile
              imageSrc={{ uri: this.props.picture.large}}
              featured
              title={`${this.props.name.first.toUpperCase()} ${this.props.name.last.toUpperCase()}`}
              caption={this.props.email}
            />
    
            <Button
              title="Settings"
              buttonStyle={{ marginTop: 20 }}
              onPress={this.handleSettingsPress}
            />
    
            <List>
              <ListItem
                title="Email"
                rightTitle={this.props.email}
                hideChevron
              />
              <ListItem
                title="Phone"
                rightTitle={this.props.phone}
                hideChevron
              />
            </List>
    
            <List>
              <ListItem
                title="Username"
                rightTitle={this.props.login.username}
                hideChevron
              />
            </List>
            <List>
            {this.props.pets.map((pet,index)=>{
                    return(<ListItem
                    key={index}
                    title={pet.name}
                    rightTitle={pet.breed}
                    hideChevron
                  />)
                })}
            </List>
            <List>
              <ListItem
                title="Birthday"
                rightTitle={this.props.dob}
                hideChevron
              />
              <ListItem
                title="City"
                rightTitle={this.props.location.city}
                hideChevron
              />
            </List>
          </ScrollView>
        );
      }
    }
    
    ProfileScreen.defaultProps = { ...me };
    
export default ProfileScreen