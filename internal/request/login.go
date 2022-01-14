package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginResponse struct {
	URL string `json:"url"`
}

func (c *Client) Login(ctx context.Context) ([]LoginResponse, error) {
	var result []LoginResponse

	for u, p := range Users {
		postBody, _ := json.Marshal(map[string]string{
			"userName": u,
			"password": p,
		})

		requestBody := bytes.NewBuffer(postBody)

		resp, err := login(requestBody)
		if err != nil {
			return nil, err
		}

		result = append(result, LoginResponse{
			URL: resp.URL,
		})
	}

	return result, nil
}

func login(requestBody io.Reader) (*LoginResponse, error) {
	url := apiPath + loginRoute

	req, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := makeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	result := LoginResponse{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	// todo clean
	fmt.Println(result.URL)

	return &result, nil
}
