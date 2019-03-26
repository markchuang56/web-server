package v1

import (
	//"../../api/v1"
	"bytes"
	"context"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"
	//"fmt"
	"io"
	//"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (s *toDoServiceServer) ReadNewRefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	log.Printf("=== Server Read New Refresh Token ===\n")
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		log.Printf("=== RF and 0 ===\n")
		return nil, err
	}

	tokenAddress := "https://api.fitbit.com/oauth2/token"
	oldToken, _ := setTokenUserID(req.FitbitToken)
	now := time.Now()
	log.Printf("%v\n", now)
	nanoSeconds := now.UnixNano()
	log.Printf("%v\n", nanoSeconds)
	nanoSeconds = nanoSeconds + 28800*1000*1000*1000

	//token.Expiry = time.Unix(0, tmpNanoSecond)
	log.Printf("=== NEW Expire ===\n")
	log.Printf("%v\n", time.Unix(0, nanoSeconds))

	newExpiryString := strconv.FormatInt(int64(nanoSeconds), 10)

	log.Printf("%v", req.FitbitToken.UserId)
	var body io.Reader
	v := url.Values{
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{oldToken.RefreshToken},
	}

	body = strings.NewReader(v.Encode())

	reqRefreh, err := http.NewRequest(http.MethodPost, tokenAddress, body)

	if err != nil {
		return nil, err
	}

	xtype := "Basic Y2xpZW50X2lkOmNsaWVudCBzZWNyZXQ="
	reqRefreh.Header.Set("Authorization", xtype)

	reqRefreh.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	reqRefreh.WithContext(ctx)
	//

	log.Printf("===============\n")
	log.Printf("%v\n", req.FitbitToken.Expiry)

	// go to refresh token
	client := oauthConfig.Client(ctx, oldToken)
	resp, err := client.Do(reqRefreh)

	//fmt.Println(err)
	if err != nil {
		log.Printf("== post err == %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)

	// extend
	xbyte := []byte{}
	xbyte = buf.Bytes()

	//

	var xxxToken map[string]interface{}

	if err := json.Unmarshal(xbyte, &xxxToken); err != nil {
		return nil, err
	}
	log.Printf("==== difference between refresh token ====\n")

	//log.Printf("acces token : %s\n", xxxToken["access_token"])
	log.Printf("token type : %s\n", xxxToken["token_type"])
	log.Printf("refresh token : %s\n", xxxToken["refresh_token"])
	log.Printf("expires in : %v\n", xxxToken["expires_in"])
	log.Printf("scope : %s\n", xxxToken["scope"])
	log.Printf("user id : %s\n", xxxToken["user_id"])

	// string to nano second
	if tmpNanoSecond, err := strconv.ParseInt(req.FitbitToken.Expiry, 10, 64); err == nil {
		log.Printf("%T, %v\n", tmpNanoSecond, tmpNanoSecond)
		// int64 nano second to timestamp
		log.Printf("%v\n", time.Unix(0, tmpNanoSecond))
	}

	log.Printf("\n")

	newToken := v1.FitbitToken{
		Api:          apiVersion,
		AccessToken:  xxxToken["access_token"].(string),
		TokenType:    xxxToken["token_type"].(string),
		RefreshToken: xxxToken["refresh_token"].(string),
		Expiry:       newExpiryString,
		UserId:       xxxToken["user_id"].(string),
	}

	log.Printf("=== GOOD REFRESH ===\n")
	return &v1.RefreshTokenResponse{
		Api:            apiVersion,
		Item:           "Refresh Token",
		FitbitTokenNew: &newToken,
	}, nil
}
