
export const goBackTo=(routeName, navigation) =>{
    console.log(`Trying to gobackto ${routeName}`)
    console.log(navigation)
    navigation.pop()
    while(navigation.state.routeName!==routeName){
        console.log(navigation)
        navigation.pop()
    }
} 