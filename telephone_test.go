package telephone

import (
	"github.com/benmills/quiz"

	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func echoServer() *httptest.Server {
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawBody, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "Path:%s Method:%s Body:%s", r.RequestURI, r.Method, rawBody)
	})

	return httptest.NewServer(server)
}

func TestRequestWithBody(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	request := Request{
		Url: server.URL,
		Body: "this is a body",
	}

	response := request.makeRequest("GET")
	test.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:this is a body")
}

func TestRequestWithParameters(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	request := Request{
		Url: server.URL,
		Parameters: Parameters{
			"foo": "bar",
		},
	}

	response := request.makeRequest("GET")
	test.Expect(response.ParsedBody).ToEqual("Path:/?foo=bar Method:GET Body:")
}

func TestGet(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	response := Request{Url: server.URL}.Get()
	test.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:")
}

func TestGetHelperMethod(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	response := Get(server.URL+"/my_path")
	test.Expect(response.ParsedBody).ToEqual("Path:/my_path Method:GET Body:")
}