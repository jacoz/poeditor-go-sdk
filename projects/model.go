package projects

import "github.com/jacoz/poeditor-go-sdk/model"

type ExportRequestType string

const (
	KeyValueJSON ExportRequestType = "key_value_json"
)

type apiResponseList struct {
	Projects []Project `json:"projects"`
}
type apiResponseView struct {
	Project ProjectDetails `json:"project"`
}

type Project struct {
	ID        int64                    `json:"id"`
	Name      string                   `json:"name"`
	Public    model.POEditorApiBoolean `json:"public"`
	Open      model.POEditorApiBoolean `json:"open"`
	CreatedAt model.POEditorApiTime    `json:"created"`
}

type ListResponse []Project

type ViewRequest struct {
	ProjectID int64 `json:"id"`
}

type ProjectDetails struct {
	Project
	Description       string `json:"description"`
	ReferenceLanguage string `json:"reference_language"`
	FallbackLanguage  string `json:"fallback_language"`
	TermsCount        int64  `json:"terms"`
}

type ExportRequest struct {
	ProjectID int64             `json:"id"`
	Language  string            `json:"language"`
	Type      ExportRequestType `json:"type"`
	//Filters  *[]string `json:"filters"`
	//Order    *string   `json:"order"`
	//Tags     *[]string `json:"tags"`
	//Options  *string   `json:"options"`
}

type ExportResponse struct {
	Url string `json:"url"`
}
