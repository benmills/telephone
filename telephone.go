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
}

func Get(url string) Response {
	return Request{Url: url}.Get()
}

func Put(url, body string) Response {
	return Request{Url: url, Body: body}.Put()
}

func (request Request) Get() Response {
	return request.makeRequest("GET")
}

func (request Request) Put() Response {
	return request.makeRequest("PUT")
}

func (request Request) makeRequest(method string) Response {
	httpRequest, _ := http.NewRequest(method, request.buildUrl(), strings.NewReader(request.Body))
	client := http.Client{}
	response, _ := client.Do(httpRequest)
	rawBody, _ := ioutil.ReadAll(response.Body)

	return Response{response, string(rawBody)}
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
