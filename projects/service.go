package projects

import (
	"errors"
	"github.com/jacoz/poeditor-go-sdk/utils"
	_ "net/http"
)

const (
	listPath   = "/projects/list"
	viewPath   = "/projects/view"
	exportPath = "/projects/export"
)

type Service struct {
	client *utils.HttpClient
}

func NewProjectsService(client *utils.HttpClient) *Service {
	return &Service{client}
}

func (s *Service) List() (*ListResponse, error) {
	res, err := s.client.NewPostRequest(listPath, nil)
	if err != nil {
		return nil, err
	}

	if res.Response.Code != utils.ValidResponseCode {
		return nil, errors.New(res.Response.Message)
	}

	var r apiResponseList
	err = utils.ConvertResultToDto(res.Result, &r)
	if err != nil {
		return nil, err
	}

	listResponse := ListResponse(r.Projects)

	return &listResponse, nil
}

func (s *Service) View(req ViewRequest) (*ProjectDetails, error) {
	res, err := s.client.NewPostRequest(viewPath, req)
	if err != nil {
		return nil, err
	}

	if res.Response.Code != utils.ValidResponseCode {
		return nil, errors.New(res.Response.Message)
	}

	var r apiResponseView
	err = utils.ConvertResultToDto(res.Result, &r)
	if err != nil {
		return nil, err
	}

	return &r.Project, nil
}

func (s *Service) Export(req ExportRequest) (*ExportResponse, error) {
	res, err := s.client.NewPostRequest(exportPath, req)
	if err != nil {
		return nil, err
	}

	if res.Response.Code != utils.ValidResponseCode {
		return nil, errors.New(res.Response.Message)
	}

	var e ExportResponse
	err = utils.ConvertResultToDto(res.Result, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
