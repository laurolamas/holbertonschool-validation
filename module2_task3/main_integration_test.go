package main

import (
	//nolint:staticcheck
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server(t *testing.T) {
	if testing.Short() {
		t.Skip("Flag `-short` provided: skipping Integration Tests.")
	}

	tests := []struct {
		name         string
		URI          string
		responseCode int
		body         string
	}{
		{
			name:         "Home page",
			URI:          "",
			responseCode: 404,
			body:         "404 page not found\n",
		},
		{
			name:         "Hello page",
			URI:          "/hello?name=Holberton",
			responseCode: 200,
			body:         "Hello Holberton!",
		},
		{
			name:         "No Parameter",
			URI:          "/hello?",
			responseCode: 200,
			body:         "Hello there!",
		},
		{
			name:         "HealthCheck",
			URI:          "/health",
			responseCode: 200,
			body:         "ALIVE",
		},
		{
			name:         "HelloHandler",
			URI:          "/hello?name=",
			responseCode: 400,
			body:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(setupRouter())
			defer ts.Close()

			res, err := http.Get(ts.URL + tt.URI)
			if err != nil {
				t.Fatal(err)
			}

			// Check that the status code is what you expect.
			expectedCode := tt.responseCode
			gotCode := res.StatusCode
			if gotCode != expectedCode {
				t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
			}

			// Check that the response body is what you expect.
			expectedBody := tt.body
			bodyBytes, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			gotBody := string(bodyBytes)
			if gotBody != expectedBody {
				t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
			}
		})
	}
}
func TestIntegration_SetupRouter(t *testing.T) {
	router := setupRouter()

	// Create a new HTTP request for the path "/health"
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the router
	router.ServeHTTP(rr, req)

	// Check the response status code
	expectedCode := http.StatusOK
	if rr.Code != expectedCode {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, expectedCode)
	}

	// Check the response body
	expectedBody := "ALIVE"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedBody)
	}
}

func TestIntegration_HelloHandler(t *testing.T) {
	router := setupRouter()

	// Create a new HTTP request for the path "/hello" with a query parameter "name"
	req, err := http.NewRequest("GET", "/hello?name=John", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the router
	router.ServeHTTP(rr, req)

	// Check the response status code
	expectedCode := http.StatusOK
	if rr.Code != expectedCode {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, expectedCode)
	}

	// Check the response body
	expectedBody := "Hello John!"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedBody)
	}
}

func TestIntegration_HealthCheckHandler(t *testing.T) {
	router := setupRouter()

	// Create a new HTTP request for the path "/health"
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the router
	router.ServeHTTP(rr, req)

	// Check the response status code
	expectedCode := http.StatusOK
	if rr.Code != expectedCode {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, expectedCode)
	}

	// Check the response body
	expectedBody := "ALIVE"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedBody)
	}
}
