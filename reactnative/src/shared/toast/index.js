import { Toast } from 'native-base'

module.exports = {
  Message: (message) => {
    Toast.show({
      text: message,
      position: 'top',
      buttonText: 'Dismiss',
      type: 'success',
      duration: 5000
    })
  },
  Error: (message) => {
    Toast.show({
      text: message,
      position: 'top',
      buttonText: 'Dismiss',
      type: 'danger',
      duration: 5000
    })
  }
}
