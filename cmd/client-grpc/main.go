package main

import (
	"context"
	"flag"
	"log"
	"time"

	//"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	api "../../pkg/api/v1"
	aux "../../pkg/aux"
	//fbitsdk "../../pkg/service-fitbit/fitbit-v1"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var gRpcClient v1.ToDoServiceClient

//var reqToken v1.FitbitToken

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	//client := v1.NewToDoServiceClient(conn)
	gRpcClient = v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// -- DEMO -- Get fitbit steps
	readActivitiesSteps(ctx)

	// -- DEMO -- Get fitbit distance
	readActivitiesDistance(ctx)

	// -- DEMO -- Get fitbit heart rate
	//readActivitiesHeartRate(ctx)

	// get fitbit user profile

	// Send User Token Test ..
	//demoRefreshToken(ctx)

}

// Get Fitbit activities steps task
func readActivitiesSteps(ctx context.Context) {
	tmpToken, err := aux.ReadTokenFromFile("user_token")
	if err != nil {
		fmt.Println("==== READ token file fail ===")
	}
	reqSteps := v1.ActivitiesStepRequest{
		Api:         apiVersion,
		Item:        "STEPS",
		FitbitToken: tmpToken,
	}

	resSteps, err := gRpcClient.ReadActivitiesStep(ctx, &reqSteps)
	if err != nil {
		log.Fatalf("Steps failed: %v", err)
	}
	//log.Printf("Steps result: <%+v>\n\n", resSteps)
	log.Printf("Steps result: <%+v>\n\n", resSteps.Api)
	log.Printf("Steps result: <%+v>\n\n", resSteps.Item)
	for k, v := range resSteps.ActivitiesSteps {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// Get Fitbit activities distance task
func readActivitiesDistance(ctx context.Context) {
	tmpToken, err := aux.ReadTokenFromFile("user_token")
	if err != nil {
		fmt.Println("==== READ token file fail ===")
	}
	reqDistance := v1.ActivitiesDistanceRequest{
		Api:         apiVersion,
		Item:        "DISTANCE",
		FitbitToken: tmpToken,
	}

	respDistance, err := gRpcClient.ReadActivitiesDistance(ctx, &reqDistance)
	if err != nil {
		log.Fatalf("Distance failed: %v", err)
	}
	//log.Printf("Distance result: <%+v>\n\n", respDistance.Api)
	log.Printf("Distance result: <%+v>\n\n", respDistance.Item)
	for k, v := range respDistance.ActivitiesDistance {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// Get Fitbit activities heart rate task
func readActivitiesHeartRate(ctx context.Context) {
	tmpToken, err := aux.ReadTokenFromFile("user_token")
	if err != nil {
		fmt.Println("==== READ token file fail ===")
	}
	reqBead := v1.ActivitiesHeartRateRequest{
		Api:         apiVersion,
		Item:        "HEART RATE",
		FitbitToken: tmpToken,
	}

	resSteps, err := gRpcClient.ReadActivitiesHeartRate(ctx, &reqBead)
	if err != nil {
		log.Fatalf("Heart Rate failed: %v", err)
	}
	log.Printf("Heart Rate result: <%+v>\n\n", resSteps)
}

// Get Fitbit manual refresh token task
func demoRefreshToken(ctx context.Context) {
	tmpToken, err := aux.ReadTokenFromFile("user_token")
	//tmpToken, err := aux.ReadTokenFromFile("tt")
	if err != nil {
		fmt.Println(err)
	}
	reqRefresh := v1.RefreshTokenRequest{
		Api:         apiVersion,
		Item:        "REFRESH TOKEN",
		FitbitToken: tmpToken,
	}

	newToken, err := gRpcClient.ReadNewRefreshToken(ctx, &reqRefresh)
	if err != nil {
		log.Fatalf("NEW TOKEN failed: %v", err)
	}
	log.Printf("NEW TOKEN result: <%+v>\n\n", newToken)
}
