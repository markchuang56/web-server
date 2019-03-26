package aux

import (
	//"../../pkg/api/v1"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"

	"encoding/json"
	//"fmt"
	"io/ioutil"
)

// save user access token to loacl storage
func SaveTokenToFile(title string, data []byte) {
	filename := title + ".json"
	erx := ioutil.WriteFile(filename, data, 0644)
	check(erx)

	//erx = ioutil.WriteFile("user_token.json", data, 0644)
	//check(erx)
}

// check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadTokenFromFile(title string) (*v1.FitbitToken, error) {
	filename := title + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var fileToken v1.FitbitToken
	json.Unmarshal(body, &fileToken)
	return &fileToken, nil
}
