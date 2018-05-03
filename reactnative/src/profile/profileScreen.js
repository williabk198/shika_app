import React, { Component } from 'react'
import { View, Keyboard, Text, Image, Dimensions, ScrollView, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
import { Tile, List, ListItem, Button } from 'react-native-elements';
import { me } from '../config/data';

const dim = Dimensions.get('window')

class ProfileScreen extends Component{
    render(){
        return(
            <ScrollView style={{backgroundColor:'#C2B280'}}>
            <Tile
              imageSrc={{ uri: this.props.picture}}
              featured
              title={`${this.props.name.first.toUpperCase()} ${this.props.name.last.toUpperCase()}`}
              caption={this.props.email}
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
                rightTitle={this.props.username}
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
                            <Button
              title="Add A Pet"
              buttonStyle={{ marginTop: 20, marginBottom:50 }}
              onPress={()=>this.props.navigation.navigate('AddPet')}
            />
    
            </List>
          </ScrollView>
        );
      }
    }
    
    ProfileScreen.defaultProps = { ...me };
    
export default ProfileScreen