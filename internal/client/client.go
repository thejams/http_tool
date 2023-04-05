package client

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// Options settings for the http Client object
type Options struct {
	Parallel int
	URLS     []string
}

type httpClient struct {
	parallel int
	urls     []string
	client   http.Client
}

// New returns a new httpClient object
func New(opts Options) *httpClient {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	return &httpClient{
		parallel: opts.Parallel,
		urls:     opts.URLS,
		client:   client,
	}
}

// Request makes a http Get request to specific urls
func (h *httpClient) Request() error {
	if len(h.urls) == 0 {
		return fmt.Errorf("No urls found for HTTP Request")
	}

	var wg sync.WaitGroup
	defaultParallelCalls := 10

	if h.parallel > 0 {
		defaultParallelCalls = h.parallel
	}

	for _, v := range h.urls {
		if runtime.NumGoroutine() <= defaultParallelCalls {
			// ASYNC CALL
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				body, err := h.MakeHttpRequest(url)

				md5, err := GetMD5Hash(body)
				if err == nil {
					fmt.Println(url, md5)
				}
			}(v)
		} else {
			// SYNC CALL
			body, err := h.MakeHttpRequest(v)

			md5, err := GetMD5Hash(body)
			if err == nil {
				fmt.Println(v, md5)
			}
		}
	}

	wg.Wait()
	return nil
}

// MakeHttpRequest makes a http call to a specific url and returns the body of the response
func (h *httpClient) MakeHttpRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error %v, requesting %v url", err, url)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error %v, requesting %v url", err, url)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%v request status != 200", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v  response ioutil ReadAll error: %s", url, err.Error())
	}

	return body, nil
}

// GetMD5Hash returns a string of MD5 hash of a []byte
func GetMD5Hash(bytes []byte) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write(bytes) // []byte(<anything>)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
