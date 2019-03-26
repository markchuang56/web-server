package v1

import (
	//"../../api/v1"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"
	//"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// Read all activities Distance task
func (s *toDoServiceServer) ReadActivitiesDistance(ctx context.Context, req *v1.ActivitiesDistanceRequest) (*v1.ActivitiesDistanceResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	rawDistance, err := userActivitiesRequest(ctx, 3, req.FitbitToken)
	if err != nil {
		return nil, err
	}
	return activitiesDistanceParser(rawDistance)

}

// Parser activities Distance
func activitiesDistanceParser(buf interface{}) (*v1.ActivitiesDistanceResponse, error) {

	xbyte, _ := buf.([]byte)
	//fmt.Println(xbyte)

	var distance map[string]interface{}

	if err := json.Unmarshal(xbyte, &distance); err != nil {
		fmt.Println("== D ERROR ==")
		return nil, err
	}

	fmt.Println("==== ACT DISTANCE ====")
	xyz := distance["activities-distance"]
	mx := xyz.([]interface{})

	distanceList := []*v1.ActivitiesDistance{}

	for _, v := range mx {
		mm := v.(map[string]interface{})
		// value, _ := v.(float64)
		dis := new(v1.ActivitiesDistance)
		dis.DateTime = mm["dateTime"].(string)
		dis.Distance = mm["value"].(string)
		distanceList = append(distanceList, dis)
	}

	return &v1.ActivitiesDistanceResponse{
		Api:                apiVersion,
		Item:               "DISTANCE",
		ActivitiesDistance: distanceList,
	}, nil
}

// Read all activities Step task
func (s *toDoServiceServer) ReadActivitiesStep(ctx context.Context, req *v1.ActivitiesStepRequest) (*v1.ActivitiesStepResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	rawStep, err := userActivitiesRequest(ctx, 2, req.FitbitToken)
	if err != nil {
		return nil, err
	}

	return activitiesStepParser(rawStep)
}

// Parser activities Steps
func activitiesStepParser(buf interface{}) (*v1.ActivitiesStepResponse, error) {
	fmt.Println("== HA HA ==")
	xbyte, _ := buf.([]byte)
	var xad map[string]interface{}

	if err := json.Unmarshal(xbyte, &xad); err != nil {
		fmt.Println("== D ERROR ==")
		return nil, err
	}

	stepsList := []*v1.ActivitiesStep{}

	fmt.Println("==== ACT STEP ====")
	xyz := xad["activities-steps"]
	mx := xyz.([]interface{})
	for _, v := range mx {
		mm := v.(map[string]interface{})
		step := new(v1.ActivitiesStep)
		step.DateTime = mm["dateTime"].(string)
		step.Steps = mm["value"].(string)
		//fmt.Println(step.DateTime)
		//fmt.Println(mm["dateTime"])

		stepsList = append(stepsList, step)
		//fmt.Println(stepsList)
	}

	return &v1.ActivitiesStepResponse{
		Api:             apiVersion,
		Item:            "STEPS",
		ActivitiesSteps: stepsList,
	}, nil
}

//GET https://api.fitbit.com/1/user/[user-id]/activities/date/[date].json
//user-id	The encoded ID of the user. Use "-" (dash) for current logged-in user.
//date	The date in the format yyyy-MM-dd
//Request Headers
//Accept-Locale	optional	The locale to use for response values.
//Accept-Language	optional	The measurement unit system to use for response values.
//Example Response
/*
{
    "activities":[
        {
            "activityId":51007,
            "activityParentId":90019,
            "calories":230,
            "description":"7mph",
            "distance":2.04,
            "duration":1097053,
            "hasStartTime":true,
            "isFavorite":true,
            "logId":1154701,
            "name":"Treadmill, 0% Incline",
            "startTime":"00:25",
            "steps":3783
        }
    ],
    "goals":{
        "caloriesOut":2826,
        "distance":8.05,
        "floors":150,
        "steps":10000
     },
    "summary":{
        "activityCalories":230,
        "caloriesBMR":1913,
        "caloriesOut":2143,
        "distances":[
            {"activity":"tracker", "distance":1.32},
            {"activity":"loggedActivities", "distance":0},
            {"activity":"total","distance":1.32},
            {"activity":"veryActive", "distance":0.51},
            {"activity":"moderatelyActive", "distance":0.51},
            {"activity":"lightlyActive", "distance":0.51},
            {"activity":"sedentaryActive", "distance":0.51},
            {"activity":"Treadmill, 0% Incline", "distance":3.28}
        ],
        "elevation":48.77,
        "fairlyActiveMinutes":0,
        "floors":16,
        "lightlyActiveMinutes":0,
        "marginalCalories":200,
        "sedentaryMinutes":1166,
        "steps":0,
        "veryActiveMinutes":0
    }
}
*/

//Activity Time Series
//Get Activity Time Series
//The Get Activity Time Series endpoint returns time series data in the specified range for a given resource in the format requested using units in the unit system that corresponds to the Accept-Language header provided.

//Considerations
//Even if you provide earlier dates in the request, the response will retrieve only data since the user's join date or the first log entry date for the requested collection.

//The activities/tracker/... resource represents the daily activity values logged by the tracker device only, excluding manual activity log entries.

//The activities/tracker/calories resource does not include the Estimated Energy Requirement for calorie estimation (EER) calculations for any dates even if they are turned on for the user's profile and use BMR level instead.

//The activities collection is maintained as a backwards compatible resource urls (e.g. activities/log/calories).

//Elevation time series (/elevation, /floors) are only available for users with compatible trackers.

//Resource URL
//There are two acceptable formats for retrieving activity time series data:

//GET /1/user/[user-id]/[resource-path]/date/[date]/[period].json
//GET /1/user/[user-id]/[resource-path]/date/[base-date]/[end-date].json
//user-id	The encoded ID of the user. Use "-" (dash) for current logged-in user.
//resource-path	The resource path; see options in the "Resource Path Options" section below.
//base-date	The range start date, in the format yyyy-MM-dd or today.
//end-date	The end date of the range.
//date	The end date of the period specified in the format yyyy-MM-dd or today.
//period	The range for which data will be returned. Options are 1d, 7d, 30d, 1w, 1m, 3m, 6m, 1y

//Resource Path Options
//ACTIVITY

//activities/calories
//activities/caloriesBMR
//activities/steps
//activities/distance
//activities/floors
//activities/elevation
//activities/minutesSedentary
//activities/minutesLightlyActive
//activities/minutesFairlyActive
//activities/minutesVeryActive
//activities/activityCalories
