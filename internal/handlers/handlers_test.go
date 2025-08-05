package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals_quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors_suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search_availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make_reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post_search_availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-10"},
	}, http.StatusOK},
	{"post_search_availability_json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-10"},
	}, http.StatusOK},
	{"post_make_reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Doe"},
		{key: "email", value: "jdoe@gmail.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			response, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, params := range e.params {
				values.Add(params.key, params.value)
			}
			response, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
