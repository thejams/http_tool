package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"http_tool/internal/client"
)

type testCase struct {
	input  string
	output string
}

func TestRequest(t *testing.T) {
	t.Run("should make a correct Request", func(t *testing.T) {
		mockServerOK := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
			rw.Header().Set("Content-type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("status OK"))
		}))

		defer mockServerOK.Close()

		opts := client.Options{
			Parallel: 1,
			URLS:     []string{mockServerOK.URL},
		}
		customClient := client.New(opts)

		err := customClient.Request()
		if err != nil {
			t.Error("Expected", nil, "Got", err)
		}
	})

	t.Run("should return error when no URL is received", func(t *testing.T) {
		opts := client.Options{
			Parallel: 1,
		}
		customClient := client.New(opts)

		err := customClient.Request()
		if err == nil {
			t.Error("Expected", "Error", "Got", err)
		}
	})
}

func TestMD5Hash(t *testing.T) {
	testCases := []testCase{
		{"MyString", "10dd3a9cb9c68ff6403e78f7e2da14d7"},
		{"MyString2", "51de57718fb4b5ec5365c092f6123f0b"},
	}

	for _, v := range testCases {
		res, _ := client.GetMD5Hash([]byte(v.input))
		if res != v.output {
			t.Error("Expected:", v.output, "Got:", res)
		}
	}
}

func TestMakeHttpRequest(t *testing.T) {
	t.Run("should return a correct HTTP Request", func(t *testing.T) {
		mockServerOK := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
			rw.Header().Set("Content-type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("status OK"))
		}))
		defer mockServerOK.Close()

		opts := client.Options{
			Parallel: 2,
			URLS:     []string{mockServerOK.URL},
		}
		customClient := client.New(opts)

		okCase := testCase{
			input:  mockServerOK.URL,
			output: "status OK",
		}

		res, _ := customClient.MakeHttpRequest(okCase.input)
		sRes := string(res)
		if sRes != okCase.output {
			t.Error("Expected:", okCase.output, "Got:", sRes)
		}
	})

	t.Run("should return empty string and error when HTTP response fails", func(t *testing.T) {
		mockServerNOK := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
			rw.Header().Set("Content-type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("status NOK"))
		}))

		defer mockServerNOK.Close()

		opts := client.Options{
			Parallel: 2,
			URLS:     []string{mockServerNOK.URL},
		}
		customClient := client.New(opts)

		noOkCase := testCase{
			input:  mockServerNOK.URL,
			output: "",
		}

		res, err := customClient.MakeHttpRequest(noOkCase.input)
		sRes := string(res)
		if sRes != noOkCase.output {
			t.Error("Expected:", noOkCase.output, "Got:", sRes)
		}
		if err == nil {
			t.Error("Expected:", "Not Nill", "Got:", nil)
		}
	})
}
