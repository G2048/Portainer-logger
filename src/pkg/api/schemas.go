package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

// Response codes is errors: 500 400
type ErrorHttp struct {
	status string
	body   string
}

func (e ErrorHttp) Error() string {
	return fmt.Sprintf("StatusCode: %s. ResponseBody: %s", e.status, e.body)
}
func NewErrorHttp(status string, body string) *ErrorHttp {
	return &ErrorHttp{status, body}
}

type QueryParams map[string]string

type JsonStringMap map[string]string

type Response interface {
	Status() int
	Body() []byte
	DecodeBodyString() (string, error)
	DecodeBodySliceMap() ([]JsonStringMap, error)
	DecodeBodyStruct(any) (any, error)
}

type ResponseApi struct {
	body   []byte
	status int
}

func (r ResponseApi) Body() []byte {
	return r.body
}
func (r ResponseApi) Status() int {
	return r.status
}

func (r ResponseApi) DecodeBodyString() (string, error) {
	return string(r.body), nil
}
func (r ResponseApi) DecodeBodySliceMap() (jsonBody []JsonStringMap, err error) {
	err = json.Unmarshal(r.body, &jsonBody)
	if err != nil {
		slog.Error(err.Error())
	}
	return jsonBody, err
}
func (r ResponseApi) DecodeBodyStruct(jsonBody any) (any, error) {
	err := json.Unmarshal(r.body, &jsonBody)
	if err != nil {
		slog.Error(err.Error())
	}
	return jsonBody, err
}
