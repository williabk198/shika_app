import React, { Component } from 'react'
import { Text, View, ScrollView, ListView } from 'react-native'
import firebase from '../firebase'
import { Item } from 'native-base'
import { List, ListItem, Button } from 'react-native-elements'
import { Me, Visitors } from '../config/data'
import { goBackTo } from '../config/util'
import { firebaseApp } from '../login/loginScreen'

class DogListScreen extends Component {
  constructor (props) {
    super(props)
    this.state = {
      dataSource: new ListView.DataSource({
        rowHasChanged: (row1, row2) => row1 !== row2
      })
    }
  }

  componentDidMount () {
    const ref = firebase.database().ref('users/users')
    this.listenForItems(this.itemsRef)
  }

  listenForItems (itemsRef) {
    itemsRef.on('value', (snap) => {
      var items = []
      snap.forEach((child) => {
        items.push({
          first: child.val().first,
          last: child.val().last,
          _key: child.key
        })
      })
      this.setState({
        dataSource: this.state.dataSource.cloneWithRows(items)
      })
    })
  }

  render () {
    return (
      <ScrollView style={{ backgroundColor: '#C2B280' }}>
        <View style={{ alignItems: 'center' }}>
          <Text style={{ fontSize: 30, fontWeight: 'bold' }}>Park Visitors:</Text>
        </View>
        <ListView
          dataSource={this.state.dataSource}
          renderRow={this._renderItem.bind(this)}
          enableEmptySections={true}
        />

        <Button
          title='Check In A Pet!'
          buttonStyle={{ marginTop: 20, marginBottom: 50 }}
          onPress={() => this.props.navigation.navigate('MyPetList')}
        />
      </ScrollView>
    )
  }
  _renderItem (item) {
    return <ListItem item={item} />
  }
}

export default DogListScreen

/*
        <List>
          {BigParkUsers.map((user, index) => (
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
          {me &&
          me.pet &&
          me.pet.picture && (
            <ListItem
              roundAvatar
              style={{ color: '#C2B280' }}
              avatar={{ uri: me.pet.picture }}
              title={me.pet.name}
              rightTitle={me.pet.gender}
              subtitle={`${me.name.first} ${me.name.last}`}
              onPress={() => {}}
            />
          )}
        </List>*/
