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
import { connect } from 'react-redux'
import { HomeButton } from '../shared/ui'
import { Tile, List, ListItem, Button } from 'react-native-elements'
import { updateUser } from '../store/actions'
const dim = Dimensions.get('window')

const mapStateToProps = (state) => {
  user = state.user
  return {
    user
  }
}

class Profile extends Component {
  constructor (props) {
    super(props)
    this.state = {
      first: 'First',
      last: 'Last',
      picture: 'http://via.placeholder.com/350x350?text=Image',
      email: 'Email',
      pets: [ { name: 'petName', breed: 'Breed' } ]
    }
  }
  componentDidMount () {
    const user = this.props.user ? this.props.user : ''
    this.props.user && this.setState(user)
  }
  render () {
    console.log(this.props.user)
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
          <Text style={{ fontWeight: 'bold', marginLeft: 10 }}>Pets</Text>
          {this.state.pets.map((pet, index) => {
            return (
              <ListItem key={index} title={pet.name} rightTitle={pet.breed} hideChevron />
            )
          })}
        </List>
        <Button
          title='Add A Pet'
          buttonStyle={{
            flex: 1,
            justifyContent: 'center',
            margin: 5,
            marginTop: 20,
            width: dim.width * 2 / 3
          }}
          onPress={() => this.props.navigation.navigate('AddPet')}
        />
        <Button
          title='Save Changes'
          buttonStyle={{ margin: 5, width: dim.width * 2 / 3 }}
          onPress={() => {}}
        />
      </ScrollView>
    )
  }
}

export const ProfileScreen = connect(mapStateToProps, {
  updateUser
})(Profile)
