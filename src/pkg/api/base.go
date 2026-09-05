package api

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"portainer-logger/src/pkg/utils"
)

type MethodHttp string

const MethodGet MethodHttp = "GET"
const MethodPatch MethodHttp = "PATCH"
const MethodPost MethodHttp = "POST"
const MethodPut MethodHttp = "PUT"
const MethodDelete MethodHttp = "DELETE"
const MethodOptions MethodHttp = "OPTIONS"
const MethodHead MethodHttp = "HEAD"
const MethodTrace MethodHttp = "TRACE"
const MethodConnect MethodHttp = "CONNECT"

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

type Response struct {
	Body   string
	Status int
}
type QueryParams map[string]string

type HttpApi struct {
	URL string
	http.Client
	// req http.Request
	header http.Header
}

func NewHttpApi(url string) *HttpApi {
	return &HttpApi{URL: url, header: http.Header{}}
}

func (h *HttpApi) AddHeader(key, value string) {
	h.header.Add(key, value)
}
func (h *HttpApi) baseAddHeaders() {
	h.header.Add("Content-Type", "application/json")
	h.header.Add("accept", "application/json")
}
func (h *HttpApi) DecodeBody(res *http.Response, clientResp *Response) *Response {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error(err.Error())
	} else {
		clientResp.Body = string(body)
	}
	return clientResp
}

func (h *HttpApi) Response(method MethodHttp, endpoint string, payload []byte, queryParams string) (*Response, *ErrorHttp) {
	var clientResponse = &Response{"", 0}
	var respErr *ErrorHttp = nil

	url := utils.MustResult(url.JoinPath(h.URL, endpoint, url.QueryEscape(queryParams)))
	slog.Info(fmt.Sprintf("Request to url: %s", url))

	h.baseAddHeaders()
	req := utils.MustResult(http.NewRequest(string(method), url, bytes.NewBuffer(payload)))
	req.URL.Query()
	req.Header = h.header

	res := utils.MustResult(h.Do(req))
	clientResponse.Status = res.StatusCode
	// flush headers
	h.header = http.Header{}

	h.DecodeBody(res, clientResponse)
	if res.StatusCode >= 500 {
		respErr = NewErrorHttp(res.Status, "")
	}
	if res.StatusCode >= 400 {
		respErr = NewErrorHttp(res.Status, clientResponse.Body)
	}

	return clientResponse, respErr
}
func (h *HttpApi) Get(url string, queryParams string) (*Response, *ErrorHttp) {
	return h.Response(MethodGet, url, nil, queryParams)
}
func (h *HttpApi) Post(url string, payload []byte) (*Response, *ErrorHttp) {
	return h.Response(MethodPost, url, payload, "")
}
func (h *HttpApi) Patch(url string, payload []byte) (*Response, *ErrorHttp) {
	return h.Response(MethodPatch, url, payload, "")
}
func (h *HttpApi) Put(url string, payload []byte) (*Response, *ErrorHttp) {
	return h.Response(MethodPut, url, payload, "")
}
func (h *HttpApi) Delete(url string, payload []byte) (*Response, *ErrorHttp) {
	return h.Response(MethodDelete, url, payload, "")
}
