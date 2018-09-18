package testhelper

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type HttpTestRequest struct {
	Method      string
	Path        string
	RequestBody string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func NewHttpTest(method string, path string, requestBody string, handler func(http.ResponseWriter, *http.Request)) *HttpTestRequest {
	return &HttpTestRequest{
		Method:      method,
		Path:        path,
		RequestBody: requestBody,
		HandlerFunc: handler,
	}
}

func (t *HttpTestRequest) DoRequestTest() (rrReturn *httptest.ResponseRecorder, err error) {

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(t.Method, t.Path, strings.NewReader(t.RequestBody))
	if err != nil {
		return nil, err
	}
	http.HandlerFunc(t.HandlerFunc).ServeHTTP(rr, req)

	return rr, nil
}
