package v1

import (
	"bytes"
	"context"
	//"fmt"
	"time"
	//"github.com/golang/protobuf/ptypes"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	//"../../../pkg/api/v1"
	//"../../api/v1"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"

	//"encoding/json"
	"golang.org/x/oauth2"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GET USER DATA
func userActivitiesRequest(ctx context.Context, cmd int, serveToken *v1.FitbitToken) (interface{}, error) {
	log.Printf("=== user Activities Request ===\n")

	log.Printf("---------\n")
	userDataSel = cmd

	userfitbitToken, _ = setTokenUserID(serveToken)

	switch cmd {
	case 0:
		userfitbitAddress = ProfileBegin + "-" + ProfileEnd
		break

	case 1:
		userfitbitAddress = ProfileBegin + userfitbitId + ProfileEnd
		break

	case 2:
		userfitbitAddress = ActivitiesBegin + userfitbitId + ActivitiesStepsEnd
		break

	case 3:
		userfitbitAddress = ActivitiesBegin + userfitbitId + ActivitiesDistanceEnd
		break

	case 4:
		userfitbitAddress = HardRateBegin + userfitbitId + HardRateEnd
		break

	case 5:
		userfitbitAddress = BodyFatBegin + userfitbitId + BodyFatEnd
		break

	case 6:
		userfitbitAddress = "https://api.fitbit.com/1.1/oauth2/introspect"
		break

	default:
		break

	}

	varResult, err := fitbitGetTask(ctx)
	log.Printf("%v\n")
	log.Printf("%v\n", err)
	if err != nil {
		log.Printf("=== need refresh token ===\n")
		return "fail", err
	}
	//fmt.Println(varResult)
	return varResult, nil
}

func fitbitGetTask(ctx context.Context) (interface{}, error) {
	log.Printf("==== FITBIT GET TASK ====\n")

	var body io.Reader
	v := url.Values{}
	body = strings.NewReader(v.Encode())

	req, err := http.NewRequest(http.MethodGet, userfitbitAddress, body)

	if err != nil {
		return "fail", err
	}

	xtype := userfitbitToken.TokenType + " " + userfitbitToken.AccessToken

	req.Header.Set("Authorization", xtype)

	req.WithContext(ctx)
	//
	//log.Printf("%v,", req)

	client := oauthConfig.Client(ctx, userfitbitToken)
	resp, err := client.Do(req)

	if err != nil {
		return "fail", err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)

	// extend
	xbyte := []byte{}
	xbyte = buf.Bytes()

	return xbyte, nil
}

var userfitbitAddress string
var userfitbitId string
var userfitbitToken *oauth2.Token
var userDataSel int

func setTokenUserID(serveToken *v1.FitbitToken) (*oauth2.Token, error) {
	var token oauth2.Token
	token.AccessToken = serveToken.AccessToken
	token.TokenType = serveToken.TokenType
	token.RefreshToken = serveToken.RefreshToken
	log.Printf("nano second %v", serveToken.Expiry)

	// string to nanosecon, to timestamp
	//token.Expiry = time.Unix(0, serveToken.Expiry)
	if tmpNanoSecond, err := strconv.ParseInt(serveToken.Expiry, 10, 64); err == nil {
		log.Printf("%T, %v\n", tmpNanoSecond, tmpNanoSecond)
		token.Expiry = time.Unix(0, tmpNanoSecond)
		log.Printf("Time Stampe = %v", token.Expiry)
	}

	userfitbitId = serveToken.UserId

	return &token, nil
}

const (
	TokenFileName = "token"
	ProfileBegin  = "https://api.fitbit.com/1/user/"
	ProfileEnd    = "/profile.json"

	SleepBegin = "https://api.fitbit.com/1.2/user/"
	//SleepEnd   = "/sleep/date/2018-12-03.json"
	SleepEnd = "/sleep/date/2019-02-03/2019-03-02.json"

	//GET https://api.fitbit.com/1.2/user/-/sleep/date/2017-04-02/2017-04-08.json

	ActivitiesBegin       = "https://api.fitbit.com/1/user/"
	ActivitiesStepsEnd    = "/activities/steps/date/today/1m.json"
	ActivitiesDistanceEnd = "/activities/distance/date/today/1m.json"

	HardRateBegin = "https://api.fitbit.com/1/user/"
	HardRateEnd   = "/activities/heart/date/today/1d.json"
	//HardRateEnd = "6Z29KN/activities/heart/date/today/1m.json"
	//HardRateEnd = "/activities/heart/date/today/1m.json"

	BodyFatBegin = "https://api.fitbit.com/1/user/"
	BodyFatEnd   = "/body/log/fat/date/2018-12-03.json"

	//distance
	// "https://api.fitbit.com/1/user/6Z29KN/profile.json"
	// "https://api.fitbit.com/1.2/user/6Z29KN/sleep/date/2018-12-03.json"
	// "https://api.fitbit.com/1/user/6Z29KN/activities/steps/date/today/1m.json"
	//"https://api.fitbit.com/1/user/6Z29KN/activities/heart/date/today/1d.json"
	//introspect := "https://api.fitbit.com/1.1/oauth2/introspect"
	//xbodyWeight := 	//"https://api.fitbit.com/1/user/6Z29KN/body/log/fat/date/2018-12-03.json"
)
