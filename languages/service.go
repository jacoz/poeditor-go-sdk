package languages

import (
	"errors"
	"github.com/jacoz/poeditor-go-sdk/utils"
	_ "net/http"
)

const (
	listPath = "/languages/list"
)

type Service struct {
	client *utils.HttpClient
}

func NewLanguagesService(client *utils.HttpClient) *Service {
	return &Service{client}
}

func (s *Service) List(req ListRequest) (*ListResponse, error) {
	res, err := s.client.NewPostRequest(listPath, req)
	if err != nil {
		return nil, err
	}

	if res.Response.Code != utils.ValidResponseCode {
		return nil, errors.New(res.Response.Message)
	}

	var r apiResponseLanguage
	err = utils.ConvertResultToDto(res.Result, &r)
	if err != nil {
		return nil, err
	}

	listResponse := ListResponse(r.Languages)

	return &listResponse, nil
}
