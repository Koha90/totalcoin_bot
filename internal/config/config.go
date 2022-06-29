package config

import (
	"encoding/json"
	"os"
)

type UserProfile struct {
	Username    string
	PhoneNumber string
	Password    string
}

func loadUserProfileJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func GetUserProfile() {

}
