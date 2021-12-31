package languages

import "github.com/jacoz/poeditor-go-sdk/model"

type apiResponseLanguage struct {
	Languages []Language `json:"languages"`
}

type ListRequest struct {
	ProjectID int64 `json:"id"`
}

type Language struct {
	Name          string                `json:"name"`
	Code          string                `json:"code"`
	Translations  int64                 `json:"translations"`
	Percentage    float64               `json:"percentage"`
	LastUpdatedAt model.POEditorApiTime `json:"updated"`
}

type ListResponse []Language
