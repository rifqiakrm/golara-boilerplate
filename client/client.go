package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func StartMocks() {
	enabledMocks = true
}

func StopMocks() {
	enabledMocks = false
}

func AddMock(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

//{}
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mockup for the given request")
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	response, err := client.Do(request)
	return response, nil
}

func Get(url string, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodGet, url)]
		if mock == nil {
			return nil, errors.New("No mockup for the given request")
		}
		return mock.Response, mock.Err
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header = headers

	client := http.Client{}

	response, err := client.Do(request)

	return response, nil
}
