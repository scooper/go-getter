package context

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/scooper/go-getter/pkg/utils"
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
	Error error
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

func Template(filepath string, data any) *Response {

	tfile, ferr := utils.GetTemplate(filepath)
	headers :=  make(map[string]string)

	if ferr != nil {
		return &Response {
			StatusCode: 500,
			Error: ferr,
			Headers: headers,
		}
	}

	defer tfile.Close()

	var buf bytes.Buffer
	io.Copy(&buf, tfile)
	fileAsStr := string(buf.Bytes())

	t, terr := template.New("").Parse(fileAsStr)
	if terr != nil {
		return &Response {
			StatusCode: 500,
			Error: terr,
			Headers: headers,
		}
	}

	bodyBuf := new(bytes.Buffer)
	
	t.Execute(bodyBuf, data)

	response := &Response {
		StatusCode: 200, // TODO: easier way to change this?
		Body: bodyBuf.String(),
		Headers: headers,
	}

	response.Headers["Content-Type"] = "text/html"

	return response
}
