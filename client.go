package poeditor_go_sdk

import (
	"github.com/jacoz/poeditor-go-sdk/languages"
	"github.com/jacoz/poeditor-go-sdk/projects"
	"github.com/jacoz/poeditor-go-sdk/utils"
)

const baseApiUrl = "https://api.poeditor.com/v2"

type Client struct {
	Projects  *projects.Service
	Languages *languages.Service
}

func NewClient(apiKey string) *Client {
	httpClient := utils.NewHttpClient(baseApiUrl, apiKey)

	return &Client{
		Projects:  projects.NewProjectsService(httpClient),
		Languages: languages.NewLanguagesService(httpClient),
	}
}
