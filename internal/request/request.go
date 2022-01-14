package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiPath = "http://localhost:8080/v1/"
	wsPath  = "ws://127.0.0.1:8080/v1/ws"

	registerRoute = "user"
	loginRoute    = "user/login"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

var Users = map[string]string{
	"userA111": "userA1234",
	"userA222": "userA1234",
	"userA333": "userA1234",
	"userA444": "userA1234",
	"userA555": "userA1234",
	"userA666": "userA1234",
}

func makeHTTPRequest(req *http.Request) ([]byte, error) {
	httpClient := http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error with %d code", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
