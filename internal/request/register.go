package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateUserResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
}

func (c *Client) Register(ctx context.Context) ([]CreateUserResponse, error) {
	var result []CreateUserResponse

	for u, p := range Users {
		postBody, _ := json.Marshal(map[string]string{
			"userName": u,
			"password": p,
		})

		requestBody := bytes.NewBuffer(postBody)

		resp, err := register(requestBody)
		if err != nil {
			return nil, err
		}

		result = append(result, CreateUserResponse{
			ID:       resp.ID,
			UserName: resp.UserName,
		})
	}

	return result, nil
}

func register(requestBody io.Reader) (*CreateUserResponse, error) {
	url := apiPath + registerRoute

	req, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := makeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	result := CreateUserResponse{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return &result, nil
}
