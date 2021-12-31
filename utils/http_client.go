package utils

import (
	"encoding/json"
	"fmt"
	"github.com/jacoz/poeditor-go-sdk/model"
	"net/http"
	"net/url"
	"strings"
)

const (
	requestApiTokenKey = "api_token"
)

type HttpClient struct {
	baseUrl string
	apiKey  string
}

func NewHttpClient(baseUrl string, apiKey string) *HttpClient {
	return &HttpClient{baseUrl, apiKey}
}

func (c *HttpClient) NewPostRequest(path string, body interface{}) (*model.ApiResponse, error) {
	apiUrl := c.baseUrl + path

	data, err := c.getPostData(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, data)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var res model.ApiResponse
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *HttpClient) getPostData(body interface{}) (*strings.Reader, error) {
	data := make(map[string]interface{})

	if body != nil {
		jsonStr, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonStr, &data)
		if err != nil {
			return nil, err
		}
	}

	data[requestApiTokenKey] = c.apiKey

	vals := url.Values{}
	for k, v := range data {
		vals.Add(k, fmt.Sprintf("%v", v))
	}

	return strings.NewReader(vals.Encode()), nil
}
