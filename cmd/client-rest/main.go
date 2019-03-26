package main

import (
	"../../pkg/api/v1"
	//fbitsdk "../../pkg/service-fitbit/fitbit-v1"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
	"../../pkg/aux"
	"strings"
	"time"
)

var svrAddress *string

func main() {
	// get configuration
	svrAddress = flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
	flag.Parse()

	demoRefreshToken()
}

func demoRefreshToken() {
	tmpToken, err := aux.ReadTokenFromFile("user_token")

	if err != nil {
		fmt.Println(err)
		return
	}
	reqRefresh := v1.RefreshTokenRequest{
		Api:         "v1",
		Item:        "REFRESH TOKEN",
		FitbitToken: tmpToken,
	}

	log.Printf("==== Rrfresh Token Demo ====\n")

	buf, _ := json.Marshal(reqRefresh)

	s := string(buf[:])

	var body string
	resp, err := http.Post(*svrAddress+"/v1/fitbit/refreshtoken", "application/json", strings.NewReader(s))

	if err != nil {
		log.Fatalf("failed to refresh token : %v", err)
		return
	}

	// Parsing and save the new refrsh token
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed to get new refresh token response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	//log.Printf("New Refresh Token response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	var rawToken map[string]interface{}

	if err := json.Unmarshal([]byte(body), &rawToken); err != nil {
		log.Printf("post fail error %v", err)
	}

	if rawToken["error"] != nil {
		fmt.Println(rawToken["error"].(string))
		return
	}

	log.Printf("api = %v and item %v\n", rawToken["api"].(string), rawToken["item"].(string))

	// get and save new refresh token
	val := rawToken["fitbitTokenNew"].(map[string]interface{})
	data, err := json.Marshal(val)

	aux.SaveTokenToFile("tt_token", data)

	// checking ...
	mx := make(map[string]interface{})
	err = json.Unmarshal(data, &mx)
	log.Printf("Expiry = %v\n", mx["expiry"].(string))
	log.Printf("RefreshToken = %v\n", mx["refreshToken"].(string))
	log.Printf("UserId = %v\n", mx["userId"])
	log.Printf("*************************************************************\n")
	log.Printf("****************** Remove the user_token.json ***************\n")
	log.Printf("***** and re-name the tt_token.json the user_token.json *****\n")
	log.Printf("*************************************************************\n")
}

// Get Fitbit activities steps task
func readActivitiesSteps() {
	/*
		tmpToken, err := clientLoadTokenFromFile("user_token")
		if err != nil {
			fmt.Println("==== READ token file fail ===")
		}
		reqSteps := v1.ActivitiesStepRequest{
			Api:         apiVersion,
			Item:        "STEPS",
			FitbitToken: tmpToken,
		}
	*/

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)
	var body string

	resp, err := http.Post(*svrAddress+"/v1/activities/step/all", "application/json", strings.NewReader(fmt.Sprintf(`
		{	"api":"v1",
		"toDo": {
			"title":"title (%s)",
			"description":"description (%s)",
			"reminder":"%s"
			}
		}
		`, pfx, pfx, pfx)))

	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}

	//
	//resp, err := http.Get(*svrAddress + "/v1/activities/step/all")
	//resp, err := http.Get(fmt.Sprintf("%s%s", *svrAddress, "/v1/activities/step/all"))

	//if err != nil {
	//log.Fatalf("failed to call Steps method: %v", err)
	//}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read ActivitiesStepAllRequest response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("ActivitiesStepAllRequest response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

}

// Get Fitbit activities distance task
func readActivitiesDistance() {

}
