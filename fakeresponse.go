//Package fakeresponse provides a mocked ResponseWriter
package fakeresponse

import (
	"net/http"
	"testing"
)

//FakeResponse type
type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

//NewFakeResponse create a new fake respone
func NewFakeResponse(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

//Header get the response headers
func (r *FakeResponse) Header() http.Header {
	return r.headers
}

//Write the reposnse
func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

//WriteHeader set the status code
func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

//Assert simple assertion method
func (r *FakeResponse) Assert(status int, body string) {
	if r.status != status {
		r.t.Errorf("expected status %+v to equal %+v", r.status, status)
	}
	if string(r.body) != body {
		r.t.Errorf("expected body %+v to equal %+v", string(r.body), body)
	}
}
