package v1

import (
	//"../../api/v1"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"
)

// Read all activities Heart Rate task
func (s *toDoServiceServer) ReadActivitiesHeartRate(ctx context.Context, req *v1.ActivitiesHeartRateRequest) (*v1.ActivitiesHeartRateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	userActivitiesRequest(ctx, 4, req.FitbitToken)

	//
	hrList := []*v1.ActivitiesHeartRate{}

	hr := new(v1.ActivitiesHeartRate)

	hr.DateTime = "2019-03-08 17:53:12"
	hr.Bead = "1352"
	hrList = append(hrList, hr)

	return &v1.ActivitiesHeartRateResponse{
		Api:                 apiVersion,
		Item:                "HEART RATE",
		ActivitiesHeartRate: hrList,
	}, nil
}

func heartRateParser(buf bytes.Buffer) (string, error) {
	fmt.Println("== HA HA ==")

	xbyte := []byte{}
	xbyte = buf.Bytes()

	var heart map[string]interface{}

	if err := json.Unmarshal(xbyte, &heart); err != nil {
		fmt.Println("== D ERROR ==")
		return "", err
	}

	// start ...
	xyz := heart["activities-heart"]

	mx := xyz.([]interface{})
	fmt.Println("----------------")
	fmt.Println(mx)
	fmt.Println("--------*******--------")
	/*
		mm := mx.(map[string]interface{})
		fmt.Println(mm["dateTime"])
		fmt.Println(mm["value"])
	*/
	for k, v := range mx {
		fmt.Println(k)
		mm := v.(map[string]interface{})
		fmt.Println(mm["dateTime"])
		fmt.Println(mm["value"])
		fmt.Println()

		val := mm["value"].(map[string]interface{})
		fmt.Println(val)
		fmt.Println()
		//for vz1, vz2 := range val {
		//	fmt.Println(vz1["customHeartRateZones"])
		//	fmt.Println(vz2["heartRateZones"])
		//}
		fmt.Println()
		//hrZones := val["heartRateZones"].(map[string]interface{})
		//fmt.Println(hrZones)
		//fmt.Println()

	}

	fmt.Println("--------*******--------")

	result := "not error"
	fmt.Println(result)
	return result, nil
}

// Heart Rate
//GET https://api.fitbit.com/1/user/[user-id]/activities/heart/date/[date]/[period].json
//GET https://api.fitbit.com/1/user/[user-id]/activities/heart/date/[base-date]/[end-date].json

//Example Request
//GET https://api.fitbit.com/1/user/-/activities/heart/date/today/1d.json
//Example Response
/*
{
    "activities-heart": [
        {
            "dateTime": "2015-08-04",
            "value": {
                "customHeartRateZones": [],
                "heartRateZones": [
                    {
                        "caloriesOut": 740.15264,
                        "max": 94,
                        "min": 30,
                        "minutes": 593,
                        "name": "Out of Range"
                    },
                    {
                        "caloriesOut": 249.66204,
                        "max": 132,
                        "min": 94,
                        "minutes": 46,
                        "name": "Fat Burn"
                    },
                    {
                        "caloriesOut": 0,
                        "max": 160,
                        "min": 132,
                        "minutes": 0,
                        "name": "Cardio"
                    },
                    {
                        "caloriesOut": 0,
                        "max": 220,
                        "min": 160,
                        "minutes": 0,
                        "name": "Peak"
                    }
                ],
                "restingHeartRate": 68
            }
        }
    ]
}
*/

//user-id	The encoded ID of the user. Use "-" (dash) for current logged-in user.
//base-date	The range start date, in the format yyyy-MM-dd or today.
//end-date	The end date of the range.
//date	The end date of the period specified in the format yyyy-MM-dd or today.
//period	The range for which data will be returned. Options are 1d, 7d, 30d, 1w, 1m.

//GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/[end-date]/[detail-level].json
//GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/[end-date]/[detail-level]/time/[start-time]/[end-time].json
//GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/1d/[detail-level].json`
//GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/1d/[detail-level]/time/[start-time]/[end-time].json
//date	The date, in the format yyyy-MM-dd or today.
//detail-level	Number of data points to include. Either 1sec or 1min. Optional.
//start-time	The start of the period, in the format HH:mm. Optional.
//end-time	The end of the period, in the format HH:mm. Optional.
//Example Request
//GET https://api.fitbit.com/1/user/-/activities/heart/date/today/1d/1sec/time/00:00/0

//Example Response
/*
{
    "activities-heart": [
        {
            "customHeartRateZones": [],
            "dateTime": "today",
            "heartRateZones": [
                {
                    "caloriesOut": 2.3246,
                    "max": 94,
                    "min": 30,
                    "minutes": 2,
                    "name": "Out of Range"
                },
                {
                    "caloriesOut": 0,
                    "max": 132,
                    "min": 94,
                    "minutes": 0,
                    "name": "Fat Burn"
                },
                {
                    "caloriesOut": 0,
                    "max": 160,
                    "min": 132,
                    "minutes": 0,
                    "name": "Cardio"
                },
                {
                    "caloriesOut": 0,
                    "max": 220,
                    "min": 160,
                    "minutes": 0,
                    "name": "Peak"
                }
            ],
            "value": "64.2"
        }
    ],
    "activities-heart-intraday": {
        "dataset": [
            {
                "time": "00:00:00",
                "value": 64
            },
            {
                "time": "00:00:10",
                "value": 63
            },
            {
                "time": "00:00:20",
                "value": 64
            },
            {
                "time": "00:00:30",
                "value": 65
            },
            {
                "time": "00:00:45",
                "value": 65
            }
        ],
        "datasetInterval": 1,
        "datasetType": "second"
    }
}
*/

type UserTatal_tx struct {
	user UserProfile_tx
}

type UserProfile_tx struct {
	user string `json:"user"`
}

type userContent_tx struct {
	aboutMe                string `json:"aboutMe"`                //":<value>,
	avatar                 string `json:"avatar"`                 //:<value>,
	avatar150              string `json:"avatar150"`              ////":<value>,
	avatar640              string `json:"avatar640"`              //":<value>,
	city                   string `json:"city"`                   //":<value>,
	clockTimeDisplayFormat string `json:"clockTimeDisplayFormat"` //":<12hour|24hour>,
	country                string `json:"country"`                //":<value>,
	dateOfBirth            string `json:"dateOfBirth"`            //":<value>,
	displayName            string `json:"displayName"`            //":<value>,
	distanceUnit           string `json:"distanceUnit"`           //":<value>,
	encodedId              string `json:"encodedId"`              //":<value>,
	foodsLocale            string `json:"foodsLocale"`            //":<value>,
	fullName               string `json:"fullName"`               //":<value>,
	gender                 string `json:"gender"`                 //":<FEMALE|MALE|NA>,
	glucoseUnit            string `json:"glucoseUnit"`            //":<value>,
	height                 string `json:"height"`                 //":<value>,
	heightUnit             string `json:"heightUnit"`             //":<value>,
	locale                 string `json:"locale"`                 //":<value>,
	memberSince            string `json:"memberSince"`            //":<value>,
	offsetFromUTCMillis    string `json:"offsetFromUTCMillis"`    //":<value>,
	startDayOfWeek         string `json:"startDayOfWeek"`         //":<value>,
	state                  string `json:"state"`                  //":<value>,
	strideLengthRunning    string `json:"strideLengthRunning"`    //":<value>,
	strideLengthWalking    string `json:"strideLengthWalking"`    //":<value>,
	timezone               string `json:"timezone"`               //":<value>,
	waterUnit              string `json:"waterUnit"`              //":<value>,
	weight                 string `json:"weight"`                 //":<value>,
	weightUnit             string `json:"weightUnit"`             //":<value>
}

type userContent struct {
	aboutMe                string //":<value>,
	avatar                 string //:<value>,
	avatar150              string ////":<value>,
	avatar640              string //":<value>,
	city                   string //":<value>,
	clockTimeDisplayFormat string //":<12hour|24hour>,
	country                string //":<value>,
	dateOfBirth            string //":<value>,
	displayName            string //":<value>,
	distanceUnit           string //":<value>,
	encodedId              string //":<value>,
	foodsLocale            string //":<value>,
	fullName               string //":<value>,
	gender                 string //":<FEMALE|MALE|NA>,
	glucoseUnit            string //":<value>,
	height                 string //":<value>,
	heightUnit             string //":<value>,
	locale                 string //":<value>,
	memberSince            string //":<value>,
	offsetFromUTCMillis    string //":<value>,
	startDayOfWeek         string //":<value>,
	state                  string //":<value>,
	strideLengthRunning    string //":<value>,
	strideLengthWalking    string //":<value>,
	timezone               string //":<value>,
	waterUnit              string //":<value>,
	weight                 string //":<value>,
	weightUnit             string //":<value>
}
