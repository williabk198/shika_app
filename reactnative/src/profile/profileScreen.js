import React, { Component } from 'react'
import {
  View,
  Keyboard,
  Text,
  Image,
  Dimensions,
  ScrollView,
  StyleSheet
} from 'react-native'
import { HomeButton } from '../shared/ui'
import { Tile, List, ListItem, Button } from 'react-native-elements'
import { me } from '../config/data'

const dim = Dimensions.get('window')

class ProfileScreen extends Component {
  constructor (props) {
    super(props)
    this.state = {
      first: '',
      last: '',
      picture: '',
      email: '',
      pets: [ { name: '', breed: '' } ]
    }
  }

  componentDidMount () {
    this.state = this.props.user
  }
  render () {
    return (
      <ScrollView style={{ backgroundColor: '#C2B280' }}>
        <Tile
          imageSrc={{ uri: this.state.picture }}
          featured
          title={`${this.state.first.toUpperCase()} ${this.state.last.toUpperCase()}`}
          caption={this.state.email}
        />
        <List>
          <ListItem title='First' rightTitle={this.state.first} hideChevron />
          <ListItem title='Last' rightTitle={this.state.last} hideChevron />
          <ListItem title='Email' rightTitle={this.state.email} hideChevron />
        </List>
        <List>
          <Text style={{ weight: 'bold' }}>Pets</Text>
          {this.state.pets.map((pet, index) => {
            return (
              <ListItem key={index} title={pet.name} rightTitle={pet.breed} hideChevron />
            )
          })}
          <Button
            title='Add A Pet'
            buttonStyle={{ marginTop: 20, marginBottom: 50 }}
            onPress={() => this.props.navigation.navigate('AddPet')}
          />
        </List>
      </ScrollView>
    )
  }
}

ProfileScreen.defaultProps = { ...me }

export default ProfileScreen
