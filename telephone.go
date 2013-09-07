package telephone

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Parameters map[string]string

type Request struct {
	Url string
	Body string
	Parameters Parameters
}

type Response struct {
	*http.Response
	ParsedBody string
	Success bool
}

func Get(url string) Response {
	return Request{Url: url}.Get()
}

func Put(url, body string) Response {
	return Request{Url: url, Body: body}.Put()
}

func Post(url, body string) Response {
	return Request{Url: url, Body: body}.Post()
}

func (request Request) Get() Response {
	return request.makeRequest("GET")
}

func (request Request) Put() Response {
	return request.makeRequest("PUT")
}

func (request Request) Post() Response {
	return request.makeRequest("POST")
}

func (request Request) makeRequest(method string) Response {
	httpRequest, _ := http.NewRequest(method, request.buildUrl(), strings.NewReader(request.Body))
	client := http.Client{}

	response, requestErr := client.Do(httpRequest)
	if requestErr != nil {
		return Response{&http.Response{}, "", false}
	}

	rawBody, ioErr := ioutil.ReadAll(response.Body)
	if ioErr != nil {
		return Response{&http.Response{}, "", false}
	}

	return Response{response, string(rawBody), true}
}

func (request Request) buildUrl() string {
	return request.Url+request.encodeParams()
}

func (request Request) encodeParams() string {
	if len(request.Parameters) <= 0 {
		return ""
	}

	values := url.Values{}
	for key, value := range(request.Parameters) {
		values.Set(key, value)
	}

	return "?"+values.Encode()
}
