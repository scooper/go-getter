package context

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Request struct {
	r *http.Request
	Method string
	MimeType string
}

func (ggrequest *Request) Args(name string) string {
	return ggrequest.r.URL.Query().Get(name)
}

func (ggrequest *Request) Form(name string) string {
	return ggrequest.r.FormValue(name)
}

func CreateRequest(r *http.Request) *Request {
	return &Request{
		r: r,
		Method: r.Method,
		//MimeType: r.Header["Content-Type"][0], //?
	}
}

// response suggested functions
// Text(content string) response
// Json() response
// Template(path string) response
// Custom(/*options to set most things like body, response code, etc.*/)

type Response struct {
	Response *http.Response
}

func Text(content string) *Response {

	headers := make(http.Header, 0)
	headers.Add("Content-Type", "text/plain")

	return &Response {
		&http.Response{
			Status: "200 OK",
			StatusCode: 200,
			Proto: "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Body: ioutil.NopCloser(bytes.NewBufferString(content)),
			ContentLength: int64(len(content)),
			//Request: ggrequest.r,
			Header: headers,
		},
	}
}
