package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	//	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"../../../pkg/api/v1"
	//"../../api/v1"
	"go-heroku/fitbit-demo/web-server/pkg/api/v1"

	//"io"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"strings"

	//fbitsdk "../../service-fitbit/fitbit-v1"
	//"encoding/json"
	"golang.org/x/oauth2"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
}

func NewToDoServiceServer() v1.ToDoServiceServer {
	return &toDoServiceServer{}
}

// checkAPI checks if the API version
func (s *toDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'",
				apiVersion, api)
		}
	}
	return nil
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	fmt.Println("=== Create ===")
	// check if the API version
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  100,
	}, nil
}

// Read all todo tasks
func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	//
	var reminder time.Time
	list := []*v1.ToDo{}

	td := new(v1.ToDo)
	//td.Reminder, err = ptypes.TimestampProto(reminder)
	td.Reminder, _ = ptypes.TimestampProto(reminder)
	list = append(list, td)

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}

// Read todo task
func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {

	fmt.Println("=== Read ===")
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get or set(test) ToDo data
	var td v1.ToDo
	var reminder time.Time

	//
	//td.Reminder, err = ptypes.TimestampProto(reminder)
	td.Reminder, _ = ptypes.TimestampProto(reminder)

	return &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: &td,
	}, nil

}

/*
func (s *toDoServiceServer) SendToken(ctx context.Context, req *v1.FitbitToken) (*empty.Empty, error) {

	fmt.Println("=== Send User Toke Test ===")
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	out := new(empty.Empty)

	//layout := "2019-03-10 02:34:07.68904 +0800 CST"

	//t, err := time.Parse(layout, req.Expiry)
	//fmt.Println(t)
	//loc := &time.Location{"America/New_York": "UTC"}
	//fmt.Println(loc)

	//newYork, _ := time.LoadLocation("America/New_York")
	//fmt.Println(newYork)

	//layout := "2019-03-10T02:34:07.689Z"
	//str := "2019-03-10T02:34:07.689Z08"
	//layout := "2006-01-02T15:04:05.000Z08"
	//str := "2017-11-12T11:45:26.371Z"
	//t, err := time.Parse(layout, str)

	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)

	fmt.Println(req.AccessToken)
	fmt.Println(req.TokenType)
	fmt.Println(req.RefreshToken)
	fmt.Println(req.Expiry)
	fmt.Println(req.UserId)

	utoken := fbitsdk.HFToken{
		AccessToken:  req.AccessToken,
		TokenType:    req.TokenType,
		RefreshToken: req.RefreshToken,
		Expiry:       req.Expiry, //t,
		UserId:       req.UserId,
	}
	fmt.Println(utoken.Expiry)
	fmt.Println(utoken.UserId)

	fmt.Println("=== Send User Toke Test END ===")

	return out, nil
}
*/
var (
	oauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8020/cb",
		Scopes:      []string{"activity", "heartrate", "location", "nutrition", "profile", "settings", "sleep", "social", "weight"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.fitbit.com/oauth2/authorize",
			TokenURL: "https://api.fitbit.com/oauth2/token",
		},
		ClientID:     "22D6FQ",
		ClientSecret: "be9c1fb74ca0d6b8c93deb35ba305093",
	}
)
