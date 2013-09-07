package telephone

import (
	e "github.com/lionelbarrow/examples"

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

func TestTelephone(t *testing.T) {
	e.Describe("Request", t,
		e.It("fails due to an invalid URL", func (ex *e.Example) {
			response := Request{Url: "!!"}.makeRequest("GET")
			ex.Expect(response.Success).ToBeFalse()
		}),

		e.It("can have a body", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			request := Request{
				Url: server.URL,
				Body: "this is a body",
			}

			response := request.makeRequest("GET")
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:this is a body")
		}),

		e.It("can have parameters", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			request := Request{
				Url: server.URL,
				Parameters: Parameters{
					"foo": "bar",
				},
			}

			response := request.makeRequest("GET")
			ex.Expect(response.ParsedBody).ToEqual("Path:/?foo=bar Method:GET Body:")
		}),

		e.It("can make a GET request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Request{Url: server.URL}.Get()
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:")
		}),

		e.It("can make a PUT request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Request{Url: server.URL}.Put()
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:PUT Body:")
		}),

		e.It("can make a POST request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Request{Url: server.URL}.Post()
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:POST Body:")
		}),
	)

	e.Describe("Get", t,
		e.It("can make a GET request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Get(server.URL)
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:GET Body:")
		}),
	)

	e.Describe("Put", t,
		e.It("can make a PUT request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Put(server.URL, "")
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:PUT Body:")
		}),

		e.It("can make a PUT request with a body", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Put(server.URL, "my body")
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:PUT Body:my body")
		}),
	)

	e.Describe("Post", t,
		e.It("can make a POST request", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Post(server.URL, "")
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:POST Body:")
		}),

		e.It("can make a PUT request with a body", func (ex *e.Example) {
			server := echoServer()
			defer server.Close()

			response := Post(server.URL, "my body")
			ex.Expect(response.ParsedBody).ToEqual("Path:/ Method:POST Body:my body")
		}),
	)
}
