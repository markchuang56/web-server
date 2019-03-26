package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
// Get User Profile
func SdkGetProfile(uid string) {
	fmt.Println("===  ... ===")
	_, err := userActivitiesRequest(1, uid)
	if err != nil {
		fmt.Println("=== error ===")
	}
}
*/

// Normal Parser User Profile
func parserUserProfile(buf bytes.Buffer) (string, error) {

	xbytes := []byte{}
	xbytes = buf.Bytes()

	var pf map[string]interface{}

	if err := json.Unmarshal(xbytes, &pf); err != nil {
		fmt.Println("== D ERROR ==")
		//panic(err)
		return "", err
	}

	user := pf["user"]
	fmt.Println("=== user ===")
	fmt.Println(user)
	mx := user.(map[string]interface{})
	fmt.Println(mx)
	var result string

	for k, v := range mx {
		fmt.Println("======")
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println()
		switch k {
		case "age":
			//value, _ := v.(float64)
			fmt.Println("AGE", v)
			break

		case "displayName":
			fmt.Println("dispalyName", v)
			break

		case "encodedId":
			fmt.Println("User ID", v)
			result, _ = v.(string)
			break

		default:
			break
		}
	}
	fmt.Println(result)
	return result, nil
}

//
type UserProfile struct {
	user string
	/*
		"user": {
		        "aboutMe":<value>,
		        "avatar":<value>,
		        "avatar150":<value>,
		        "avatar640":<value>,
		        "city":<value>,
		        "clockTimeDisplayFormat":<12hour|24hour>,
		        "country":<value>,
		        "dateOfBirth":<value>,
		        "displayName":<value>,
		        "distanceUnit":<value>,
		        "encodedId":<value>,
		        "foodsLocale":<value>,
		        "fullName":<value>,
		        "gender":<FEMALE|MALE|NA>,
		        "glucoseUnit":<value>,
		        "height":<value>,
		        "heightUnit":<value>,
		        "locale":<value>,
		        "memberSince":<value>,
		        "offsetFromUTCMillis":<value>,
		        "startDayOfWeek":<value>,
		        "state":<value>,
		        "strideLengthRunning":<value>,
		        "strideLengthWalking":<value>,
		        "timezone":<value>,
		        "waterUnit":<value>,
		        "weight":<value>,
		        "weightUnit":<value>
		    }
	*/
}

//GET https://api.fitbit.com/1/user/[user-id]/profile.json
//user-id	The encoded ID of the user. Use "-" (dash) for current logged-in user.

/*
// User Profile parser for USER ID
func sdkParserUserId(buf bytes.Buffer) (string, error) {
	fmt.Println("== HA HA ==")

	xbyte := []byte{}
	xbyte = buf.Bytes()

	var xad map[string]interface{}

	if err := json.Unmarshal(xbyte, &xad); err != nil {
		fmt.Println("== D ERROR ==")
		//panic(err)
		return "", err
	}

	xyz := xad["user"]
	mx := xyz.(map[string]interface{})
	var result string

	for k, v := range mx {
		switch k {
		case "age":
			//value, _ := v.(float64)
			fmt.Println("AGE", v)
			break

		case "displayName":
			fmt.Println("dispalyName", v)
			break

		case "encodedId":
			fmt.Println("User ID", v)
			result, _ = v.(string)
			break

		default:
			break
		}
	}
	fmt.Println(result)
	return result, nil
}
*/
