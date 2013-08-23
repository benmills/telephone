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

func TestGet(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	response := Request{url: server.URL+"/my_path"}.Get()
	test.Expect(response.ParsedBody).ToEqual("Path:/my_path Method:GET Body:")
	test.Expect(response.StatusCode).ToEqual(200)
}

func TestGetWithBody(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	request := Request{
		url: server.URL,
		body: "this is a body",
	}

	response := request.Get()
	test.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:this is a body")
}

func TestGetWithParameters(t *testing.T) {
	test := quiz.Test(t)
	server := echoServer()
	defer server.Close()

	request := Request{
		url: server.URL,
		parameters: Parameters{
			"foo": "bar",
		},
	}

	response := request.Get()
	test.Expect(response.ParsedBody).ToEqual("Path:/?foo=bar Method:GET Body:")
}
