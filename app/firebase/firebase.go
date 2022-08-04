package firebase

import firebase "firebase.google.com/go"

func declare() {

	config := &firebase.Config{
		DatabaseURL: "https://database-name.firebaseio.com",
	}
}
