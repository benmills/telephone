package telephone

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Parameters map[string]string

type Request struct {
	url string
	body string
	parameters Parameters
}

type Response struct {
	*http.Response
	ParsedBody string
}

func (request Request) Get() Response {
	return makeRequest("GET", request.buildUrl(), request.body)

}

func makeRequest(method string, url, body string) Response {
	request, _ := http.NewRequest(method, url, strings.NewReader(body))
	client := http.Client{}
	response, _ := client.Do(request)
	rawBody, _ := ioutil.ReadAll(response.Body)

	return Response{response, string(rawBody)}
}

func (request Request) buildUrl() string {
	return request.url+request.encodeParams()
}

func (request Request) encodeParams() string {
	if len(request.parameters) <= 0 {
		return ""
	}

	values := url.Values{}
	for key, value := range(request.parameters) {
		values.Set(key, value)
	}

	return "?"+values.Encode()
}
