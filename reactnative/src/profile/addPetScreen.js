import React, { Component } from 'react'
import { View, Keyboard, Text, Image, Dimensions, TextInput, ScrollView, StyleSheet} from 'react-native'
import {HomeButton} from '../shared/ui'
import { Tile, List, ListItem, Button, Input } from 'react-native-elements';
import ImagePicker from 'react-native-image-picker'


const dim = Dimensions.get('window')

class AddPetScreen extends Component{
    state={
        name:'',
        breed:'',
        sex:'',
        picture:'',
    }
    

    selectPhotoTapped() {
        console.log('Select photo tapped')
        const options = {
          quality: 1.0,
          maxWidth: 500,
          maxHeight: 500,
          storageOptions: {
            skipBackup: true
          }
        };
    
        //ImagePicker.showImagePicker(options, (response) => {
        ImagePicker.launchImageLibrary(options, (response)  => {
            console.log('Response = ', response);
    
          if (response.didCancel) {
            console.log('User cancelled photo picker');
          }
          else if (response.error) {
            console.log('ImagePicker Error: ', response.error);
          }
          else if (response.customButton) {
            console.log('User tapped custom button: ', response.customButton);
          }
          else {
            let source = { uri: response.uri };
    
            // You can also display the image using data:
            // let source = { uri: 'data:image/jpeg;base64,' + response.data };
    
            this.setState({
              picture: response.uri
            });
          }
        });
      }


    render(){
        return(
            <ScrollView style={{backgroundColor:'#C2B280'}}>
            <Tile
              imageSrc={{uri: this.state.picture!==''?this.state.picture:'https://www.babybedding.com/images/fabric/silver-gray-minky-fabric_smaller.jpg'}}
              featured
              title={this.state.picture!==''?'':'Select Photo'}
              caption=''
              onPress={this.selectPhotoTapped.bind(this)}
            />
            <List>
                <Text>Name:</Text>
                <TextInput
                    style={{marginTop:-10}}
                    label='Name'
                    placeholder='Lucky'
                    onChangeText={(name) => this.setState({name})}
                    value={this.state.name}
                />
                <Text>Breed:</Text>
                <TextInput
                    style={{marginTop:-10}}
                    label='Breed'
                    placeholder='Chihuahua'
                    onChangeText={(breed) => this.setState({breed})}
                    value={this.state.breed}
                />
                <Text>Sex:</Text>
                <TextInput
                    style={{marginTop:-10}}
                    label='Sex'
                    placeholder='Boy'
                    onChangeText={(sex) => this.setState({sex})}
                    value={this.state.sex}
                />
                <Button
                title="Add Pet"
                buttonStyle={{ marginTop: 0, marginBottom:50 }}
                onPress={()=>{this.props.navigation.goBack()}}
                />
            </List>
          </ScrollView>
        );
      }
    }
    
export default AddPetScreen