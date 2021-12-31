package model

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// type ApiRequest interface{}

type ApiResponse struct {
	Response struct {
		Status  string `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"response"`
	Result *map[string]interface{} `json:"result"`
}

type POEditorApiTime struct {
	time.Time
}

func (t *POEditorApiTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse("2006-01-02T15:04:05+0000", s)
	
	return
}

type POEditorApiBoolean bool

func (b *POEditorApiBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*b = true
	} else if asString == "0" || asString == "false" {
		*b = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}
