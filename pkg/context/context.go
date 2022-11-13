package context

import (
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
	}
}

// response suggested functions
// Json() response
// Template(path string) response
// Custom(/*options to set most things like body, response code, etc.*/)

type Response struct {
	//Response *http.Response,
	StatusCode int
	Headers map[string]string
	Body string
}

func Text(content string) *Response {
	response := &Response{
		Body: content,
		StatusCode: 200,
		Headers: make(map[string]string),
	}

	response.Headers["Content-Type"] = "text/plain"

	return response
}
